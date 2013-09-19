/*
Package collada implements a schema for importing and exporting collada (.dea) documents

Collada 1.5 Specification
http://www.khronos.org/files/collada_spec_1_5.pdf

Only collada 1.5 will be supported

As of this release only a subset of the total schema is interpreted and validation tests should
be included before this package is considered production read

This collada package aims to provide strongly typed access to a collada document without
any extra indexing or processing properties, some common indexing will be provided by collada/util

*/
package collada

import (
	"bufio"
	"encoding/xml"
	"io"
	"os"
)

type Version string

const (
	Version1_5_0 Version = "1.5.0"
)

type Uri string

type UpAxis string

const (
	Xup UpAxis = "X_UP"
	Yup UpAxis = "Y_UP"
	Zup UpAxis = "Z_UP"
)

//Animation ategorizes the declaration of animation information.
type Animation struct {
	//TODO
}

//AnimationClip defines a section of the animation curves to be used together as an animation clip.
type AnimationClip struct {
	//TODO
}

//Channel declares an output channel of an animation.
type Channel struct {
	//TODO
}

//InstanceAnimation instantiates a COLLADA animation resource.
type InstanceAnimation struct {
	//TODO
}

//LibraryAnimationClips provides a library in which to place <animation_clip> elements.
type LibraryAnimationClips struct {
	//TODO
}

//LibraryAnimations provides a library in which to place <animation> elements.
type LibraryAnimations struct {
	//TODO
}

//Sampler declares an interpolation sampling function for an animation.
type Sampler struct {
	//TODO
}

//Camera declares a view into the scene hierarchy or scene graph.
//The camera contains elements that describe the camera’s optics and imager.
type Camera struct {
	HasId
	HasName
	HasAsset
	Optics Optics  `xml:"optics"`
	Imager *Imager `xml:"imager"`
	HasExtra
}

//Imager represents the image sensor of a camera (for example, film or CCD).
type Imager struct {
	HasTechnique
	HasExtra
}

//InstanceCamera instantiates a COLLADA camera resource.
type InstanceCamera struct {
	HasSid
	HasName
	HasUrl
	HasExtra
}

//Provides a library in which to place <camera> elements.
type LibraryCameras struct {
	HasId
	HasName
	HasAsset
	Camera []Camera `xml:"camera"`
	HasExtra
}

//Optics represents the apparatus on a camera that projects the image onto the image sensor.
type Optics struct {
	HasTechniqueCommon
	HasTechnique
	HasExtra
}

//Orthographic describes the field of view of an orthographic camera.
type Orthographic struct {
	//TODO
}

//Perspective describes the field of view of a perspective camera.
type Perspective struct {
	//TODO
}

//Controller categorizes the declaration of generic control information.
type Controller struct {
	//TODO
}

//InstanceController instantiates a a COLLADA controller resource.
type InstanceController struct {
	HasSid
	HasName
	HasUrl
	Skeleton     []Skeleton    `xml:"skeleton"`
	BindMaterial *BindMaterial `xml:"bind_material"`
	HasExtra
}

//BindMaterial binds a specific material to a piece of geometry, binding varying and uniform parameters at the same time.
type BindMaterial struct {
	Param []ParamCore `xml:"param"`
	HasTechniqueCommon
	HasTechnique
	HasExtra
}

//Joints associates joint, or skeleton, nodes with attribute data.
type Joints struct {
	//TODO
}

//LibraryControllers provides a library in which to place <controller> elements.
type LibraryControllers struct {
	//TODO
}

//Morph describes the data required to blend between sets of static meshes.
type Morph struct {
	//TODO
}

//Skeleton indicates where a skin controller is to start searching for the joint nodes that it needs.
type Skeleton struct {
	//TODO
}

//Skin contains vertex and primitive information sufficient to describe blend-weight skinning.
type Skin struct {
	//TODO
}

// Targets teclares morph targets, their weights, and any user-defined attributes associated with them.
type Targets struct {
	//TODO
}

// VertexWeights describes the combination of joints and weights used by a skin.
type VertexWeights struct {
	//TODO
}

// Accessor declares an access pattern to one of the array elements <float_array>, <int_array>, <Name_array>, <bool_array>, and <IDREF_array>.
type Accessor struct {
	//TODO
}

// BoolArray declares the storage for a homogenous array of Boolean values.
type BoolArray struct {
	HasCount
	HasId
	HasName
	Bools
}

