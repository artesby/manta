package manta

import (
	"testing"

	"github.com/dotabuff/manta/dota"
	"github.com/stretchr/testify/assert"
)

func TestMatch1788648401(t *testing.T) { testScenarios[1788648401].test(t) }

func TestMatch1786687320(t *testing.T) { testScenarios[1786687320].test(t) }
func TestMatch1785937100(t *testing.T) { testScenarios[1785937100].test(t) }
func TestMatch1785899023(t *testing.T) { testScenarios[1785899023].test(t) }
func TestMatch1785874713(t *testing.T) { testScenarios[1785874713].test(t) }
func TestMatch1781640270(t *testing.T) { testScenarios[1781640270].test(t) }
func TestMatch1764592109(t *testing.T) { testScenarios[1764592109].test(t) }
func TestMatch1763193771(t *testing.T) { testScenarios[1763193771].test(t) }
func TestMatch1763177231(t *testing.T) { testScenarios[1763177231].test(t) }
func TestMatch1734886116(t *testing.T) { testScenarios[1734886116].test(t) }
func TestMatch1731962898(t *testing.T) { testScenarios[1731962898].test(t) }
func TestMatch1716444111(t *testing.T) { testScenarios[1716444111].test(t) }
func TestMatch1712853372(t *testing.T) { testScenarios[1712853372].test(t) }
func TestMatch1648457986(t *testing.T) { testScenarios[1648457986].test(t) }
func TestMatch1605340040(t *testing.T) { testScenarios[1605340040].test(t) }
func TestMatch1582611189(t *testing.T) { testScenarios[1582611189].test(t) }
func TestMatch1560315800(t *testing.T) { testScenarios[1560315800].test(t) }
func TestMatch1560294294(t *testing.T) { testScenarios[1560294294].test(t) }
func TestMatch1560289528(t *testing.T) { testScenarios[1560289528].test(t) }

type testScenario struct {
	matchId                string
	replayUrl              string
	debugLevel             uint
	debugTick              uint32
	skipPacketEntities     bool
	expectGameBuild        uint32
	expectEntityEvents     int32
	expectCombatLogDamage  int32
	expectCombatLogHealing int32
	expectCombatLogDeaths  int32
	expectCombatLogEvents  int32
	expectUnitOrderEvents  int32
	expectPlayer6Name      string
	expectPlayer6Steamid   uint64
}

