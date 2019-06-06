package inkrecognizer

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v1.0/inkrecognizer"

// Application enumerates the values for application.
type Application string

const (
	// Drawing ...
	Drawing Application = "drawing"
	// Mixed ...
	Mixed Application = "mixed"
	// Writing ...
	Writing Application = "writing"
)

// PossibleApplicationValues returns an array of possible values for the Application const type.
func PossibleApplicationValues() []Application {
	return []Application{Drawing, Mixed, Writing}
}

// Category enumerates the values for category.
type Category string

const (
	// InkBullet ...
	InkBullet Category = "inkBullet"
	// InkDrawing ...
	InkDrawing Category = "inkDrawing"
	// InkWord ...
	InkWord Category = "inkWord"
	// Line ...
	Line Category = "line"
	// ListItem ...
	ListItem Category = "listItem"
	// Paragraph ...
	Paragraph Category = "paragraph"
	// Unknown ...
	Unknown Category = "unknown"
	// WritingRegion ...
	WritingRegion Category = "writingRegion"
)

// PossibleCategoryValues returns an array of possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{InkBullet, InkDrawing, InkWord, Line, ListItem, Paragraph, Unknown, WritingRegion}
}

// Class enumerates the values for class.
type Class string

const (
	// ClassContainer ...
	ClassContainer Class = "container"
	// ClassLeaf ...
	ClassLeaf Class = "leaf"
)

// PossibleClassValues returns an array of possible values for the Class const type.
func PossibleClassValues() []Class {
	return []Class{ClassContainer, ClassLeaf}
}

// Container enumerates the values for container.
type Container string

const (
	// ContainerLine ...
	ContainerLine Container = "line"
	// ContainerListItem ...
	ContainerListItem Container = "listItem"
	// ContainerParagraph ...
	ContainerParagraph Container = "paragraph"
	// ContainerWritingRegion ...
	ContainerWritingRegion Container = "writingRegion"
)

// PossibleContainerValues returns an array of possible values for the Container const type.
func PossibleContainerValues() []Container {
	return []Container{ContainerLine, ContainerListItem, ContainerParagraph, ContainerWritingRegion}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindInkDrawing ...
	KindInkDrawing Kind = "inkDrawing"
	// KindInkWriting ...
	KindInkWriting Kind = "inkWriting"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindInkDrawing, KindInkWriting}
}

// Leaf enumerates the values for leaf.
type Leaf string

const (
	// LeafInkBullet ...
	LeafInkBullet Leaf = "inkBullet"
	// LeafInkDrawing ...
	LeafInkDrawing Leaf = "inkDrawing"
	// LeafInkWord ...
	LeafInkWord Leaf = "inkWord"
	// LeafUnknown ...
	LeafUnknown Leaf = "unknown"
)

// PossibleLeafValues returns an array of possible values for the Leaf const type.
func PossibleLeafValues() []Leaf {
	return []Leaf{LeafInkBullet, LeafInkDrawing, LeafInkWord, LeafUnknown}
}

// RasterOp enumerates the values for raster op.
type RasterOp string

const (
	// CopyPen ...
	CopyPen RasterOp = "copyPen"
	// MaskPen ...
	MaskPen RasterOp = "maskPen"
	// NoOperation ...
	NoOperation RasterOp = "noOperation"
)

// PossibleRasterOpValues returns an array of possible values for the RasterOp const type.
func PossibleRasterOpValues() []RasterOp {
	return []RasterOp{CopyPen, MaskPen, NoOperation}
}

// Shape enumerates the values for shape.
type Shape string

const (
	// ShapeBlockArrow ...
	ShapeBlockArrow Shape = "blockArrow"
	// ShapeCircle ...
	ShapeCircle Shape = "circle"
	// ShapeCloud ...
	ShapeCloud Shape = "cloud"
	// ShapeCurve ...
	ShapeCurve Shape = "curve"
	// ShapeDiamond ...
	ShapeDiamond Shape = "diamond"
	// ShapeDrawing ...
	ShapeDrawing Shape = "drawing"
	// ShapeEllipse ...
	ShapeEllipse Shape = "ellipse"
	// ShapeEquilateralTriangle ...
	ShapeEquilateralTriangle Shape = "equilateralTriangle"
	// ShapeHeart ...
	ShapeHeart Shape = "heart"
	// ShapeHexagon ...
	ShapeHexagon Shape = "hexagon"
	// ShapeIsoscelesTriangle ...
	ShapeIsoscelesTriangle Shape = "isoscelesTriangle"
	// ShapeLine ...
	ShapeLine Shape = "line"
	// ShapeParallelogram ...
	ShapeParallelogram Shape = "parallelogram"
	// ShapePentagon ...
	ShapePentagon Shape = "pentagon"
	// ShapePolyLine ...
	ShapePolyLine Shape = "polyLine"
	// ShapeQuadrilateral ...
	ShapeQuadrilateral Shape = "quadrilateral"
	// ShapeRectangle ...
	ShapeRectangle Shape = "rectangle"
	// ShapeRightTriangle ...
	ShapeRightTriangle Shape = "rightTriangle"
	// ShapeSquare ...
	ShapeSquare Shape = "square"
	// ShapeStarCrossed ...
	ShapeStarCrossed Shape = "starCrossed"
	// ShapeStarSimple ...
	ShapeStarSimple Shape = "starSimple"
	// ShapeTrapezoid ...
	ShapeTrapezoid Shape = "trapezoid"
	// ShapeTriangle ...
	ShapeTriangle Shape = "triangle"
)