// FloatArray declares the storage for a homogenous array of floating-point values.
type FloatArray struct {
	HasCount
	HasId
	HasName
	Digits    uint8  `xml:"digits,attr"`
	Magnitude uint16 `xml:"magnitude,attr"`
	Floats
}

// IdRefArray declares the storage for a homogenous array of ID reference values.
type IdRefArray struct {
	HasCount
	HasId
	HasName
	IdRefs
}

// IntArray stores a homogenous array of integer values.
type IntArray struct {
	HasCount
	HasId
	HasName
	MinInclusive *int `xml:"minInclusive,attr"`
	MaxInclusive *int `xml:"maxInclusive,attr"`
	Ints
}

// NameArray stores a homogenous array of symbolic name values.
type NameArray struct {
	HasCount
	HasId
	HasName
	Names
}

// ParamCore declares parametric information for its parent element.
type ParamCore struct {
	//TODO
}

// SidRefArray declares the storage for a homogenous array of scoped-identifier reference values.
type SidRefArray struct {
	HasCount
	HasId
	HasName
	SidRefs
}

// Source declares a data repository that provides values according to
// the semantics of an <input> element that refers to it.
type Source struct {
	HasId
	HasName
	HasAsset
	BoolArray   *BoolArray   `xml:"bool_array"`
	FloatArray  *FloatArray  `xml:"float_array"`
	IdRefArray  *IdRefArray  `xml:"IDREF_array"`
	IntArray    *IntArray    `xml:"int_array"`
	NameArray   *NameArray   `xml:"Name_array"`
	SidRefArray *SidRefArray `xml:"SIDREF_array"`
	// TokenArray *TokenArray `xml:"token_array"`
	HasTechniqueCommon
	HasTechnique
}

// InputShared declares the input semantics of a data source.
type InputShared struct {
	Offset   uint   `xml:"offset,attr"`
	Semantic string `xml:"semantic,attr"`
	Soruce   string `xml:"source,attr"`
	Set      uint   `xml:"set,attr,omitempty"`
}

// InputUnshared declares the input semantics of a data source.
type InputUnshared struct {
	Semantic string `xml:"semantic,attr"`
	Source   string `xml:"source,attr"`
}

// Extra provides arbitrary additional information about or related to its parent element.
type Extra struct {
	HasId
	HasName
	HasType
	HasAsset
	HasTechnique
}

//Technique (core) Declares the information used to process some portion of the content. Each technique conforms to an associated profile.
type TechniqueCore struct {
	Profile string `xml:"profile,attr"`
	Xmlns   string `xml:"xmlns,attr"`
	XML     string `xml:",innerxml"`
}

//TechniqueCommon specifies the information for a specific element for the common profile that all COLLADA implementations must support.
type TechniqueCommon struct {
	XML string `xml:",innerxml"`
}

//ControlVertices describes the control vertices (CVs) of a spline.
type ControlVertices struct {
	//TODO
}

//Geometry describes the visual shape and appearance of an object in a scene.
type Geometry struct {
	HasId
	HasName
	HasAsset
	//TODO
	// ConvexMesh *ConvexMesh `xml:"convex_mesh"`
	Mesh   *Mesh   `xml:"mesh"`
	Spline *Spline `xml:"spline"`
	//TODO
	// Brep *Brep `xml:"brep"`
	HasExtra
}

//InstanceGeometry instantiates a COLLADA geometry resource.
type InstanceGeometry struct {
	HasSid
	HasName
	HasUrl
	BindMaterial *BindMaterial `xml:"bind_material"`
	HasExtra
}

//LibraryGeometries provides a library in which to place <geometry> elements.
type LibraryGeometries struct {
	HasId
	HasName
	HasAsset
	Geometry []Geometry `xml:"geometry"`
	HasExtra
}

//Lines declares the binding of geometric primitives and vertex attributes for a <mesh>element.
type Lines struct {
	HasName
	HasCount
	HasMaterial
	HasSharedInput
	HasP
	HasExtra
}

//Linestrips declares a binding of geometric primitives and vertex attributes for a <mesh>element.
type Linestrips struct {
	HasName
	HasCount
	HasMaterial
	HasSharedInput
	HasPs
	HasExtra
}

