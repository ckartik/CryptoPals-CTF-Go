package main

import "testing"

func TestFixedXOR(t *testing.T) {
	S1C2Input := "1c0111001f010100061a024b53535009181c"
	S1C2Input2 := "686974207468652062756c6c277320657965"
	S1C2Answer := "746865206b696420646f6e277420706c6179"
	S1C2Result := FixedXOR(S1C2Input, S1C2Input2)
	if S1C2Result != S1C2Answer {
		t.Errorf("Error: Expected 746865206b696420646f6e277420706c6179, got %v", S1C2Result)
	}
}
