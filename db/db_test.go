package db

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/mitchellh/go-homedir"
)

func TestInit(t *testing.T) {
	home, _ := homedir.Dir()
	dbfilepath := filepath.Join(home, "Sites.db")
	want := []byte("sites")
	got := Init(dbfilepath)
	fmt.Printf("wanted %s  got %s ", want, got)
}

func TestAddsite(t *testing.T) {
	u, _ := time.ParseDuration("1s")
	s, _ := AddSite("www.google.com", u)
	want := "1s"
	got := s
	fmt.Printf("wanted %s  got %s ", want, got)
}

//func TestDeletesite(t *testing.T) {
//	x := DeleteSite(&Sites{})
//	want := ""
//	got := x
//	fmt.Printf("wanted %s  got %s ", want, got)
//}
