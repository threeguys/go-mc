package blocks

const (
	Air = 0
	Stone = 1
	GrassBlock = 2
	Dirt = 3
	Cobblestone = 4
	WoodPlanks = 5
	Saplings = 6
	Bedrock = 7
	Water = 8
	StationaryWater = 9
	Lava = 10
	StationaryLava = 11
	Sand = 12
	Gravel = 13
	GoldOre = 14
	IronOre = 15
	CoalOre = 16
	Wood = 17
	Leaves = 18
	Sponge = 19
	Glass = 20
	LapisLazuliOre = 21
	LapisLazuliBlock = 22
	Dispenser = 23
	Sandstone = 24
	NoteBlock = 25
	Bed = 26
	PoweredRail = 27
	DetectorRail = 28
	StickyPiston = 29
	Cobweb = 30
	Grass = 31
	DeadBush = 32
	Piston = 33
	PistonExtension = 34
	Wool = 35
	BlockMovedByPiston = 36
	Dandelion = 37
	Rose = 38
	BrownMushroom = 39
	RedMushroom = 40
	BlockOfGold = 41
	BlockOfIron = 42
	DoubleSlabs = 43
	Slabs = 44
	Bricks = 45
	TNT = 46
	Bookshelf = 47
	MossStone = 48
	Obsidian = 49
	Torch = 50
	Fire = 51
	MonsterSpawner = 52
	OakWoodStairs = 53
	Chest = 54
	RedstoneWire = 55
	DiamondOre = 56
	BlockOfDiamond = 57
	CraftingTable = 58
	WheatCrop = 59
	Farmland = 60
	Furnace = 61
	BurningFurnace = 62
	SignPost = 63
	WoodenDoor = 64
	Ladders = 65
	Rail = 66
	CobblestoneStairs = 67
	WallSign = 68
	Lever = 69
	StonePressurePlate = 70
	IronDoor = 71
	WoodenPressurePlate = 72
	RedstoneOre = 73
	GlowingRedstoneOre = 74
	RedstoneTorchInactive = 75
	RedstoneTorchActive = 76
	StoneButton = 77
	Snow = 78
	Ice = 79
	SnowBlock = 80
	Cactus = 81
	Clay = 82
	SugarCane = 83
	Jukebox = 84
	Fence = 85
	Pumpkin = 86
	Netherrack = 87
	SoulSand = 88
	Glowstone = 89
	NetherPortal = 90
	JackOLantern = 91
	CakeBlock = 92
	RedstoneRepeaterInactive = 93
	RedstoneRepeaterActive = 94
	LockedChest = 95
	Trapdoor = 96
	MonsterEgg = 97
	StoneBricks = 98
	HugeBrownMushroom = 99
	HugeRedMushroom = 100
	IronBars = 101
	GlassPane = 102
	Melon = 103
	PumpkinStem = 104
	MelonStem = 105
	Vines = 106
	FenceGate = 107
	BrickStairs = 108
	StoneBrickStairs = 109
	Mycelium = 110
	LilyPad = 111
	NetherBrick = 112
	NetherBrickFence = 113
	NetherBrickStairs = 114
	NetherWart = 115
	EnchantmentTable = 116
	BrewingStand = 117
	Cauldron = 118
	EndPortal = 119
	EndPortalBlock = 120
	EndStone = 121
	DragonEgg = 122
	RedstoneLampInactive = 123
	RedstoneLampActive = 124
	WoodenDoubleSlab = 125
	WoodenSlab = 126
	Cocoa = 127
	SandstoneStairs = 128
	EmeraldOre = 129
	EnderChest = 130
	TripwireHook = 131
	Tripwire = 132
	BlockOfEmerald = 133
	SpruceWoodStairs = 134
	BirchWoodStairs = 135
	JungleWoodStairs = 136
	CommandBlock = 137
	Beacon = 138
	CobblestoneWall = 139
	FlowerPot = 140
	Carrots = 141
	Potatoes = 142
	WoodenButton = 143
	MobHead = 144
	Anvil = 145
	TrappedChest = 146
	WeightedPressurePlateLight = 147
	WeightedPressurePlateHeavy = 148
	RedstoneComparatorInactive = 149
	RedstoneComparatorActive = 150
	DaylightSensor = 151
	BlockOfRedstone = 152
	NetherQuartzOre = 153
	Hopper = 154
	BlockOfQuartz = 155
	QuartzStairs = 156
	ActivatorRail = 157
	Dropper = 158
	StainedClay = 159
	HayBlock = 170
	Carpet = 171
	HardenedClay = 172
	BlockOfCoal = 173
)

type Block struct {
	id uint16
	data, light, skyLight byte
}

type Section struct {
	yIndex int8
	blocks [4096]*Block
}

type TileTick struct {
}

type TileEntity struct {
	id int32
	x, y, z float64
}
