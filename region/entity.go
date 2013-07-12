package region

type Entity struct {
	id int32
	x, y, z float64
	dX, dY, dZ float64
	yaw, pitch float32
	fallDistance int32
	fire int32
	air uint8
	onGround bool
	invulnerable bool
	portalCooldown int16
	uuid [128]byte
}