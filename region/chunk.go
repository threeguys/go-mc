package region

type Chunk struct {
	x, z int
	lastUpdate int
	populated bool
	biomes [256]byte
	heightMap [256]int32
	sections []*Section
	entities []*Entity
	tileEntities []*TileEntity
	tileTicks []*TileTick
}