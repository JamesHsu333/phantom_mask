package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/JamesHsu333/kdan/config"
	data "github.com/JamesHsu333/kdan/internal/data/test"
	"github.com/JamesHsu333/kdan/pkg/database/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
)

type User struct {
	Name              string            `json:"name"`
	CashBalance       float64           `json:"cashBalance"`
	PurchaseHistories []PurchaseHistory `json:"purchaseHistories"`
}

type PurchaseHistory struct {
	PharmacyName      string  `json:"pharmacyName"`
	MaskName          string  `json:"maskName"`
	TransactionAmount float64 `json:"transactionAmount"`
	TransactionDate   string  `json:"transactionDate"`
}

type Pharmacy struct {
	Name         string  `json:"name"`
	CashBalance  float64 `json:"cashBalance"`
	OpeningHours string  `json:"openingHours"`
	Masks        []Mask  `json:"masks"`
}

type Mask struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := initPostgreSQL(ctx, cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	users := readUsers()
	pharmacies := readPharmacies()

	querier := data.New(db)

	ingestPharmacies(ctx, querier, pharmacies)
	ingestUsers(ctx, querier, users)
}

func readUsers() []User {
	userJson, err := os.Open("./data/users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer userJson.Close()

	var users []User
	userBytes, _ := ioutil.ReadAll(userJson)
	json.Unmarshal([]byte(userBytes), &users)

	return users
}

func readPharmacies() []Pharmacy {
	pharmacyJson, err := os.Open("./data/pharmacies.json")
	if err != nil {
		log.Fatal(err)
	}
	defer pharmacyJson.Close()

	var pharmacies []Pharmacy
	pharmacyBytes, _ := ioutil.ReadAll(pharmacyJson)
	json.Unmarshal([]byte(pharmacyBytes), &pharmacies)

	return pharmacies
}

var weekdayMap = map[string]int32{
	"Mon":  0,
	"Tue":  1,
	"Wed":  2,
	"Thur": 3,
	"Fri":  4,
	"Sat":  5,
	"Sun":  6,
}

func parseOpeningHours(s string) [][2]int32 {
	sa := strings.Split(s, " / ")
	re := regexp.MustCompile(`(\d{2}):(\d{2}) - (\d{2}):(\d{2})`)

	r := make([][2]int32, 0, 10)

	for _, ss := range sa {
		p := re.FindStringIndex(ss)
		ranges := strings.Split(ss[p[0]:p[1]], " - ")
		days := strings.Split(ss[:p[0]-1], " - ")
		if len(days) != 1 {
		} else {
			days = strings.Split(ss[:p[0]-1], ", ")
		}

		for _, d := range days {
			dd := weekdayMap[d]
			h1, _ := strconv.ParseInt(ranges[0][:2], 10, 0)
			m1, _ := strconv.ParseInt(ranges[0][2:], 10, 0)
			t1 := int32(h1)*60 + int32(m1)
			h2, _ := strconv.ParseInt(ranges[1][:2], 10, 0)
			m2, _ := strconv.ParseInt(ranges[1][2:], 10, 0)
			t2 := int32(h2)*60 + int32(m2)

			r1 := dd*24*60 + t1
			if t1 > t2 {
				if (dd + 1) > 6 {
					dd = 0
				} else {
					dd += 1
				}
			}
			r2 := dd*24*60 + t2

			if r1 > r2 {
				r = append(r, [2]int32{r1, 10080})
				r = append(r, [2]int32{0, r2})
			} else {
				r = append(r, [2]int32{r1, r2})
			}
		}
	}
	return r
}

var pharmaciesMap = make(map[string]int32)

func ingestPharmacies(ctx context.Context, querier *data.Queries, pharmacies []Pharmacy) {
	for _, pharmacy := range pharmacies {
		minutesRanges := parseOpeningHours(pharmacy.OpeningHours)
		openingHours := make(pgtype.Multirange[pgtype.Range[pgtype.Int4]], 0, len(minutesRanges))

		for _, r := range minutesRanges {
			openingHours = append(openingHours, pgtype.Range[pgtype.Int4]{
				Lower:     pgtype.Int4{Int32: r[0], Valid: true},
				Upper:     pgtype.Int4{Int32: r[1], Valid: true},
				LowerType: pgtype.Inclusive,
				UpperType: pgtype.Inclusive,
				Valid:     true,
			})
		}

		args := data.CreatePharmacyParams{
			Name:                    pharmacy.Name,
			OpeningHours:            openingHours,
			OpeningHoursDescription: pharmacy.OpeningHours,
			CashBalance:             &pharmacy.CashBalance,
		}

		resp, err := querier.CreatePharmacy(ctx, args)
		if err != nil {
			log.Fatal(err)
		}
		pharmaciesMap[pharmacy.Name] = resp.ID
		ingestMasks(ctx, querier, pharmacy.Masks, resp.ID)
	}
}

var masksMap = make(map[string]int32)

func ingestMasks(ctx context.Context, querier *data.Queries, masks []Mask, pharmacyId int32) {
	for _, mask := range masks {
		if _, exist := masksMap[mask.Name]; exist {
			continue
		}

		resp, err := querier.CreateMask(ctx, mask.Name)
		if err != nil {
			log.Fatal(err)
		}
		masksMap[mask.Name] = resp.ID

		querier.CreateMaskPrice(ctx, data.CreateMaskPriceParams{
			PharmacyID: pharmacyId,
			MaskID:     resp.ID,
			Price:      mask.Price,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

var usersMap = make(map[string]int32)

func ingestUsers(ctx context.Context, querier *data.Queries, users []User) {
	for _, user := range users {
		args := data.CreateUserParams{
			Name:        user.Name,
			CashBalance: &user.CashBalance,
		}
		resp, err := querier.CreateUser(ctx, args)
		if err != nil {
			log.Fatal(err)
		}
		usersMap[user.Name] = resp.ID
		ingestPurchaseHistory(ctx, querier, user.PurchaseHistories, resp.ID)
	}
}

func ingestPurchaseHistory(ctx context.Context, querier *data.Queries, purchases []PurchaseHistory, userId int32) {
	for _, p := range purchases {
		t, err := time.Parse(time.DateTime, p.TransactionDate)
		if err != nil {
			log.Fatal(err)
		}

		args := data.CreatePurchaseParams{
			UserID:            userId,
			PharmacyID:        pharmaciesMap[p.PharmacyName],
			MaskID:            masksMap[p.MaskName],
			TransactionAmount: p.TransactionAmount,
			TransactionDate:   t,
		}
		_, err = querier.CreatePurchase(ctx, args)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func initPostgreSQL(ctx context.Context, cfg postgres.Config) (*pgx.Conn, error) {
	db, err := postgres.NewPsqlDB(ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Postgresql init")
	}
	return db, err
}
