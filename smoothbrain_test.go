package smoothbrain

import (
	"log"
	"testing"
)

func TestSmoothBrain(t *testing.T) {
	const sampleJSON = `
{
  "store": {
    "name": "Acme Supplies",
    "location": {
      "address": "456 Maple Avenue",
      "city": "Springfield",
      "state": "CA",
      "zipCode": "90123"
    },
    "inventory": [
      {
        "category": "Tools",
        "sku": "HAM001",
        "description": "Hammer",
        "price": 12.99,
        "stock": 25
      },
      {
        "category": "Hardware",
        "sku": "SCR015",
        "description": "Screwdriver Set",
        "price": 24.95,
        "stock": 10,
        "variations": [
          { "size": "Small", "inStock": true },
          { "size": "Medium", "inStock": false }
        ]
      }
    ],
    "employees": [
      { "name": "Bob Johnson", "role": "Manager"},
      { "name": "Sarah Lee", "role": "Sales Associate"}
    ]
  }
}`
	sbj := New()
	err := sbj.Unmarshal([]byte(sampleJSON))
	if err != nil {
		t.Error(err)
	}
	data, err := sbj.Marshal()
	if err != nil {
		t.Error(err)
	}
	log.Println(string(data))
}
