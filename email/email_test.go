package email

import (
	"encoding/json"
	"testing"
)


func BenchmarkUnmarshalNomarlJSON(b *testing.B) {
	em := &Email{}

	bb := []byte(`{"subject":"Xin chao","body":"Mot hai ba bon nam sau bay tam chin muoi muoi mot muoi hai muoi ba muoi bon muoi nam","to":["1@gmail.com","2@gmail.com","3@gmail.com"],"cc":["1@gmail.com","2@gmail.com","3@gmail.com"],"bcc":["1@gmail.com","2@gmail.com","3@gmail.com"]}`)
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bb, em)
	}
}

func BenchmarkUnmarshalEasyJSON(b *testing.B) {
	em := &Email{}

	bb := []byte(`{"subject":"Xin chao","body":"Mot hai ba bon nam sau bay tam chin muoi muoi mot muoi hai muoi ba muoi bon muoi nam","to":["1@gmail.com","2@gmail.com","3@gmail.com"],"cc":["1@gmail.com","2@gmail.com","3@gmail.com"],"bcc":["1@gmail.com","2@gmail.com","3@gmail.com"]}`)
	for i := 0; i < b.N; i++ {
		em.UnmarshalJSON(bb)
	}
}

func BenchmarkNomarlJSON(b *testing.B) {
	em := Email{
		Subject: "Xin chao",
		Body:    "Mot hai ba bon nam sau bay tam chin muoi muoi mot muoi hai muoi ba muoi bon muoi nam",
		To:      []string{"1@gmail.com", "2@gmail.com", "3@gmail.com"},
		Cc:      []string{"1@gmail.com", "2@gmail.com", "3@gmail.com"},
		Bcc:     []string{"1@gmail.com", "2@gmail.com", "3@gmail.com"},
	}

	for i := 0; i < b.N; i++ {
		json.Marshal(em)
	}
}

func BenchmarkEasyJSON(b *testing.B) {
	em := Email{
		Subject: "Xin chao",
		Body:    "Mot hai ba bon nam sau bay tam chin muoi muoi mot muoi hai muoi ba muoi bon muoi nam",
		To:      []string{"1@gmail.com", "2@gmail.com", "3@gmail.com"},
		Cc:      []string{"1@gmail.com", "2@gmail.com", "3@gmail.com"},
		Bcc:     []string{"1@gmail.com", "2@gmail.com", "3@gmail.com"},
	}

	for i := 0; i < b.N; i++ {
		em.MarshalJSON()
	}
}
