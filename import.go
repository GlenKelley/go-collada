package collada

import (
    "os"
    "bufio"
    "encoding/xml"
)

type Collada struct {
    XMLName string      `xml:"COLLADA"`
    Namespace string    `xml:"xmlns,attr,omitempty"`
    Version string      `xml:"version,attr"`
    HasAsset
    LibraryCameras LibraryCameras `xml:"library_cameras,omitempty"`
    LibraryLights LibraryLights `xml:"library_lights,omitempty"`
    LibraryImages LibraryImages `xml:"library_images,omitempty"`
    LibraryEffects LibraryEffects `xml:"library_effects,omitempty"`
    LibraryMaterials LibraryMaterials `xml:"library_materials,omitempty"`
    LibraryGeometries LibraryGeometries `xml:"library_geometries,omitempty"`
    LibraryVisualScenes LibraryVisualScenes `xml:"library_visual_scenes,omitempty"`
    Scene *Scene       `xml:"scene,omitempty"`
    HasExtras
}

type LibraryCameras struct {
    HasId
    HasName
    HasAsset
    Cameras []Camera `xml:"camera"`
    HasExtras
}

type Camera struct {
    HasId
    HasName
    HasAsset
    Optics Optics `xml:"optics"`
    Imager *Imager `xml:"imager,omitempty"`
    HasExtras
}

type Optics struct {
    TechniqueCommon TechniqueCommon `xml:"technique_common"`
    Technique []Technique `xml:"technique,omitempty"`
    HasExtras
}

type Imager struct {
    //TODO
}

type LibraryLights struct {
    HasId
    HasName
    HasAsset
    Lights []Light `xml:"light"`
    HasExtras
}

type Light struct {
    HasId
    HasName
    HasAsset
    TechniqueCommon TechniqueCommon `xml:"technique_common"`
    Techniques []Technique `xml:"technique"`
    HasExtras
}

type LibraryEffects struct {
    HasId
    HasName
    HasAsset
    Effects []Effect `xml:"effect"`
    HasExtras
}

type Effect struct {
    HasId
    HasName
    HasAsset
    // Annotations []Annotate `xml:"annotate"`
    Images []Image `xml:"image,omitempty"`
    // NewParams []NewParam `xml:"newparam"`
    ProfileCommon *ProfileCommon `xml:"profile_COMMON,omitempty"`
    ProfileGLES *ProfileGLES `xml:"profile_GLES,omitempty"`
    ProfileCG *ProfileCG `xml:"profile_CG,omitempty"`
    ProfileGLSL *ProfileGLSL `xml:"profile_GLSL,omitempty"`
    HasExtras
}

type ProfileCommon struct {
    XML string `xml:",innerxml"`
}
type ProfileCG struct {
    XML string `xml:",innerxml"`
}
type ProfileGLES struct {
    XML string `xml:",innerxml"`
}
type ProfileGLSL struct {
    XML string `xml:",innerxml"`
}

type Annotate struct {
    
}

type LibraryImages struct {
    HasId
    HasName
    HasAsset
    Images []Image `xml:"image"`
    HasExtras
}

type Image struct {
    HasId
    HasName
    Format string `xml:"token,attr"`
    Height uint `xlm:"height,attr"`
    Width uint `xml:"width,attr"`
    Depth uint `xml:"depth,attr"`
    HasAsset
    Data string `xml:"data,chardata"`
    InitFrom string `xml:"init_from"`
    HasExtras
}

type LibraryMaterials struct {
    HasId
    HasName
    HasAsset
    Materials []Material `xml:"material"`
    HasExtras
}

type Material struct {
    HasId
    HasName
    HasAsset
    InstanceEffect InstanceEffect `xml:"instance_effect"`
    HasExtras
}

type InstanceEffect struct {
    HasUrl
    HasSid
    HasName
    //TODO
    HasExtras
}