//Mesh describes basic geometric meshes using vertex and primitive information.
type Mesh struct {
	Source     []Source     `xml:"source"`
	Vertices   Vertices     `xml:"vertices"`
	Lines      []Lines      `xml:"lines"`
	Linestrips []Linestrips `xml:"linestrips"`
	Polygons   []Polygons   `xml:"polygons"`
	Polylist   []Polylist   `xml:"polylist"`
	Triangles  []Triangles  `xml:"triangles"`
	Trifans    []Trifans    `xml:"trifans"`
	Tristrips  []Tristrips  `xml:"tristrips"`
	HasExtra
}

//Polygons declares the binding of geometric primitives and vertex attributes for a <mesh>element.
type Polygons struct {
	HasName
	HasCount
	HasMaterial
	HasSharedInput
	HasPs
	HasPhs
	HasExtra
}
type H Ints
type HasPhs struct {
	Ph []Ph `xml:"ph"`
}
type Ph struct {
	P P   `xml:"p"`
	H []H `xml:"h"`
}

//Polylist declares the binding of geometric primitives and vertex attributes for a <mesh>element.
type Polylist struct {
	HasName
	HasCount
	HasMaterial
	HasSharedInput
	VCount *Ints `xml:"vcount"`
	HasP
	HasExtra
}

//Spline describes a multisegment spline with control vertex (CV) and segment information.
type Spline struct {
	//TODO
}

//Triangles provides the information needed to bind vertex attributes together and then organize those vertices into individual triangles.
type Triangles struct {
	HasName
	HasCount
	HasMaterial
	HasSharedInput
	HasP
	HasExtra
}

//Trifans provides the information needed to bind vertex attributes together and then organize those vertices into connected triangles.
type Trifans struct {
	HasName
	HasCount
	HasMaterial
	HasSharedInput
	HasP
	HasExtra
}

//Tristrips provides the information needed to bind vertex attributes together and then organize those vertices into connected triangles
type Tristrips struct {
	HasName
	HasCount
	HasMaterial
	HasSharedInput
	HasP
	HasExtra
}

//Vertices declares the attributes and identity of mesh-vertices.
type Vertices struct {
	HasId
	HasName
	Input []InputUnshared `xml:"input"`
	HasExtra
}

//AmbientCore (core) Describes an ambient light source.
type AmbientCore struct {
	//TODO
}

//Color describes the color of its parent light element.
type Color struct {
	//TODO
}

//Directional describes a directional light source.
type Directional struct {
	//TODO
}

//InstanceLight instantiates a COLLADA light resource.
type InstanceLight struct {
	HasSid
	HasName
	HasUrl
	HasExtra
}

//LibraryLights provides a library in which to place <image> elements.
type LibraryLights struct {
	HasId
	HasName
	HasAsset
	Light []Light `xml:"light"`
	HasExtra
}

//Light declares a light source that illuminates a scene.
type Light struct {
	HasId
	HasName
	HasAsset
	HasTechniqueCommon
	HasTechnique
	HasExtra
}

//Point describes a point light source.
type Point struct {
	//TODO
}

//Spot describes a spot light source.
type Spot struct {
	//TODO
}

//Formula defines a formula.
type Formula struct {
	//TODO
}

//InstanceFormula instantiates a COLLADA formula resource.
type InstanceFormula struct {
	//TODO
}

//LibraryFormulas provides a library in which to place <formula> elements.
type LibraryFormulas struct {
	//TODO
}

//Coverage provides information about the location of the visual scene in physical space.
type Coverage struct {
	GeographicLocation []GeographicLocation `xml:"geographic_location"`
}

//Unit defines unit of distance for COLLADA elements and objects.
type Unit struct {
	HasName
	Meter float64 `xml:"meter,attr"`
}

//Asset defines asset-management information regarding its parent element.
type Asset struct {
	Contributor []Contributor `xml:"contributor"`
	Coverage    *Coverage     `xml:"coverage"`
	Created     string        `xml:"created"`
	Keywords    string        `xml:"keywords,omitempty"`
	Modified    string        `xml:"modified"`
	Revision    string        `xml:"revision,omitempty"`
	Subject     string        `xml:"subject,omitempty"`
	Title       string        `xml:"title,omitempty"`
	Unit        *Unit         `xml:"unit"`
	UpAxis      UpAxis        `xml:"up_axis,omitempty"`
}