// PossibleShapeValues returns an array of possible values for the Shape const type.
func PossibleShapeValues() []Shape {
	return []Shape{ShapeBlockArrow, ShapeCircle, ShapeCloud, ShapeCurve, ShapeDiamond, ShapeDrawing, ShapeEllipse, ShapeEquilateralTriangle, ShapeHeart, ShapeHexagon, ShapeIsoscelesTriangle, ShapeLine, ShapeParallelogram, ShapePentagon, ShapePolyLine, ShapeQuadrilateral, ShapeRectangle, ShapeRightTriangle, ShapeSquare, ShapeStarCrossed, ShapeStarSimple, ShapeTrapezoid, ShapeTriangle}
}

// Tip enumerates the values for tip.
type Tip string

const (
	// Ellipse ...
	Ellipse Tip = "ellipse"
	// Rectangle ...
	Rectangle Tip = "rectangle"
)

// PossibleTipValues returns an array of possible values for the Tip const type.
func PossibleTipValues() []Tip {
	return []Tip{Ellipse, Rectangle}
}

// Unit enumerates the values for unit.
type Unit string

const (
	// Cm ...
	Cm Unit = "cm"
	// In ...
	In Unit = "in"
	// Mm ...
	Mm Unit = "mm"
)

// PossibleUnitValues returns an array of possible values for the Unit const type.
func PossibleUnitValues() []Unit {
	return []Unit{Cm, In, Mm}
}

// Unit1 enumerates the values for unit 1.
type Unit1 string

const (
	// Unit1Cm ...
	Unit1Cm Unit1 = "cm"
	// Unit1In ...
	Unit1In Unit1 = "in"
	// Unit1Mm ...
	Unit1Mm Unit1 = "mm"
)

// PossibleUnit1Values returns an array of possible values for the Unit1 const type.
func PossibleUnit1Values() []Unit1 {
	return []Unit1{Unit1Cm, Unit1In, Unit1Mm}
}

// AlternatePatternItem ...
type AlternatePatternItem struct {
	// Category - Possible values include: 'LeafInkDrawing', 'LeafInkBullet', 'LeafInkWord', 'LeafUnknown'
	Category Leaf `json:"category,omitempty"`
	// Points - Array of point objects that represent points that are relevant to the type of recognition unit. For example, for leaf node of inkDrawing category that represents a triangle, points would include the x,y coordinates of the vertices of the recognized triangle. The points represent the coordinates of points used to create the perfectly drawn shape that is closest to the original input. They may not exactly match.
	Points *[]PointDetailsPattern `json:"points,omitempty"`
	// RotationAngle - The angular orientation of an object relative to the horizontal axis
	RotationAngle *float64 `json:"rotationAngle,omitempty"`
	// Confidence - A number between 0 and 1 which indicates the confidence level in the result
	Confidence *float64 `json:"confidence,omitempty"`
	// RecognizedString - The recognized string from an inkWord or the name of a recognized shape in an inkDrawing object
	RecognizedString *string `json:"recognizedString,omitempty"`
}

// AnalysisRequest this shows the expected contents of a request
type AnalysisRequest struct {
	// ApplicationType - This describes the domain of the client application. Possible values include: 'Drawing', 'Writing', 'Mixed'
	ApplicationType Application `json:"applicationType,omitempty"`
	// Unit - This represents the physical units of the ink strokes. It is up to the application developer to decide how to convert the device specific units to physical units before calling the service. The conversion factor can be different based on the type of the device used. Possible values include: 'Unit1Mm', 'Unit1Cm', 'Unit1In'
	Unit Unit1 `json:"unit,omitempty"`
	// UnitMultiple -  This is a scaling factor to be applied to the point coordinates when interpreting them in the physical units specified.
	UnitMultiple *float64 `json:"unitMultiple,omitempty"`
	// Language - The IETF BCP 47 language code (for ex. en-US, en-GB, hi-IN etc.) of the expected language for the handwritten content in the ink strokes. The response will include results from this language.
	Language *string `json:"language,omitempty"`
	// Strokes - This is the array of strokes sent for recognition. Best results are produced when the order of strokes added in the array matches the order in which the user created them. Changing the stroke order may produce unexpected results.
	Strokes *[]Stroke `json:"strokes,omitempty"`
}