type LibraryVisualScenes struct {
    Scenes []VisualScene `xml:"visual_scene"`
}

type LibraryGeometries struct {
    HasId
    HasName
    HasAsset
    HasExtras
    Geometries []Geometry `xml:"geometry"`
}

type Geometry struct {
    HasId
    HasName
    HasAsset
    Meshes []Mesh `xml:"mesh"`
    ConvexMeshes []ConvexMesh `xml:"convex_mesh"`
    Splies []Spline `xml:"spline"`
    HasExtras
}

type Mesh struct {
    Sources []Source `xml:"source"`
    Vertices Vertices `xml:"vertices"`
    Lines []Lines `xml:"lines"`
    LineStrips []LineStrips `xml:"linestrips"`
    Polygons []Polygons `xml:"polygons"`
    Polylists []Polylist `xml:"polylist"`
    Triangles []Triangles `xml:"triangles"`
    Trifans []Trifans `xml:"trifans"`
    Tristrips []Tristrips `xml:"tristrips"`
    HasExtras
}

type Vertices struct {
    HasId
    HasName
    HasExtras
    Inputs []Input `xml:"input"`
}

type Input struct {
    Offset uint `xml:"offset,attr,omitempty"`
    Semantic string `xml:"semantic,attr,omitempty"`
    Source string `xml:"source,attr,omitempty"`
    Set uint `xml:"set,attr,omitempty"`
}

type BasicGeometry struct {
    HasName
    Count uint `xml:"count,attr"`
    Material string `xml:"material,attr"`
    HasExtras
    Inputs []Input `xml:"input"`
    P []string `xml:"p"`
}

type Lines BasicGeometry
type LineStrips BasicGeometry
type Triangles BasicGeometry
type Trifans BasicGeometry
type Tristrips BasicGeometry

type Polylist struct {
    BasicGeometry
    VCount string `xml:"vcount"`
}

type Polygons struct {
    BasicGeometry
    Ph []Ph `xml:"ph"`
}

type Ph struct {
    P string `xml:"p"`
    H []string `xml:"h"`
}

type Source struct {
    HasId
    HasName
    HasAsset
    IdRef *IdRefArray `xml:"IDREF_array"`
    Name *NameArray `xml:"Name_array"`
    Bool *BoolArray `xml:"bool_array"`
    Float *FloatArray `xml:"float_array"`
    Int *IntArray `xml:"int_array"`
    TechniqueCommon TechniqueCommon `xml:"technique_common"`
    Techniques []Technique `xml:"technique"`
}

type ValueArray struct {
    HasId
    HasName
    Count uint `xml:"count,attr"`
    Values string `xml:",chardata"`
}

type IdRefArray ValueArray

type IntArray struct {
    ValueArray
    MinInclusive int64 `xml:"minInclusive,attr,omitempty"`
    MaxInclusive int64 `xml:"maxInclusive,attr,omitempty"`
}

type FloatArray struct {
    ValueArray
    Digits int `xml:"digits,attr,omitempty"`
    Magnitude int `xml:"magnitude,attr,omitempty"`
}

type BoolArray ValueArray
type NameArray ValueArray

type ConvexMesh struct {
    Mesh
    ConvexHullOf string `xml:"convex_hull_of"`
}

type Spline struct {
    //TODO
}

type Asset struct {
    Contributor     []Contributor `xml:"contributor,omitempty"`
    Created         string `xml:"created,omitempty"`
    Keywords        string `xml:"keywords,omitempty"`
    Modified        string `xml:"modified,omitempty"`
    Revision        string `xml:"revision,omitempty"`
    Subject         string `xml:"subject,omitempty"`
    Title           string `xml:"title,omitempty"`
    Unit            string `xml:"unit,omitempty"`
    UpAxis          string `xml:"up_axis,omitempty"`
}

