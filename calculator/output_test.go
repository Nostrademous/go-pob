package calculator

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"

	"go-pob/builds"
)

func TestOutput(t *testing.T) {
	file, err := os.ReadFile("../testdata/builds/Fireball.xml")
	testza.AssertNoError(t, err)

	build, err := builds.ParseBuild(file)
	testza.AssertNoError(t, err)

	calculator := &Calculator{PoB: build}
	env := calculator.BuildOutput(OutputModeMain)

	testza.AssertEqual(t, 0.9523809523809523, env.Player.OutputTable[OutTableMainHand]["TotalMin"])
	testza.AssertEqual(t, 2.8571428571428568, env.Player.OutputTable[OutTableMainHand]["TotalMax"])
	testza.AssertEqual(t, 1.9047619047619047, env.Player.OutputTable[OutTableMainHand]["AverageHit"])
	testza.AssertEqual(t, 1.8857142857142855, env.Player.OutputTable[OutTableMainHand]["AverageDamage"])
	testza.AssertEqual(t, 2.2628571428571425, env.Player.OutputTable[OutTableMainHand]["TotalDPS"])

	marshal, err := json.MarshalIndent(env, "", "  ")
	testza.AssertNoError(t, err)

	err = os.WriteFile("out.json", marshal, 0777)
	testza.AssertNoError(t, err)
}