//COLLADA declares the root of the document that contains some of the content in the COLLADA schema.
type Collada struct {
	Version Version `xml:"version,attr"`
	Xmlns   Uri     `xml:"xmlns,attr,omitempty"`
	Base    Uri     `xml:"base,attr,omitempty"`
	HasAsset
	LibraryAnimationClips []LibraryAnimationClips `xml:"library_animation_clips"`
	LibraryAnimations     []LibraryAnimations     `xml:"library_animations"`
	// LibraryArticulatedSystems []LibraryArticulatedSystems `xml:"library_animation_clips"`
	LibraryCameras     []LibraryCameras     `xml:"library_cameras"`
	LibraryControllers []LibraryControllers `xml:"library_controllers"`
	// LibraryEffects []LibraryEffects `xml:"library_effects"`
	// LibraryForceFields []LibraryForceFields `xml:"library_force_fields"`
	LibraryFormulas   []LibraryFormulas   `xml:"library_formulas"`
	LibraryGeometries []LibraryGeometries `xml:"library_geometries"`
	// LibraryImages []LibraryImages `xml:"library_images"`
	// LibraryJoints []LibraryJoints `xml:"library_joints"`
	// LibraryKinematicModels []LibraryKinematicModels `xml:"library_kinematic_models"`
	// LibraryKinematicScenes []LibraryKinematicScenes `xml:"library_kinematic_scenes"`
	LibraryLights []LibraryLights `xml:"library_lights"`
	// LibraryMaterials []LibraryMaterials `xml:"library_materials"`
	// LibraryPhysicsNodes []LibraryPhysicsNodes `xml:"library_physics_nodes"`
	// LibraryPhysicsMaterials []LibraryPhysicsMaterials `xml:"library_physics_materials"`
	// LibraryPhysicsScenes []LibraryPhysicsScenes `xml:"library_physics_scenes"`
	// LibraryPhysicsScenes []LibraryPhysicsScenes `xml:"library_physics_scenes"`
	LibraryVisualScenes []LibraryVisualScenes `xml:"library_visual_scenes"`
	Scene               *Scene                `xml:"scene"`
	HasExtra
}

//Contributor defines authoring information for asset management.
type Contributor struct {
	Author        string `xml:"author,omitempty"`
	AuthorEmail   string `xml:"author_email,omitempty"`
	AuthorWebsite string `xml:"author_website,omitempty"`
	AuthoringTool string `xml:"authoring_tool,omitempty"`
	Comments      string `xml:"comments,omitempty"`
	Copyright     string `xml:"copyright,omitempty"`
	SourceData    Uri    `xml:"source_data,omitempty"`
}

//GeographicLocation defines an asset’s location for asset management.
type GeographicLocation struct {
	//TODO
}

//Newparam creates a new, named parameter object and assigns it a type and an initial value.
type Newparam struct {
	//TODO
}

//ParamReference references a predefined parameter.
type ParamReference struct {
	//TODO
}

//Setparam assigns a new value to a previously defined parameter.
type Setparam struct {
	//TODO
}

//EvaluateScene declares information specifying how to evaluate a <visual_scene>.
type EvaluateScene struct {
	//TODO
}

//InstanceNode instantiates a COLLADA node resource.
type InstanceNode struct {
	HasSid
	HasName
	HasUrl
	Proxy Uri `xml:"proxy,attr,omitempty"`
	HasExtra
}

//InstanceVisualScene instantiates a COLLADA visual_scene resource.
type InstanceVisualScene struct {
	HasSid
	HasName
	HasUrl
	HasExtra
}

//LibraryNodes provides a library in which to place <node> elements.
type LibraryNodes struct {
	//TODO
}

//LibraryVisualScenes provides a library in which to place <visual_scene> elements.
type LibraryVisualScenes struct {
	HasId
	HasName
	HasAsset
	VisualScene []VisualScene `xml:"visual_scene"`
	HasExtra
}

//Node embodies the hierarchical relationship of elements in a scene.
type Node struct {
	HasId
	HasName
	HasSid
	HasType
	Layer string `xml:"layer,attr,omitempty"`
	HasAsset
	Lookat             []Lookat             `xml:"lookat"`
	Matrix             []Matrix             `xml:"matrix"`
	Rotate             []Rotate             `xml:"rotate"`
	Scale              []Scale              `xml:"scale"`
	Skew               []Skew               `xml:"skew"`
	Translate          []Translate          `xml:"translate"`
	InstanceCamera     []InstanceCamera     `xml:"instance_camera"`
	InstanceController []InstanceController `xml:"instance_controller"`
	InstanceGeometry   []InstanceGeometry   `xml:"instance_geometry"`
	InstanceLight      []InstanceLight      `xml:"instance_light"`
	InstanceNode       []InstanceNode       `xml:"instance_node"`
	HasNodes
	HasExtra
}