type Contributor struct {
    Author          string `xml:"author,omitempty"`
    AuthoringTool   string `xml:"authoring_tool,omitempty"`
    Comments        string `xml:"comments,omitempty"`
    Copyright       string `xml:"copyright,omitempty"`
    SourceData      string `xml:"sourceData,omitempty"`
}

type Unit struct {
    HasName
    Meter float64   `xml:"meter,attr"`
}

type Extra struct {
    HasId
    HasName
    HasType
    HasAsset
    Techniques []Technique `xml:"technique"`
}

type Scene struct {
    HasExtras
    // Physics InstanceWithExtra    `xml:"instance_physics_scene"`
    Visual InstanceWithExtra      `xml:"instance_visual_scene"`
}

type VisualScene struct {
    HasId
    HasName
    HasAsset
    HasExtras
    HasNodes
    Evaluate []EvaluateScene    `xml:"evaluate_scene"`
}

type Node struct {
    HasId
    HasName
    HasSid
    HasType
    Layer   string    `xml:"layer,attr"`
    HasAsset
    HasNodes
    HasExtras
    InstanceCamera  []InstanceWithExtra   `xml:"instance_camera"`
    InstanceController []InstanceController   `xml:"instance_controller"`
    InstanceGeometry []InstanceGeometry   `xml:"instance_geometry"`
    InstanceLight   []InstanceWithExtra   `xml:"instance_light"`
    InstanceNode    []InstanceWithExtra   `xml:"instance_node"`
    
    LookAt  []LookAt      `xml:"lookat"`
    Matrix  []Matrix      `xml:"matrix"`
    Rotate  []Rotate      `xml:"rotate"`
    Scale   []Scale       `xml:"scale"`
    Skew    []Skew        `xml:"skew"`
    Translate []Translate `xml:"translate"`
}

type Translate V3
type Rotate V3
type Scale V3
type Skew V3
type LookAt M3
type Matrix M4
type Uri string

type InstanceWithExtra struct {
    HasSid
    HasName
    HasUrl
    HasExtras
}

type InstanceController struct {
    //Skeleton
    //BindMaterial
    HasExtras
}

type InstanceGeometry struct {
    HasUrl
    HasExtras
    BindMaterial *BindMaterial `xml:"bind_material"`
}

type BindMaterial struct {
    HasExtras
    TechniqueCommon []TechniqueCommon `xml:"technique_common"`
    Technique []Technique `xml:"technique"`
}

type Technique struct {
    Profile string `xml:"profile,attr"`
    Content string `xml:",innerxml"`
}

type TechniqueCommon struct {
    Content string `xml:",innerxml"`
}

type InstanceMaterial struct {
    HasSid
    HasName
    HasExtras
    Symbol string `xml:"symbol"`
    Target string `xml:"target"`
    // Bind []Bind `xml:"bind"`
    // BindVertexInput `xml:"bind_vertex_input"`
}

type EvaluateScene struct {
    //TODO
}

type HasType struct {
    Type string `xml:"type,attr,omitempty"`
}

type HasSid struct {
    Sid string `xml:"sid,attr,omitempty"`
}

type HasName struct {
    Name string `xml:"name,attr,omitempty"`
}

type HasUrl struct {
    Url string `xml:"url,attr,omitempty"`
}

type V3 struct {
    HasSid
    V string `xml:",chardata"`
}

type V4 struct {
    HasSid
    V string `xml:",chardata"`
}

type M3 struct {
    HasSid
    V string `xml:",chardata"`
}

type M4 struct {
    HasSid
    V string `xml:",chardata"`
}

type HasId struct {
    Id string `xml:"id,attr,omitempty"`
}

type HasAsset struct {
    Asset *Asset `xml:"asset,omitempty"`
}

type HasExtras struct {
    Extras []Extra `xml:"extra,omitempty"`
}

type HasNodes struct {
    Nodes []Node `xml:"node,omitempty"`
}

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
    decoder := xml.NewDecoder(file)
    collada := &Collada{}
    err = decoder.Decode(collada)
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