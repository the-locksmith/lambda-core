package bsp

import (
	"math"
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/primitives/face"
	"github.com/galaco/bsp/lumps"
	"github.com/go-gl/mathgl/mgl32"
)

func GenerateFacesFromBSP(file *bsp.Bsp) ([]float32, [][]uint16) {
	var verts []float32
	var expFaces [][]uint16

	fl := *file.GetLump(bsp.LUMP_FACES).GetContents()
	faces := (fl).(lumps.Face).GetData().(*[]face.Face)

	vs := *file.GetLump(bsp.LUMP_VERTEXES).GetContents()
	vertexes := (vs).(lumps.Vertex).GetData().(*[]mgl32.Vec3)

	sf := *file.GetLump(bsp.LUMP_SURFEDGES).GetContents()
	surfEdges := (sf).(lumps.Surfedge).GetData().(*[]int32)

	ed := *file.GetLump(bsp.LUMP_EDGES).GetContents()
	edges := (ed).(lumps.Edge).GetData().(*[][2]uint16)


	for _,v := range *vertexes {
		verts = append(verts, v.X(), v.Y(), v.Z())
	}

	for _,f := range *faces {
		expF := make([]uint16, 0)
		//// Just so we render triangles

		// All surfedges associated with this face
		// surfEdges are basically indices into the edges lump
		surfEdges := (*surfEdges)[f.FirstEdge:(f.FirstEdge+int32(f.NumEdges))]
		for _,surfEdge := range surfEdges {
			edge := (*edges)[int(math.Abs(float64(surfEdge)))]

			// Fix reverse ordering
			if surfEdge >= 0 {
				expF = append(expF, edge[0], edge[1])
			} else {
				//negative surfedge reverses edge order
				expF = append(expF, edge[1], edge[0])
			}
		}

		expFaces = append(expFaces, expF)
	}

	return verts, expFaces
}