var testScenarios = map[int64]testScenario{
	1788648401: {
		matchId:         "1788648401",
		replayUrl:       "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1788648401.dem",
		expectGameBuild: 1036,
		debugLevel:      10,
		debugTick:       95573,
	},

	1786687320: {
		matchId:         "1786687320",
		replayUrl:       "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1786687320.dem",
		expectGameBuild: 1033,
	},

	1785937100: {
		matchId:                "1785937100",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1785937100.dem",
		expectGameBuild:        1033,
		expectEntityEvents:     1965109,
		expectCombatLogDamage:  955729,
		expectCombatLogHealing: 33158,
		expectCombatLogDeaths:  1345,
		expectCombatLogEvents:  41529,
		expectUnitOrderEvents:  52359,
		expectPlayer6Name:      "JiimoxD",
		expectPlayer6Steamid:   76561198203594628,
	},
	1785899023: {
		matchId:                "1785899023",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1785899023.dem",
		expectGameBuild:        1033,
		expectEntityEvents:     2419045,
		expectCombatLogDamage:  1803248,
		expectCombatLogHealing: 48337,
		expectCombatLogDeaths:  1954,
		expectCombatLogEvents:  78804,
		expectUnitOrderEvents:  58269,
		expectPlayer6Name:      "+27",
		expectPlayer6Steamid:   76561198063151170,
	},
	1785874713: {
		matchId:                "1785874713",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1785874713.dem",
		expectGameBuild:        1033,
		expectEntityEvents:     1381012,
		expectCombatLogDamage:  513912,
		expectCombatLogHealing: 33359,
		expectCombatLogDeaths:  749,
		expectCombatLogEvents:  21840,
		expectUnitOrderEvents:  40240,
		expectPlayer6Name:      "San-Say",
		expectPlayer6Steamid:   76561198020188611,
	},
	1781640270: {
		matchId:                "1781640270",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1781640270.dem",
		expectGameBuild:        1027,
		expectEntityEvents:     1057562,
		expectCombatLogDamage:  345422,
		expectCombatLogHealing: 25884,
		expectCombatLogDeaths:  645,
		expectCombatLogEvents:  18313,
		expectUnitOrderEvents:  39222,
		expectPlayer6Name:      "GGGGGGGGGG",
		expectPlayer6Steamid:   76561198032710514,
	},
	1764592109: {
		matchId:                "1764592109",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1764592109.dem",
		expectGameBuild:        1017,
		expectEntityEvents:     1827933,
		expectCombatLogDamage:  1008784,
		expectCombatLogHealing: 33784,
		expectCombatLogDeaths:  1631,
		expectCombatLogEvents:  42228,
		expectUnitOrderEvents:  36172,
		expectPlayer6Name:      "Doffo",
		expectPlayer6Steamid:   76561198087353732,
	},
	1763193771: {
		matchId:                "1763193771",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1763193771.dem",
		expectGameBuild:        1016,
		expectEntityEvents:     1203640,
		expectCombatLogDamage:  623594,
		expectCombatLogHealing: 19530,
		expectCombatLogDeaths:  1022,
		expectCombatLogEvents:  24436,
		expectUnitOrderEvents:  31994,
		expectPlayer6Name:      "Monst_er",
		expectPlayer6Steamid:   76561198201328510,
	},
	1763177231: {
		matchId:                "1763177231",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1763177231.dem",
		expectGameBuild:        1016,
		expectEntityEvents:     1221874,
		expectCombatLogDamage:  479350,
		expectCombatLogHealing: 22611,
		expectCombatLogDeaths:  977,
		expectCombatLogEvents:  20043,
		expectUnitOrderEvents:  35975,
		expectPlayer6Name:      "Supercowman",
		expectPlayer6Steamid:   76561198013311415,
	},
	1734886116: {
		matchId:                "1734886116",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1734886116.dem",
		expectGameBuild:        1003,
		expectEntityEvents:     2049211,
		expectCombatLogDamage:  1048805,
		expectCombatLogHealing: 25089,
		expectCombatLogDeaths:  1447,
		expectCombatLogEvents:  42307,
		expectUnitOrderEvents:  59775,
		expectPlayer6Name:      "Eggard",
		expectPlayer6Steamid:   76561197972599979,
	},
	1731962898: {
		matchId:                "1731962898",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1731962898.dem",
		expectGameBuild:        1003,
		expectEntityEvents:     1183267,
		expectCombatLogDamage:  415560,
		expectCombatLogHealing: 20018,
		expectCombatLogDeaths:  690,
		expectCombatLogEvents:  24296,
		expectUnitOrderEvents:  27525,
		expectPlayer6Name:      "Snayp8",
		expectPlayer6Steamid:   76561198047587062,
	},
	1716444111: {
		matchId:                "1716444111",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1716444111.dem",
		expectGameBuild:        995,
		expectEntityEvents:     2854511,
		expectCombatLogDamage:  1398735,
		expectCombatLogHealing: 49659,
		expectCombatLogDeaths:  2169,
		expectCombatLogEvents:  76921,
		expectUnitOrderEvents:  48822,
		expectPlayer6Name:      "GangBang",
		expectPlayer6Steamid:   76561198136700681,
	},
	1712853372: {
		matchId:                "1712853372",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1712853372.dem",
		expectGameBuild:        991,
		expectEntityEvents:     1708696,
		expectCombatLogDamage:  671297,
		expectCombatLogHealing: 23467,
		expectCombatLogDeaths:  1099,
		expectCombatLogEvents:  30381,
		expectUnitOrderEvents:  48107,
		expectPlayer6Name:      "BFG",
		expectPlayer6Steamid:   76561198047707927,
	},
	1648457986: {
		matchId:                "1648457986",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1648457986.dem",
		expectGameBuild:        962,
		expectEntityEvents:     742252,
		expectCombatLogDamage:  224773,
		expectCombatLogHealing: 5914,
		expectCombatLogDeaths:  466,
		expectCombatLogEvents:  10170,
		expectUnitOrderEvents:  17822,
		expectPlayer6Name:      "Grinder",
		expectPlayer6Steamid:   76561198207988337,
	},
	1605340040: {
		matchId:                "1605340040",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1605340040.dem",
		expectGameBuild:        955,
		expectEntityEvents:     1283574,
		expectCombatLogDamage:  522367,
		expectCombatLogHealing: 31721,
		expectCombatLogDeaths:  795,
		expectCombatLogEvents:  21116,
		expectUnitOrderEvents:  40669,
		expectPlayer6Name:      "Az ★ | Big A Man",
		expectPlayer6Steamid:   76561198156504817,
	},
	1582611189: {
		matchId:                "1582611189",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1582611189.dem",
		expectGameBuild:        944,
		expectEntityEvents:     1427025,
		expectCombatLogDamage:  599388,
		expectCombatLogHealing: 28576,
		expectCombatLogDeaths:  930,
		expectCombatLogEvents:  23800,
		expectUnitOrderEvents:  40237,
		expectPlayer6Name:      "The13ananaMan",
		expectPlayer6Steamid:   76561198068424443,
	},
	1560315800: {
		matchId:                "1560315800",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560315800.dem",
		expectGameBuild:        928,
		expectEntityEvents:     2781076,
		expectCombatLogDamage:  1332418,
		expectCombatLogHealing: 57874,
		expectCombatLogDeaths:  1645,
		expectCombatLogEvents:  51288,
		expectUnitOrderEvents:  63992,
		expectPlayer6Name:      "ariethebeast",
		expectPlayer6Steamid:   76561198065323776,
	},
	1560294294: {
		matchId:                "1560294294",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560294294.dem",
		expectGameBuild:        928,
		expectEntityEvents:     1611898,
		expectCombatLogDamage:  768154,
		expectCombatLogHealing: 11565,
		expectCombatLogDeaths:  954,
		expectCombatLogEvents:  24535,
		expectUnitOrderEvents:  30657,
		expectPlayer6Name:      "Laslo",
		expectPlayer6Steamid:   76561198034549887,
	},
	1560289528: {
		matchId:                "1560289528",
		replayUrl:              "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560289528.dem",
		expectGameBuild:        928,
		expectEntityEvents:     2270022,
		expectCombatLogDamage:  1180993,
		expectCombatLogHealing: 57511,
		expectCombatLogDeaths:  1449,
		expectCombatLogEvents:  49146,
		expectUnitOrderEvents:  63387,
		expectPlayer6Name:      "It takes a tree to tango",
		expectPlayer6Steamid:   76561197993050562,
	},
}