// AnalysisResponse this shows the expected contents of a response from the service
type AnalysisResponse struct {
	autorest.Response `json:"-"`
	// Unit - This represents the physical units of the ink strokes. It is up to the application developer to decide how to convert the device specific units to physical units before calling the service. The conversion factor can be different based on the type of the device used. Possible values include: 'Mm', 'Cm', 'In'
	Unit Unit `json:"unit,omitempty"`
	// UnitMultiple -  This is a scaling factor to be applied to the point coordinates when interpreting them in the physical units specified.
	UnitMultiple *float64 `json:"unitMultiple,omitempty"`
	// Language - This is the language used for recognizing handwriting from the ink strokes in the request. Set this to the user’s preferred language.
	Language         *string                `json:"language,omitempty"`
	RecognitionUnits *[]RecognitionUnitItem `json:"recognitionUnits,omitempty"`
}

// DrawingAttributesPattern the properties to use when rendering ink
type DrawingAttributesPattern struct {
	// Width - The width of the stylus used to draw the stroke
	Width *float64 `json:"width,omitempty"`
	// Color - This shows the components of the color in rgba format
	Color *DrawingAttributesPatternColor `json:"color,omitempty"`
	// Height - The height of the stylus used to draw the stroke
	Height *float64 `json:"height,omitempty"`
	// FitToCurve -  This indicates whether Bezier smoothing is used to render the stroke
	FitToCurve *bool `json:"fitToCurve,omitempty"`
	// RasterOp - Possible values include: 'NoOperation', 'CopyPen', 'MaskPen'
	RasterOp RasterOp `json:"rasterOp,omitempty"`
	// IgnorePressure -  This indicates whether the thickness of a rendered Stroke changes according the amount of pressure applied.
	IgnorePressure *bool `json:"ignorePressure,omitempty"`
	// Tip - This specifies the tip to be used to draw a stroke. Possible values include: 'Ellipse', 'Rectangle'
	Tip Tip `json:"tip,omitempty"`
}

// DrawingAttributesPatternColor this shows the components of the color in rgba format
type DrawingAttributesPatternColor struct {
	// R - The red component of the color
	R *float64 `json:"r,omitempty"`
	// G - The green component of the color
	G *float64 `json:"g,omitempty"`
	// B - The blue component of the color
	B *float64 `json:"b,omitempty"`
	// A - The alpha component of the color
	A *float64 `json:"a,omitempty"`
}

// ErrorModel ...
type ErrorModel struct {
	// Code - This represents the error code
	Code *string `json:"code,omitempty"`
	// Message - This represents the error message
	Message *string `json:"message,omitempty"`
	// Target - This represents the target of the error message
	Target *string `json:"target,omitempty"`
	// Details - This gives details of the reason(s) for the error
	Details *[]ErrorModelDetailsItem `json:"details,omitempty"`
}

// ErrorModelDetailsItem ...
type ErrorModelDetailsItem struct {
	// Code - This represents the error code
	Code *string `json:"code,omitempty"`
	// Message - This represents the error message
	Message *string `json:"message,omitempty"`
	// Target - This represents the target of the error message
	Target *string `json:"target,omitempty"`
}

// PointDetailsPattern this holds all the properties of one point
type PointDetailsPattern struct {
	// X - This represents the x coordinate of the point
	X *float64 `json:"x,omitempty"`
	// Y - This represents the y coordinate of the point
	Y *float64 `json:"y,omitempty"`
}

