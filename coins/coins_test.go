package coins

// func TestNextDenominationPlatinum(t *testing.T) {
// 	expected := Gold
// 	got, err := NextDenomination(Platinum)

// 	if got != expected {
// 		t.Errorf("got %v, wanted %v", got.Denomination, expected.Denomination)
// 	}

// 	if err != nil {
// 		t.Errorf("should not return error %v", err)
// 	}
// }

// func TestNextDenominationGold(t *testing.T) {
// 	expected := Electrum
// 	got, err := NextDenomination(Gold)

// 	if got != expected {
// 		t.Errorf("got %v, wanted %v", got.Denomination, expected.Denomination)
// 	}

// 	if err != nil {
// 		t.Errorf("should not return error %v", err)
// 	}
// }

// func TestNextDenominationElectrum(t *testing.T) {
// 	expected := Silver
// 	got, err := NextDenomination(Electrum)

// 	if got != expected {
// 		t.Errorf("got %v, wanted %v", got.Denomination, expected.Denomination)
// 	}

// 	if err != nil {
// 		t.Errorf("should not return error %v", err)
// 	}
// }

// func TestNextDenominationSilver(t *testing.T) {
// 	expected := Copper
// 	got, err := NextDenomination(Silver)

// 	if got != expected {
// 		t.Errorf("got %v, wanted %v", got.Denomination, expected.Denomination)
// 	}

// 	if err != nil {
// 		t.Errorf("should not return error %v", err)
// 	}
// }

// func TestNextDenominationCopper(t *testing.T) {
// 	expected := Copper
// 	got, err := NextDenomination(Copper)

// 	if got != expected {
// 		t.Errorf("got %v, wanted %v", got.Denomination, expected.Denomination)
// 	}

// 	if err == nil {
// 		t.Errorf("copper should result in an error")
// 	}
// }

// func TestNextDenominationNotValid(t *testing.T) {
// 	mud := Piece{
// 		Denomination: "mp",
// 		ExchangeRate: 99999,
// 	}

// 	expected := Copper
// 	got, err := NextDenomination(mud)

// 	if got != expected {
// 		t.Errorf("got %v, wanted %v", got.Denomination, expected.Denomination)
// 	}

// 	if err == nil {
// 		t.Errorf("copper should result in an error")
// 	}
// }
