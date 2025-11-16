package data

import "pestapi/model"

var Pests = []model.Pest{
	{
		ID:             1,
		CommonName:     "Penggerek Buah Kopi (PBKo)",
		ScientificName: "Hypothenemus hampei",
		PestType:       "Serangga",
		AffectedParts:  []string{"Buah", "Biji"},
		Description:    "Serangga hitam kecil yang merusak buah kopi.",
		Symptoms:       []string{"Lubang kecil pada buah", "Buah rontok dini"},
		ImageURL:       "https://example.com/pbko.jpg",
		ControlMethods: model.ControlMethods{
			Organic:  []string{"Ambil buah terserang", "Gunakan atraktan"},
			Chemical: []string{"Pestisida Fipronil"},
		},
	},
}