// RecognitionUnitItem this represents the recognized entity
type RecognitionUnitItem struct {
	// ID - The identifier of the recognition unit. This id is used to indicate parent/child relationship between different recognition units.
	ID *int32 `json:"id,omitempty"`
	// Category - Possible values include: 'WritingRegion', 'Paragraph', 'ListItem', 'Line', 'InkBullet', 'InkDrawing', 'InkWord', 'Unknown'
	Category   Category                `json:"category,omitempty"`
	Alternates *[]AlternatePatternItem `json:"alternates,omitempty"`
	// Center - The coordinates (x,y) of the center of the recognition unit.
	Center *PointDetailsPattern `json:"center,omitempty"`
	// Points - Array of point objects that represent points that are relevant to the type of recognition unit. For example, for a leaf node of inkDrawing category that represents a triangle, points would include the x, y coordinates of the vertices of the recognized triangle. The points represent the coordinates used to create the perfectly drawn shape that is closest to the original input. They may not exactly match.
	Points *[]PointDetailsPattern `json:"points,omitempty"`
	// ChildIds - An array of integers representing the identifier of each child of the current recognition unit.
	ChildIds *[]int32 `json:"childIds,omitempty"`
	// Class - Possible values include: 'ClassContainer', 'ClassLeaf'
	Class Class `json:"class,omitempty"`
	// ParentID - The id of the parent node in the tree structure of the recognition results. parent = 0 indicates that there is no dedicated parent node for this unit.
	ParentID *int32 `json:"parentId,omitempty"`
	// BoundingRectangle - The bounding rectangle of the recognized unit represented by the coordinates of the top left corner (x,y) along with width (w) and height (h) of the rectangle. Note that this rectangle is not rotated. So for  rotated objects such as slanted handwriting, it will cover the entire object. The unit will be matched to the one specified in the original request (mm by default.)
	BoundingRectangle *RecognitionUnitItemBoundingRectangle `json:"boundingRectangle,omitempty"`
	// RotatedBoundingRectangle - This property provides the rotated bounding rectangle that covers the entire recognized object along the angle of rotation of the object. Note that this is NOT the same as rotating the boundingRectangle by the rotation angle.
	RotatedBoundingRectangle *[]PointDetailsPattern `json:"rotatedBoundingRectangle,omitempty"`
	// StrokeIds - This is an array of integers representing the list of stroke Identifier integers from the input request body that belong to this recognition unit.
	StrokeIds *[]int32 `json:"strokeIds,omitempty"`
	// RecognizedText - The string represents the text that was recognized. It can be an empty string if the recognizer cannot determine the text.
	RecognizedText *string `json:"recognizedText,omitempty"`
	// Confidence - The class represents the type of the recognition unit. A recognition unit can be a leaf node or a container node. Container nodes typically have leaf nodes as children.
	Confidence *float64 `json:"confidence,omitempty"`
	// RotationAngle - This represents the angle at which the unit is rotated in degrees with respect to the positive X axis.
	RotationAngle *float64 `json:"rotationAngle,omitempty"`
	// RecognizedObject - Possible values include: 'ShapeDrawing', 'ShapeSquare', 'ShapeRectangle', 'ShapeCircle', 'ShapeEllipse', 'ShapeTriangle', 'ShapeIsoscelesTriangle', 'ShapeEquilateralTriangle', 'ShapeRightTriangle', 'ShapeQuadrilateral', 'ShapeDiamond', 'ShapeTrapezoid', 'ShapeParallelogram', 'ShapePentagon', 'ShapeHexagon', 'ShapeBlockArrow', 'ShapeHeart', 'ShapeStarSimple', 'ShapeStarCrossed', 'ShapeCloud', 'ShapeLine', 'ShapeCurve', 'ShapePolyLine'
	RecognizedObject Shape `json:"recognizedObject,omitempty"`
}

// RecognitionUnitItemBoundingRectangle the bounding rectangle of the recognized unit represented by the
// coordinates of the top left corner (x,y) along with width (w) and height (h) of the rectangle. Note that
// this rectangle is not rotated. So for  rotated objects such as slanted handwriting, it will cover the
// entire object. The unit will be matched to the one specified in the original request (mm by default.)
type RecognitionUnitItemBoundingRectangle struct {
	// TopX - This represents the top left x coordinate
	TopX *float64 `json:"topX,omitempty"`
	// TopY - This represents the top left y coordinate
	TopY *float64 `json:"topY,omitempty"`
	// Width - This represents width of the bounding rectangle
	Width *float64 `json:"width,omitempty"`
	// Height - The represents the height of the bounding rectangle
	Height *float64 `json:"height,omitempty"`
}

// Stroke ...
type Stroke struct {
	// ID - This is treated as a unique identifier for each stroke within a request. If the id is repeated within the same request, the service will return an error.
	ID *int32 `json:"id,omitempty"`
	// Language - The IETF BCP 47 language code (for ex. en-US, en-GB, hi-IN etc.) of the expected language for the handwritten content in this stroke. The response will include results from this language.
	Language *string `json:"language,omitempty"`
	// Points - A string of comma separated floating point values that represent the x and y coordinates of points that are part of the stroke. (X1,Y1, X2,Y2…).  It is recommended to have a precision of 8 digits after the decimal to obtain most accurate recognition results. The origin (0,0) of the canvas is assumed to be at the top left corner of the canvas
	Points            *string                   `json:"points,omitempty"`
	DrawingAttributes *DrawingAttributesPattern `json:"drawingAttributes,omitempty"`
	// Kind - This is an optional property which influences the decision about what the stroke kind is between inkWriting and inkDrawing. This property should be set ONLY if the type of user content is known ahead of time. Not setting this value implies the kind is not known ahead of time. Kind represents the type of content the stroke is a part of. Possible values include: 'KindInkDrawing', 'KindInkWriting'
	Kind Kind `json:"kind,omitempty"`
}
