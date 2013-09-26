/*
Package collada implements a schema for importing and exporting collada V1.5 (.dea) documents
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

type Id string

const (
	Xup UpAxis = "X_UP"
	Yup UpAxis = "Y_UP"
	Zup UpAxis = "Z_UP"
)

type Opaque string

const (
	OpaqueAlphaZero = "A_ZERO"
	OpaqueAlphaOne  = "A_ONE"
	OpaqueRgbZero   = "RGB_ZERO"
	OpaqueRgbOne    = "RGB_ONE"
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
	Camera []*Camera `xml:"camera"`
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
	Skeleton     []*Skeleton    `xml:"skeleton"`
	BindMaterial *BindMaterial `xml:"bind_material"`
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
	HasId
	HasCount
	HasName
	Digits    uint8  `xml:"digits,attr,omitempty"`
	Magnitude uint16 `xml:"magnitude,attr,omitempty"`
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
	Source   Uri `xml:"source,attr"`
	Set      uint   `xml:"set,attr,omitempty"`
}

// InputUnshared declares the input semantics of a data source.
type InputUnshared struct {
	Semantic string `xml:"semantic,attr"`
	Source   Uri `xml:"source,attr"`
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
	Xmlns   string `xml:"xmlns,attr,omitempty"`
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
	Geometry []*Geometry `xml:"geometry"`
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
	Source     []*Source     `xml:"source"`
	Vertices   Vertices     `xml:"vertices"`
	Lines      []*Lines      `xml:"lines"`
	Linestrips []*Linestrips `xml:"linestrips"`
	Polygons   []*Polygons   `xml:"polygons"`
	Polylist   []*Polylist   `xml:"polylist"`
	Triangles  []*Triangles  `xml:"triangles"`
	Trifans    []*Trifans    `xml:"trifans"`
	Tristrips  []*Tristrips  `xml:"tristrips"`
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
	Ph []*Ph `xml:"ph"`
}
type Ph struct {
	P P   `xml:"p"`
	H []*H `xml:"h"`
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
	Input []*InputUnshared `xml:"input"`
	HasExtra
}

//AmbientCore (core) Describes an ambient light source.
type AmbientCore struct {
	//TODO
}

//Color describes the color of its parent light element.
type Color struct {
	HasSid
	Float3
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
	Light []*Light `xml:"light"`
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
	GeographicLocation []*GeographicLocation `xml:"geographic_location"`
}

//Unit defines unit of distance for COLLADA elements and objects.
type Unit struct {
	HasName
	Meter float64 `xml:"meter,attr"`
}

//Asset defines asset-management information regarding its parent element.
type Asset struct {
	Contributor []*Contributor `xml:"contributor"`
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
	XMLName string  `xml:"COLLADA"`
	Xmlns   Uri     `xml:"xmlns,attr,omitempty"`
	Version Version `xml:"version,attr"`
	Base    Uri     `xml:"base,attr,omitempty"`
	HasAsset
	LibraryAnimationClips []*LibraryAnimationClips `xml:"library_animation_clips"`
	LibraryAnimations     []*LibraryAnimations     `xml:"library_animations"`
	// LibraryArticulatedSystems []*LibraryArticulatedSystems `xml:"library_animation_clips"`
	LibraryCameras     []*LibraryCameras     `xml:"library_cameras"`
	LibraryControllers []*LibraryControllers `xml:"library_controllers"`
	LibraryLights      []*LibraryLights      `xml:"library_lights"`
	LibraryImages      []*LibraryImages      `xml:"library_images"`
	LibraryEffects     []*LibraryEffects     `xml:"library_effects"`
	// LibraryForceFields []*LibraryForceFields `xml:"library_force_fields"`
	LibraryFormulas []*LibraryFormulas `xml:"library_formulas"`
	// LibraryJoints []*LibraryJoints `xml:"library_joints"`
	// LibraryKinematicModels []*LibraryKinematicModels `xml:"library_kinematic_models"`
	// LibraryKinematicScenes []*LibraryKinematicScenes `xml:"library_kinematic_scenes"`
	LibraryMaterials  []*LibraryMaterials  `xml:"library_materials"`
	LibraryGeometries []*LibraryGeometries `xml:"library_geometries"`
	// LibraryPhysicsNodes []*LibraryPhysicsNodes `xml:"library_physics_nodes"`
	// LibraryPhysicsMaterials []*LibraryPhysicsMaterials `xml:"library_physics_materials"`
	// LibraryPhysicsScenes []*LibraryPhysicsScenes `xml:"library_physics_scenes"`
	// LibraryPhysicsScenes []*LibraryPhysicsScenes `xml:"library_physics_scenes"`
	LibraryVisualScenes []*LibraryVisualScenes `xml:"library_visual_scenes"`
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
	VisualScene []*VisualScene `xml:"visual_scene"`
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
	Lookat             []*Lookat             `xml:"lookat"`
	Matrix             []*Matrix             `xml:"matrix"`
	Translate          []*Translate          `xml:"translate"`
	Rotate             []*Rotate             `xml:"rotate"`
	Scale              []*Scale              `xml:"scale"`
	Skew               []*Skew               `xml:"skew"`
	InstanceCamera     []*InstanceCamera     `xml:"instance_camera"`
	InstanceController []*InstanceController `xml:"instance_controller"`
	InstanceGeometry   []*InstanceGeometry   `xml:"instance_geometry"`
	InstanceLight      []*InstanceLight      `xml:"instance_light"`
	InstanceNode       []*InstanceNode       `xml:"instance_node"`
	HasNodes
	HasExtra
}

//Scene embodies the entire set of information that can be visualized from the contents of a COLLADA resource.
type Scene struct {
	InstancePhysicsScene    []*InstancePhysicsScene   `xml:"instance_physics_scene"`
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
	EvaluateScene []*EvaluateScene `xml:"evaluate_scene"`
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

//Annotate Adds a strongly typed annotation remark to the parent object.
type Annotate struct {
	//TODO
}

//BindVertexInput Binds geometry vertex inputs to effect vertex inputs upon instantiation.
type BindVertexInput struct {
	//TODO
}

//Effect Provides a self-contained description of a COLLADA effect.
type Effect struct {
	HasId
	HasName
	HasAsset
	HasAnnotate
	HasNewparam
	ProfileBridge *ProfileBridge `xml:"profile_BRIDGE"`
	ProfileCg     *ProfileCg     `xml:"profile_CG"`
	ProfileGles   *ProfileGles   `xml:"profile_GLES"`
	ProfileGles2  *ProfileGles2  `xml:"profile_GLES2"`
	ProfileGlsl   *ProfileGlsl   `xml:"profile_GLSL"`
	ProfileCommon *ProfileCommon `xml:"profile_COMMON"`
	HasExtra
}

//InstanceEffect Instantiates a COLLADA effect.
type InstanceEffect struct {
	HasId
	HasName
	HasUrl
	TechniqueHint []*TechniqueHint `xml:"technique_hint"`
	Setparam      []*Setparam      `xml:"setparam"`
}

//LibraryEffects Provides a library in which to place <effect> assets.
type LibraryEffects struct {
	HasId
	HasName
	HasAsset
	Effect []*Effect `xml:"effect"`
	HasExtra
}

//TechniqueFx Holds a description of the textures, samplers, shaders, parameters, and passes necessary for rendering this effect using one method.
type TechniqueFx struct {
	HasId
	HasSid
	HasAsset
	HasAnnotate
	Blinn      *Blinn      `xml:"blinn"`
	ConstantFx *ConstantFx `xml:"constant"`
	Lambert    *Lambert    `xml:"lambert"`
	Phone      *Phong      `xml:"phong"`
	Pass       *Pass       `xml:"pass"`
	HasExtra
}

//TechniqueHint Adds a hint for a platform of which technique to use in this effect
type TechniqueHint struct {
	Platform string `xml:"platform,attr,omitempty"`
	Ref      string `xml:"ref,attr"`
	Profile  string `xml:"profile,attr,omitempty"`
}

//BindMaterial Binds a specific material to a piece of geometry, binding varying and uniform parameters at the same time.
type BindMaterial struct {
	Param []*ParamCore `xml:"param"`
	HasTechniqueCommon
	HasTechnique
	HasExtra
}

//InstanceMaterialGeometry Instantiates a COLLADA material resource.
type InstanceMaterialGeometry struct {
	//TODO
}

//LibraryMaterials Provides a library in which to place <material> assets.
type LibraryMaterials struct {
	HasId
	HasName
	HasAsset
	Material []*Material `xml:"material"`
	HasExtra
}

//Material Defines the equations necessary for the visual appearance of geometry and screenspace image processing
type Material struct {
	HasId
	HasName
	HasAsset
	InstanceEffect InstanceEffect `xml:"instance_effect"`
	HasExtra
}

//Array Creates a parameter of a one-dimensional array type.
type Array struct {
	//TODO
}

//Modifier Provides additional information about the volatility or linkage of a <newparam>declaration.
type Modifier struct {
	//TODO
}

//Newparam Creates a new, named parameter object and assigns it a type and an initial value. See Chapter 5: Core Elements Reference.
type Newparam struct {
	//TODO
}

//ParamReference (reference) References a predefined parameter. See Chapter 5: Core Elements Reference.
type ParamReference struct {
	//TODO
}

//SamplerStates Allows users to modify an effect’s sampler state from a material.
type SamplerStates struct {
	//TODO
}

//Semantic Provides metadata that describes the purpose of a parameter declaration.
type Semantic struct {
	//TODO
}

//Setparam Assigns a new value to a previously defined parameter. See main entry in Chapter 5: Core Elements Reference.
type Setparam struct {
	Ref string `xml:"ref,attr"`
}

//Usertype Creates an instance of a structured class for a parameter.
type Usertype struct {
	//TODO
}

//ProfileBridge Provides support for referencing effect profiles written with external standards.
type ProfileBridge struct {
	HasId
	HasAsset
	HasNewparam
}

//ProfileCg Declares a platform-specific representation of an effect written in the NVIDIA®Cg language.
type ProfileCg struct {
	//TODO
}

//ProfileCommon Opens a block of platform-independent declarations for the common, fixed-function shader.
type ProfileCommon struct {
	HasId
	HasAsset
	HasNewparam
	HasTechniqueFx
	HasExtra
}
type ProfileGles struct {
	//TODO
}
type ProfileGles2 struct {
	//TODO
}
type ProfileGlsl struct {
	//TODO
}
type Blinn struct {
	//TODO
}
type ColorClear struct {
	//TODO
}
type ColorTarget struct {
	//TODO
}

// A type that describes color attributes of fixed-function shader elements
type FxCommonColorOrTextureType struct {
	Opaque  Opaque          `xml:"opaque,attr,omitempty"`
	Color   *Color          `xml:"color"`
	Param   *ParamReference `xml:"param"`
	Texture *Texture        `xml:"texture"`
}

type Texture struct {
	Texture  string `xml:"texture,attr"`
	TexCoord string `xml:"texcoord,attr"`
	HasExtra
}

// A type that describes the scalar attributes of fixed-function shader elements inside <profile_COMMON> effects. See main entry.
type FxCommonFloatOrParamType struct {
	Float *Float          `xml:"float"`
	Param *ParamReference `xml:"param"`
}

//Constant Produces a constantly shaded surface that is independent of lighting.
type ConstantFx struct {
	//TODO
}

//DepthClear Specifies whether a render target image is to be cleared, and which value to use.
type DepthClear struct {
	//TODO
}

//DepthTarget Specifies which <image> will receive the depth information from the output of this pass.
type DepthTarget struct {
	//TODO
}

//Draw Instructs the FX Runtime what kind of geometry to submit.
type Draw struct {
	//TODO
}

//Evaluate Contains evaluation elements for a rendering pass.
type Evaluate struct {
	//TODO
}

//InstanceMaterialRendering Instantiates a COLLADA material resource for a screen effect.
type InstanceMaterialRendering struct {
	//TODO
}

//Lambert Produces a diffuse shaded surface that is independent of lighting.
type Lambert struct {
	//TODO
}

//Pass Provides a static declaration of all the render states, shaders, and settings for one rendering pipeline.
type Pass struct {
	//TODO
}

//Phong Produces a shaded surface where the specular reflection is shaded
type Phong struct {
	Emission          *FxCommonColorOrTextureType `xml:"emission"`
	AmbientFx         *FxCommonColorOrTextureType `xml:"ambient"`
	Diffuse           *FxCommonColorOrTextureType `xml:"diffuse"`
	Specular          *FxCommonColorOrTextureType `xml:"specular"`
	Shininess         *FxCommonFloatOrParamType   `xml:"shininess"`
	Reflective        *FxCommonColorOrTextureType `xml:"reflective"`
	Reflectivity      *FxCommonFloatOrParamType   `xml:"reflectivity"`
	Transparent       *FxCommonColorOrTextureType `xml:"transparent"`
	Transparency      *FxCommonFloatOrParamType   `xml:"transparency"`
	IndexOfRefraction *FxCommonFloatOrParamType   `xml:"index_of_refraction"`
}

//According the Phong BRDF approximation.
type According struct {
	//TODO
}

//Render Describes one effect pass to evaluate a scene.
type Render struct {
	//TODO
}

//States Contains all rendering states to set up for the parent pass.
type States struct {
	//TODO
}

//StencilClear Specifies whether a render target image is to be cleared, and which value to use.
type StencilClear struct {
	//TODO
}

//StencilTarget Specifies which <image> will receive the stencil information from the output of this pass
type StencilTarget struct {
	//TODO
}

//Binary Identifies or provides a shader in binary form.
type Binary struct {
	//TODO
}

//BindAttribute Binds semantics to vertex attribute inputs of a shader.
type BindAttribute struct {
	//TODO
}

//BindUniform Binds values to uniform inputs of a shader or binds values to effect parameters upon instantiation.
type BindUniform struct {
	//TODO
}

//Code Provides an inline block of source code.
type Code struct {
	//TODO
}

//Compiler Contains command-line or runtime-invocation options for a shader compiler.
type Compiler struct {
	//TODO
}

//Include Imports source code or precompiled binary shaders into the FX Runtime by referencing an external resource.
type Include struct {
	//TODO
}

//Linker Contains command-line or runtime-invocation options for shader linkers to combine shaders into programs.
type Linker struct {
	//TODO
}

//Program Links multiple shaders together to produce a pipeline for geometry processing.
type Program struct {
	//TODO
}
type Shader struct {
	//TODO
}
type Sources struct {
	//TODO
}
type Alpha struct {
	//TODO
}
type Argument struct {
	//TODO
}
type Create2d struct {
	//TODO
}
type Create3d struct {
	//TODO
}
type CreateCube struct {
	//TODO
}
type Format struct {
	//TODO
}
type Image struct {
	//TODO
}
type InitFrom struct {
	//TODO
}
type InstanceImage struct {
	//TODO
}
type LibraryImages struct {
	//TODO
}
type Rgb struct {
	//TODO
}
type FxSamplerCommon struct {
	//TODO
}
type Sampler1D struct {
	//TODO
}
type Sampler2D struct {
	//TODO
}
type Sampler3D struct {
	//TODO
}
type SamplerCube struct {
	//TODO
}
type SamplerDepth struct {
	//TODO
}
type SamplerRect struct {
	//TODO
}
type Texcombiner struct {
	//TODO
}
type Texenv struct {
	//TODO
}

//TexturePipeline Defines a set of texturing commands that will be converted into multitexturing operations using glTexEnv in regular and combiner mode.
type TexturePipeline struct {
	//TODO
}

type P struct {
    Ints
}

type Floats struct {
    Values
}
type Bools struct {
    Values
}
type IdRefs struct {
    Values
}
type Ints struct {
    Values
}
type Names struct {
    Values
}
type SidRefs struct {
    Values
}
type Tokens struct {
    Values
}

type Values struct {
	V string `xml:",chardata"`
}

type Float3x3 struct {
    Floats
}
type Float4x4 struct {
    Floats
}
type Float4 struct {
    Floats
}
type Float3 struct {
    Floats
}

type Float struct {
	HasSid
	Value float64 `xml:",chardata"`
}

type HasNewparam struct {
	Newparam []*Newparam `xml:"newparam"`
}
type HasAnnotate struct {
	Annotate []*Annotate `xml:"annotate"`
}
type HasSharedInput struct {
	Input []*InputShared `xml:"input"`
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
	Id Id `xml:"id,attr,omitempty"`
}
type HasUrl struct {
	Url Uri `xml:"url,attr,omitempty"`
}
type HasSid struct {
	Sid string `xml:"sid,attr,omitempty"`
}
type HasAsset struct {
	Asset *Asset `xml:"asset,omitempty"`
}
type HasNodes struct {
	Node []*Node `xml:"node"`
}
type HasExtra struct {
	Extra []*Extra `xml:"extra"`
}
type HasTechniqueCommon struct {
	TechniqueCommon TechniqueCommon `xml:"technique_common"`
}
type HasTechnique struct {
	TechniqueCore []*TechniqueCore `xml:"technique,omitempty"`
}
type HasTechniqueFx struct {
	TechniqueFx []*TechniqueFx `xml:"technique,omitempty"`
}
type HasP struct {
	P *P `xml:"p"`
}
type HasPs struct {
	P []*P `xml:"p"`
}

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
	return collada.ExportToWriter(file)
}

func (collada *Collada) ExportToWriter(writer io.Writer) error {
	w := bufio.NewWriter(writer)
	w.WriteString(xml.Header)
	w.Flush()
	encoder := xml.NewEncoder(writer)
	encoder.Indent("", " ")
	return encoder.Encode(collada)
}
