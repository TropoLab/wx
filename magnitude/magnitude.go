package magnitude

import (
	"math"
)

// 1.2. METEOROLOGICAL CONVENTIONS
// Although the Earth is approximately spherical, you need not always use spherical coordinates. For the weather at a point or in a small region such as a town, state, or province, you can use local right-hand Cartesian (rectangular) coordinates, as sketched in Fig. 1.1. Usually, this coordinate system is aligned with x pointing east, y pointing north, and z point- ing up. Other orientations are sometimes used.
// Velocity components U, V, and W correspond to motion in the x, y, and z directions. For example, a positive value of U is a velocity component from west to east, while negative is from east to west. Similarly, V is positive northward, and W is positive upward (Fig. 1.1).
// In polar coordinates, horizontal velocities can be expressed as a direction (α), and speed or magni- tude (M). Historically, horizontal wind directions are based on the compass, with 0° to the north (the positive y direction), and with degrees increasing in a clockwise direction through 360°. Negative angles are not usually used. Unfortunately, this dif- fers from the usual mathematical convention of 0° in the x direction, increasing counter-clockwise through 360° (Fig. 1.2).
// Historically winds are named by the direction from which they come, while in mathematics an- gles give the direction toward which things move. Thus, a west wind is a wind from the west; namely, from 270°. It corresponds to a positive value of U, with air moving in the positive x direction.
// Because of these differences, the usual trigono- metric equations cannot be used to convert between (U, V) and (α, M). Use the following equations in- stead, where α is the compass direction from which winds come.

func magnitude(U float64, V float64) float64 {

	M := math.Pow(math.Pow(U, 2)+math.Pow(V, 2), 1/2)

	return M
}