//Scene embodies the entire set of information that can be visualized from the contents of a COLLADA resource.
type Scene struct {
	InstancePhysicsScene    []InstancePhysicsScene   `xml:"instance_physics_scene"`
	InstanceVisualScene     *InstanceVisualScene     `xml:"instance_visual_scene"`
	InstanceKinematicsScene *InstanceKinematicsScene `xml:"instance_kinematics_scene"`
	HasExtra
}
type InstancePhysicsScene struct {
	//TODO
}
type InstanceKinematicsScene struct {
	//TODO
}

//VisualScene embodies the entire set of information that can be visualized from the contents of a COLLADA resource.
type VisualScene struct {
	HasId
	HasName
	HasAsset
	HasNodes
	EvaluateScene []EvaluateScene `xml:"evaluate_scene"`
	HasExtra
}

//Lookat contains a position and orientation transformation suitable for aiming a camera.
type Lookat struct {
	HasSid
	Float3x3
}

//Matrix describes transformations that embody mathematical changes to points within a coordinate system or the coordinate system itself.
type Matrix struct {
	HasSid
	Float4x4
}

//Rotate specifies how to rotate an object around an axis.
type Rotate struct {
	HasSid
	Float4
}

//Scale specifies how to change an object’s size.
type Scale struct {
	HasSid
	Float3
}

//Skew specifies how to deform an object along one axis.
type Skew struct {
	HasSid
	Float3
}

//Translate changes the position of an object in a coordinate system without any rotation.
type Translate struct {
	HasSid
	Float3
}

type P Ints

type Floats Values
type Bools Values
type IdRefs Values
type Ints Values
type Names Values
type SidRefs Values
type Tokens Values

type Values struct {
	V string `xml:",chardata"`
}

type Float3x3 Floats
type Float4x4 Floats
type Float4 Floats
type Float3 Floats

type HasSharedInput struct {
	Input []InputShared `xml:"input"`
}
type HasMaterial struct {
	Material string `xml:"material,attr,omitempty"`
}
type HasName struct {
	Name string `xml:"name,attr,omitempty"`
}
type HasCount struct {
	Count int `xml:"count,attr,omitempty"`
}
type HasType struct {
	Type string `xml:"type,attr,omitempty"`
}
type HasId struct {
	Id string `xml:"id,attr,omitempty"`
}
type HasUrl struct {
	Url Uri `xml:"url,attr,omitempty"`
}
type HasSid struct {
	Id string `xml:"sid,attr,omitempty"`
}
type HasAsset struct {
	Asset *Asset `xml:"asset,omitempty"`
}
type HasNodes struct {
	Node []*Node `xml:"node"`
}
type HasExtra struct {
	Extra []Extra `xml:"extra"`
}
type HasTechniqueCommon struct {
	TechniqueCommon TechniqueCommon `xml:"technique_common"`
}
type HasTechnique struct {
	TechniqueCore []TechniqueCore `xml:"technique,omitempty"`
}
type HasP struct {
	P *P `xml:"p"`
}
type HasPs struct {
	P []P `xml:"p"`
}