func (s testScenario) test(t *testing.T) {
	assert := assert.New(t)

	if s.debugTick == 0 {
		debugLevel = s.debugLevel
	}

	defer func() {
		debugLevel = 0
	}()

	buf := mustGetReplayData(s.matchId, s.replayUrl)
	parser, err := NewParser(buf)
	if err != nil {
		t.Errorf("unable to instantiate parser: %s", err)
		return
	}

	if s.skipPacketEntities {
		parser.ProcessPacketEntities = false
	}

	gotFileInfo := false
	gotCombatLogDamage := int32(0)
	gotCombatLogHealing := int32(0)
	gotCombatLogDeaths := int32(0)
	gotCombatLogEvents := int32(0)
	gotUnitOrderEvents := int32(0)
	gotEntityEvents := int32(0)
	gotPlayer6Name := "<Missing>"
	gotPlayer6Steamid := uint64(0)

	if s.debugTick > 0 {
		parser.Callbacks.OnCNETMsg_Tick(func(m *dota.CNETMsg_Tick) error {
			if m.GetTick() >= s.debugTick {
				debugLevel = s.debugLevel
			}
			return nil
		})
	}

	parser.Callbacks.OnCDOTAUserMsg_SpectatorPlayerUnitOrders(func(m *dota.CDOTAUserMsg_SpectatorPlayerUnitOrders) error {
		gotUnitOrderEvents += 1
		return nil
	})

	parser.Callbacks.OnCDemoFileInfo(func(m *dota.CDemoFileInfo) error {
		gotFileInfo = true
		return nil
	})

	parser.OnPacketEntity(func(pe *PacketEntity, pet EntityEventType) error {
		gotEntityEvents += 1

		if pe.ClassName == "CDOTA_PlayerResource" {
			if v, ok := pe.FetchString("m_vecPlayerData.0006.m_iszPlayerName"); ok {
				gotPlayer6Name = v
			} else if v, ok := pe.FetchString("m_iszPlayerNames.0006"); ok {
				gotPlayer6Name = v
			}

			if v, ok := pe.FetchUint64("m_vecPlayerData.0006.m_iPlayerSteamID"); ok {
				gotPlayer6Steamid = v
			} else if v, ok := pe.FetchUint64("m_iPlayerSteamIDs.0006"); ok {
				gotPlayer6Steamid = v
			}
		}

		return nil
	})

	parser.OnGameEvent("dota_combatlog", func(m *GameEvent) error {
		gotCombatLogEvents += 1

		t, err := m.GetInt32("type")
		assert.Nil(err)

		switch dota.DOTA_COMBATLOG_TYPES(t) {
		case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DAMAGE:
			v, err := m.GetInt32("value")
			assert.Nil(err)
			gotCombatLogDamage += v
		case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_DEATH:
			gotCombatLogDeaths += 1
		case dota.DOTA_COMBATLOG_TYPES_DOTA_COMBATLOG_HEAL:
			v, err := m.GetInt32("value")
			assert.Nil(err)
			gotCombatLogHealing += v
		}

		return nil
	})

	err = parser.Start()
	assert.Nil(err, s.matchId)

	assert.True(gotFileInfo)
	assert.Equal(s.expectGameBuild, parser.GameBuild, s.matchId)
	if s.expectEntityEvents > 0 {
		assert.Equal(s.expectEntityEvents, gotEntityEvents, s.matchId)
	}
	if s.expectCombatLogDamage > 0 {
		assert.Equal(s.expectCombatLogDamage, gotCombatLogDamage, s.matchId)
	}
	if s.expectCombatLogHealing > 0 {
		assert.Equal(s.expectCombatLogHealing, gotCombatLogHealing, s.matchId)
	}
	if s.expectCombatLogDeaths > 0 {
		assert.Equal(s.expectCombatLogDeaths, gotCombatLogDeaths, s.matchId)
	}
	if s.expectCombatLogEvents > 0 {
		assert.Equal(s.expectCombatLogEvents, gotCombatLogEvents, s.matchId)
	}
	if s.expectUnitOrderEvents > 0 {
		assert.Equal(s.expectUnitOrderEvents, gotUnitOrderEvents, s.matchId)
	}
	if s.expectPlayer6Name != "" {
		assert.Equal(s.expectPlayer6Name, gotPlayer6Name, s.matchId)
	}
	if s.expectPlayer6Steamid > 0 {
		assert.Equal(s.expectPlayer6Steamid, gotPlayer6Steamid, s.matchId)
	}
}

func BenchmarkParseMatch(b *testing.B) {
	assert := assert.New(b)

	buf := mustGetReplayData("1560315800", "https://s3-us-west-2.amazonaws.com/manta.dotabuff/1560315800.dem")

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		parser, err := NewParser(buf)
		assert.Nil(err)
		err = parser.Start()
		assert.Nil(err)
	}

	b.ReportAllocs()
}
