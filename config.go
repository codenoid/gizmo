//=============================================================
// config.go
//-------------------------------------------------------------
// Configuration for the game. Loaded both from config file
// and static configuration.
//=============================================================
package main

import (
	"encoding/json"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"os"
)

//=============================================================
// Game constants
//=============================================================
// General constants
const (
	GameTitle   = "GoD - Game or Die"
	GameVersion = "0.0"
)

var gSoundFiles = []string{
	"jump.wav",
	"shot.mp3",
	"shot2.mp3",
	"explosion.mp3",
	"shot3.mp3",
}

// World constants
const (
	wMaxInvFPS           = 1 / 60.0
	wShadowLength        = 5
	wShadowDepth         = 1.5
	wPixelsPerChunk      = 64
	wPixelSize           = 1
	wBorderSize          = 4
	wStaticBorderSize    = 0
	wStaticColor32       = 0xFFFFFFFE
	wStaticColor8        = 0xFE
	wFloodFill8          = 0xFC
	wFloodFill32         = 0xFFFFFFFC
	wFloodFillVisited8   = 0xFB
	wFloodFillVisited32  = 0xFFFFFFFB
	wBackground32        = 0xFFFFFF8F
	wBackground8         = 0x8F
	wBackgroundNew32     = 0xFFFFFF9F // New background when updated (so we don't have to update bg sprite)
	wBackgroundNew8      = 0x9F
	wLadder8             = 0xAF
	wLadder32            = 0xFFFFFFAF
	wShadow8             = 0xBF
	wShadow32            = 0xFFFFFFBF
	wViewMax             = 400 // 450
	wParticleDefaultLife = 5
	wGravity             = -9.82
	wParticlesMax        = 5000
	wAmmoMax             = 1000
	wDoorLen             = 30
	wDoorHeight          = 40
	wLightsMax           = 10
	wMiddleTextSize      = 22
	wFPSTextSize         = 40
	wDeathScreenText     = "You Died"
	wAssetObjectsPath    = "assets/objects/"
	wAssetMobsPath       = "assets/mobs/"
	wAssetMapsPath       = "assets/maps/"
	wAssetMixedPath      = "assets/mixed/"
	wConfigFile          = "configuration.json"
)

//=============================================================
// Global variables
//=============================================================
type Global struct {
	gWindowHeight   int
	gWindowWidth    int
	gVsync          bool
	gUndecorated    bool
	gWorld          *world
	gCamera         *camera
	gParticleEngine *particleEngine
	gClearColor     pixel.RGBA
	gWin            *pixelgl.Window
	gController     *controller
	gAmmoEngine     *ammoEngine
	gTextures       *textures
	gRand           *fRand
	gPlayer         *mob
	gSounds         *sound
	gUI             *UI
	gMap            *Map
	gMapColor       *mapColor
	gFont           *font
	gMainMenu       *menu
	gOptionsMenu    *menu
	gVariableConfig *variableConfig
	uTime           float32
}

var global = &Global{
	gVariableConfig: &variableConfig{},
	//gWindowHeight:   468,
	//gWindowWidth:    1024,
	//gVsync:          false,
	//gUndecorated:    false,
	gWorld:          &world{},
	gCamera:         &camera{},
	gController:     &controller{},
	gParticleEngine: &particleEngine{},
	gClearColor:     pixel.RGBA{0, 0, 0, 1.0},
	gWin:            &pixelgl.Window{},
	gAmmoEngine:     &ammoEngine{},
	gTextures:       &textures{},
	gRand:           &fRand{},
	gSounds:         &sound{},
	gUI:             &UI{},
	gMap:            &Map{},
	gMapColor:       &mapColor{},
	gFont:           &font{},
	gMainMenu:       &menu{},
	gOptionsMenu:    &menu{},
	gPlayer: &mob{
		graphics: graphics{
			animated:    true,
			sheetFile:   fmt.Sprintf("%v%v", wAssetMobsPath, "player.png"),
			walkFrames:  []int{8, 9, 10, 11, 12, 13, 14},
			idleFrames:  []int{0, 2, 3, 4, 5, 6},
			shootFrames: []int{26},
			jumpFrames:  []int{15, 16, 17, 18, 19, 20},
			climbFrames: []int{1, 7},
			frameWidth:  12.0,
		},
		maxLife: 100.0,
	},
}

type variableConfig struct {
	Vsync             bool `json:"Vsync"`
	Fullscreen        bool `json:"Fullscreen"`
	WindowHeight      int  `json:"WindowHeight"`
	WindowWidth       int  `json:"WindowWidth"`
	UndecoratedWindow bool `json:"UndecoratedWindow"`
	currentMap        int  `json:"CurrentMap"`
	KeyShoot          int  `json:"KeyShoot"`
	KeyJump           int  `json:"KeyJump"`
	KeyLeft           int  `json:"KeyLeft"`
	KeyRight          int  `json:"KeyRight"`
	KeyClimb          int  `json:"KeyClimb"`
	KeyAction         int  `json:"KeyAction"`
	KeyDrop           int  `json:"KeyDrop"`
	KeyPickup         int  `json:"KeyPickup"`
	KeyDuck           int  `json:"KeyDuck"`
}

//=============================================================
// Configuration file loading
//=============================================================
func (v *variableConfig) LoadConfiguration() {
	configFile, err := os.Open(wConfigFile)
	defer configFile.Close()
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(v)
	if err != nil {
		panic("Configuration file not valid.")
	}
}

//=============================================================
// Configuration file loading
//=============================================================
func (v *variableConfig) SaveConfiguration() {
	// TBD
}