//
// type LibraryCameras struct {
//     HasId
//     HasName
//     HasAsset
//     Cameras []Camera `xml:"camera"`
//     HasExtra
// }
//
// type Camera struct {
//     HasId
//     HasName
//     HasAsset
//     Optics Optics `xml:"optics"`
//     Imager *Imager `xml:"imager"`
//     HasExtra
// }
//
// type Optics struct {
//     TechniqueCommon TechniqueCommon `xml:"technique_common"`
//     Technique []Technique `xml:"technique,omitempty"`
//     HasExtra
// }
//
// type Imager struct {
//     //TODO
// }
//
// type LibraryLights struct {
//     HasId
//     HasName
//     HasAsset
//     Lights []Light `xml:"light"`
//     HasExtra
// }
//
// type Light struct {
//     HasId
//     HasName
//     HasAsset
//     TechniqueCommon TechniqueCommon `xml:"technique_common"`
//     Techniques []Technique `xml:"technique"`
//     HasExtra
// }
//
// type LibraryEffects struct {
//     HasId
//     HasName
//     HasAsset
//     Effects []Effect `xml:"effect"`
//     HasExtra
// }
//
// type Effect struct {
//     HasId
//     HasName
//     HasAsset
//     // Annotations []Annotate `xml:"annotate"`
//     Images []Image `xml:"image,omitempty"`
//     // NewParams []NewParam `xml:"newparam"`
//     ProfileCommon *ProfileCommon `xml:"profile_COMMON,omitempty"`
//     ProfileGLES *ProfileGLES `xml:"profile_GLES,omitempty"`
//     ProfileCG *ProfileCG `xml:"profile_CG,omitempty"`
//     ProfileGLSL *ProfileGLSL `xml:"profile_GLSL,omitempty"`
//     HasExtra
// }
//
// type ProfileCommon struct {
//     XML string `xml:",innerxml"`
// }
// type ProfileCG struct {
//     XML string `xml:",innerxml"`
// }
// type ProfileGLES struct {
//     XML string `xml:",innerxml"`
// }
// type ProfileGLSL struct {
//     XML string `xml:",innerxml"`
// }
//
// type Annotate struct {
//
// }
//
// type LibraryImages struct {
//     HasId
//     HasName
//     HasAsset
//     Images []Image `xml:"image"`
//     HasExtra
// }
//
// type Image struct {
//     HasId
//     HasName
//     Format string `xml:"token,attr"`
//     Height uint `xlm:"height,attr"`
//     Width uint `xml:"width,attr"`
//     Depth uint `xml:"depth,attr"`
//     HasAsset
//     Data string `xml:"data,chardata"`
//     InitFrom string `xml:"init_from"`
//     HasExtra
// }
//
// type LibraryMaterials struct {
//     HasId
//     HasName
//     HasAsset
//     Materials []Material `xml:"material"`
//     HasExtra
// }
//
// type Material struct {
//     HasId
//     HasName
//     HasAsset
//     InstanceEffect InstanceEffect `xml:"instance_effect"`
//     HasExtra
// }
//
// type InstanceEffect struct {
//     HasUrl
//     HasSid
//     HasName
//     //TODO
//     HasExtra
// }
//
// type LibraryVisualScenes struct {
//     Scenes []VisualScene `xml:"visual_scene"`
// }
//
// type LibraryGeometries struct {
//     HasId
//     HasName
//     HasAsset
//     HasExtra
//     Geometries []Geometry `xml:"geometry"`
// }
//
// type Geometry struct {
//     HasId
//     HasName
//     HasAsset
//     Meshes []Mesh `xml:"mesh"`
//     ConvexMeshes []ConvexMesh `xml:"convex_mesh"`
//     Splies []Spline `xml:"spline"`
//     HasExtra
// }
//
// type Mesh struct {
//     Sources []Source `xml:"source"`
//     Vertices Vertices `xml:"vertices"`
//     Lines []Lines `xml:"lines"`
//     LineStrips []LineStrips `xml:"linestrips"`
//     Polygons []Polygons `xml:"polygons"`
//     Polylists []Polylist `xml:"polylist"`
//     Triangles []Triangles `xml:"triangles"`
//     Trifans []Trifans `xml:"trifans"`
//     Tristrips []Tristrips `xml:"tristrips"`
//     HasExtra
// }
//
// type Vertices struct {
//     HasId
//     HasName
//     HasExtra
//     Inputs []Input `xml:"input"`
// }
//
// type Input struct {
//     Offset uint `xml:"offset,attr,omitempty"`
//     Semantic string `xml:"semantic,attr,omitempty"`
//     Source string `xml:"source,attr,omitempty"`
//     Set uint `xml:"set,attr,omitempty"`
// }
//
// type BasicGeometry struct {
//     HasName
//     Count uint `xml:"count,attr"`
//     Material string `xml:"material,attr"`
//     HasExtra
//     Inputs []Input `xml:"input"`
//     P string `xml:"p"`
// }
//
// type Lines BasicGeometry
// type LineStrips BasicGeometry
// type Triangles BasicGeometry
// type Trifans BasicGeometry
// type Tristrips BasicGeometry
//
// type Polylist struct {
//     BasicGeometry
//     VCount string `xml:"vcount"`
// }
//
// type Polygons struct {
//     BasicGeometry
//     Ph Ph `xml:"ph"`
// }
//
// type Ph struct {
//     P string `xml:"p"`
//     H string `xml:"h"`
// }
//
// type Source struct {
//     HasId
//     HasName
//     HasAsset
//     IdRef *IdRefArray `xml:"IDREF_array"`
//     Name *NameArray `xml:"Name_array"`
//     Bool *BoolArray `xml:"bool_array"`
//     Float *FloatArray `xml:"float_array"`
//     Int *IntArray `xml:"int_array"`
//     TechniqueCommon TechniqueCommon `xml:"technique_common"`
//     Techniques []Technique `xml:"technique"`
// }
//
// type ValueArray struct {
//     HasId
//     HasName
//     Count uint `xml:"count,attr"`
//     Values string `xml:",chardata"`
// }
//
// type IdRefArray ValueArray
//
// type IntArray struct {
//     ValueArray
//     MinInclusive int64 `xml:"minInclusive,attr,omitempty"`
//     MaxInclusive int64 `xml:"maxInclusive,attr,omitempty"`
// }
//
// type FloatArray struct {
//     ValueArray
//     Digits int `xml:"digits,attr,omitempty"`
//     Magnitude int `xml:"magnitude,attr,omitempty"`
// }
//
// type BoolArray ValueArray
// type NameArray ValueArray
//
// type ConvexMesh struct {
//     Mesh
//     ConvexHullOf string `xml:"convex_hull_of"`
// }
//
// type Spline struct {
//     //TODO
// }
//
// type Asset struct {
//     Contributor     []Contributor `xml:"contributor,omitempty"`
//     Created         string `xml:"created,omitempty"`
//     Keywords        string `xml:"keywords,omitempty"`
//     Modified        string `xml:"modified,omitempty"`
//     Revision        string `xml:"revision,omitempty"`
//     Subject         string `xml:"subject,omitempty"`
//     Title           string `xml:"title,omitempty"`
//     Unit            string `xml:"unit,omitempty"`
//     UpAxis          string `xml:"up_axis,omitempty"`
// }
//
// type Contributor struct {
//     Author          string `xml:"author,omitempty"`
//     AuthoringTool   string `xml:"authoring_tool,omitempty"`
//     Comments        string `xml:"comments,omitempty"`
//     Copyright       string `xml:"copyright,omitempty"`
//     SourceData      string `xml:"sourceData,omitempty"`
// }
//
// type Unit struct {
//     HasName
//     Meter float64   `xml:"meter,attr"`
// }
//
// type Extra struct {
//     HasId
//     HasName
//     HasType
//     HasAsset
//     Techniques []Technique `xml:"technique"`
// }
//
// type Scene struct {
//     HasExtra
//     // Physics InstanceWithExtra    `xml:"instance_physics_scene"`
//     Visual InstanceWithExtra      `xml:"instance_visual_scene"`
// }
//
// type VisualScene struct {
//     HasId
//     HasName
//     HasAsset
//     HasExtra
//     HasNodes
//     Evaluate []EvaluateScene    `xml:"evaluate_scene"`
// }
//
// type Node struct {
//     HasId
//     HasName
//     HasSid
//     HasType
//     Layer   string    `xml:"layer,attr"`
//     HasAsset
//     HasNodes
//     HasExtra
//     InstanceCamera  []InstanceWithExtra   `xml:"instance_camera"`
//     InstanceController []InstanceController   `xml:"instance_controller"`
//     InstanceGeometry []InstanceGeometry   `xml:"instance_geometry"`
//     InstanceLight   []InstanceWithExtra   `xml:"instance_light"`
//     InstanceNode    []InstanceWithExtra   `xml:"instance_node"`
//
//     LookAt  []LookAt      `xml:"lookat"`
//     Matrix  []Matrix      `xml:"matrix"`
//     Rotate  []Rotate      `xml:"rotate"`
//     Scale   []Scale       `xml:"scale"`
//     Skew    []Skew        `xml:"skew"`
//     Translate []Translate `xml:"translate"`
// }
//
// type Translate V3
// type Rotate V3
// type Scale V3
// type Skew V3
// type LookAt M3
// type Matrix M4
// type Uri string
//
// type InstanceWithExtra struct {
//     HasSid
//     HasName
//     HasUrl
//     HasExtra
// }
//
// type InstanceController struct {
//     //Skeleton
//     //BindMaterial
//     HasExtra
// }
//
// type InstanceGeometry struct {
//     HasUrl
//     HasExtra
//     BindMaterial *BindMaterial `xml:"bind_material"`
// }
//
// type BindMaterial struct {
//     HasExtra
//     TechniqueCommon []TechniqueCommon `xml:"technique_common"`
//     Technique []Technique `xml:"technique"`
// }
//
// type Technique struct {
//     Profile string `xml:"profile,attr"`
//     Content string `xml:",innerxml"`
// }
//
// type TechniqueCommon struct {
//     Content string `xml:",innerxml"`
// }
//
// type InstanceMaterial struct {
//     HasSid
//     HasName
//     HasExtra
//     Symbol string `xml:"symbol"`
//     Target string `xml:"target"`
//     // Bind []Bind `xml:"bind"`
//     // BindVertexInput `xml:"bind_vertex_input"`
// }
//
// type EvaluateScene struct {
//     //TODO
// }
//
//
// type HasSid struct {
//     Sid string `xml:"sid,attr,omitempty"`
// }
//
// type HasName struct {
//     Name string `xml:"name,attr,omitempty"`
// }
//
// type HasUrl struct {
//     Url string `xml:"url,attr,omitempty"`
// }
//
// type V3 struct {
//     HasSid
//     V string `xml:",chardata"`
// }
//
// type V4 struct {
//     HasSid
//     V string `xml:",chardata"`
// }
//
// type M3 struct {
//     HasSid
//     V string `xml:",chardata"`
// }
//
// type M4 struct {
//     HasSid
//     V string `xml:",chardata"`
// }
//
// type HasId struct {
//     Id string `xml:"id,attr,omitempty"`
// }
//
// type HasAsset struct {
//     Asset *Asset `xml:"asset,omitempty"`
// }
//
// type HasExtra struct {
//     Extras []Extra `xml:"extra,omitempty"`
// }
//
// type HasNodes struct {
//     Nodes []Node `xml:"node,omitempty"`
// }

