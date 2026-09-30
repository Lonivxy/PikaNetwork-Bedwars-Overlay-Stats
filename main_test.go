//go:build windows

package main

import "testing"

func TestFormatLeaderboardPlace(t *testing.T) {
	previous := fuzzyLB
	defer func() { fuzzyLB = previous }()

	fuzzyLB = true
	cases := map[int64]string{
		1:      "#1",
		1000:   "#1000",
		1001:   "#1k+",
		12345:  "#12k+",
		100070: "#100k+",
	}
	for place, want := range cases {
		if got := formatLeaderboardPlace(place); got != want {
			t.Fatalf("formatLeaderboardPlace(%d) = %q, want %q", place, got, want)
		}
	}

	fuzzyLB = false
	if got := formatLeaderboardPlace(12345); got != "#12345" {
		t.Fatalf("fuzzy-disabled place = %q, want #12345", got)
	}
}

func TestRankOrderAndColumnComparison(t *testing.T) {
	if rankOrder("Champion") <= rankOrder("Titan") || rankOrder("Titan") <= rankOrder("Elite") || rankOrder("Elite") <= rankOrder("VIP") {
		t.Fatal("rank order is not Champion > Titan > Elite > VIP")
	}

	champion := &PlayerStats{Username: "alpha", GamesRank: "Champion", Kills: 10}
	vip := &PlayerStats{Username: "bravo", GamesRank: "VIP", Kills: 99}
	if comparePlayersByColumn(champion, vip, "rank") <= 0 {
		t.Fatal("rank comparison did not put Champion above VIP")
	}
	if comparePlayersByColumn(champion, vip, "kills") >= 0 {
		t.Fatal("kills comparison did not order the lower value first")
	}
}

func TestPlayerLeaderboardEntryDoesNotUseAnotherPlayer(t *testing.T) {
	response := map[string]any{
		"Kills": map[string]any{
			"entries": []any{
				map[string]any{"id": "OtherPlayer", "value": "999", "place": float64(1)},
				map[string]any{"id": "SomeoneElse", "value": "888", "place": float64(2)},
			},
		},
	}
	value, place := playerLeaderboardEntry(response, "Kills", "TargetPlayer")
	if value != 0 || place != 0 {
		t.Fatalf("unmatched multi-entry response returned value=%d place=%d", value, place)
	}
}