//     Library_Geometries    LibraryGeometries
//     Library_Visual_Scenes LibraryVisualScenes
// }
//
// type LibraryGeometries struct{
//     XMLName  xml.Name   `xml:"library_geometries"`
//     Geometry []Geometry
// }
//
// type Geometry struct{
//     XMLName xml.Name  `xml:"geometry"`
//     Id      string    `xml:"attr"`
//     Mesh    Mesh
// }
//
// type Mesh struct {
//     XMLName  xml.Name `xml:"mesh"`
//     Source   []Source
//     Polylist Polylist
// }
//
// type Source struct{
//     XMLName     xml.Name   `xml:"source"`
//     Id          string     `xml:"attr"`
//     Float_array FloatArray `xml:"float_array"`
// }
//
// type FloatArray struct{
//     XMLName xml.Name `xml:"float_array"`
//     Id      string   `xml:"attr"`
//     CDATA   string   `xml:"chardata"`
//     Count   string   `xml:"attr"`
// }
//
// type Polylist struct{
//     XMLName xml.Name  `xml:"polylist"`
//     Id      string    `xml:"attr"`
//     Count   string    `xml:"attr"`
//
//     // List of integers, each specifying the number of vertices for one polygon
//     VCount  string    `xml:"vcount"`
//
//     // list of integers that specify the vertex attributes
//     P       string    `xml:"p"`
// }
//
// type LibraryVisualScenes struct {
//     XMLName      xml.Name       `xml:"library_visual_scenes"`
//     VisualScene  VisualScene
// }
//
// type VisualScene struct{
//     XMLName      xml.Name       `xml:"visual_scene"`
// }

func LoadDocument(filename string) (*Collada, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	collada, err := LoadDocumentFromReader(file)
	return collada, err
}

func LoadDocumentFromReader(reader io.Reader) (*Collada, error) {
	decoder := xml.NewDecoder(reader)
	collada := &Collada{}
	err := decoder.Decode(collada)
	if err != nil {
		return nil, err
	}
	return collada, nil
}

func (collada *Collada) Export(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString(xml.Header)
	w.Flush()
	encoder := xml.NewEncoder(file)
	encoder.Indent("", " ")
	return encoder.Encode(collada)
}
