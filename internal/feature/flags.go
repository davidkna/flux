// Code generated by the feature package; DO NOT EDIT.

package feature

import (
	"context"

	"github.com/influxdata/flux/internal/pkg/feature"
)

type (
	Flag       = feature.Flag
	Flagger    = feature.Flagger
	StringFlag = feature.StringFlag
	FloatFlag  = feature.FloatFlag
	IntFlag    = feature.IntFlag
	BoolFlag   = feature.BoolFlag
)

var aggregateTransformationTransport = feature.MakeBoolFlag(
	"Aggregate Transformation Transport",
	"aggregateTransformationTransport",
	"Jonathan Sternberg",
	false,
)

// AggregateTransformationTransport - Enable Transport interface for AggregateTransformation
func AggregateTransformationTransport() BoolFlag {
	return aggregateTransformationTransport
}

var groupTransformationGroup = feature.MakeBoolFlag(
	"Group Transformation Group",
	"groupTransformationGroup",
	"Sean Brickley",
	true,
)

// GroupTransformationGroup - Enable GroupTransformation interface for the group function
func GroupTransformationGroup() BoolFlag {
	return groupTransformationGroup
}

var optimizeUnionTransformation = feature.MakeBoolFlag(
	"Optimize Union Transformation",
	"optimizeUnionTransformation",
	"Jonathan Sternberg",
	false,
)

// OptimizeUnionTransformation - Optimize the union transformation
func OptimizeUnionTransformation() BoolFlag {
	return optimizeUnionTransformation
}

var narrowTransformationDifference = feature.MakeBoolFlag(
	"Narrow Transformation Difference",
	"narrowTransformationDifference",
	"Markus Westerlind",
	false,
)

// NarrowTransformationDifference - Enable the NarrowTransformation implementation of difference
func NarrowTransformationDifference() BoolFlag {
	return narrowTransformationDifference
}

var narrowTransformationFill = feature.MakeBoolFlag(
	"Narrow Transformation Fill",
	"narrowTransformationFill",
	"Sunil Kartikey",
	false,
)

// NarrowTransformationFill - Enable the NarrowTransformation implementation of Fill
func NarrowTransformationFill() BoolFlag {
	return narrowTransformationFill
}

var optimizeAggregateWindow = feature.MakeBoolFlag(
	"Optimize Aggregate Window",
	"optimizeAggregateWindow",
	"Jonathan Sternberg",
	true,
)

// OptimizeAggregateWindow - Enables a version of aggregateWindow written in Go
func OptimizeAggregateWindow() BoolFlag {
	return optimizeAggregateWindow
}

var labelPolymorphism = feature.MakeBoolFlag(
	"Label polymorphism",
	"labelPolymorphism",
	"Markus Westerlind",
	false,
)

// LabelPolymorphism - Enables label polymorphism in the type system
func LabelPolymorphism() BoolFlag {
	return labelPolymorphism
}

var optimizeSetTransformation = feature.MakeBoolFlag(
	"Optimize Set Transformation",
	"optimizeSetTransformation",
	"Jonathan Sternberg",
	false,
)

// OptimizeSetTransformation - Enables a version of set that is optimized
func OptimizeSetTransformation() BoolFlag {
	return optimizeSetTransformation
}

var unusedSymbolWarnings = feature.MakeBoolFlag(
	"Unused Symbol Warnings",
	"unusedSymbolWarnings",
	"Markus Westerlind",
	false,
)

// UnusedSymbolWarnings - Enables warnings for unused symbols
func UnusedSymbolWarnings() BoolFlag {
	return unusedSymbolWarnings
}

var queryConcurrencyIncrease = feature.MakeIntFlag(
	"Query Concurrency Increase",
	"queryConcurrencyIncrease",
	"Jonathan Sternberg, Adrian Thurston",
	0,
)

// QueryConcurrencyIncrease - Additional dispatcher workers to allocate on top of the minimimum allowable computed by the engine
func QueryConcurrencyIncrease() IntFlag {
	return queryConcurrencyIncrease
}

var strictNullLogicalOps = feature.MakeBoolFlag(
	"StrictNullLogicalOps",
	"strictNullLogicalOps",
	"Owen Nelson",
	false,
)

// Strictnulllogicalops - When enabled, nulls in logical expressions should match the behavior language spec.
func Strictnulllogicalops() BoolFlag {
	return strictNullLogicalOps
}

var prettyError = feature.MakeBoolFlag(
	"Pretty Error",
	"prettyError",
	"Markus Westerlind",
	false,
)

// PrettyError - Enables formatting with codespan for errors
func PrettyError() BoolFlag {
	return prettyError
}

var salsaDatabase = feature.MakeBoolFlag(
	"Salsa Database",
	"salsaDatabase",
	"Markus Westerlind",
	false,
)

// SalsaDatabase - Enables the salsa database for semantic analysis
func SalsaDatabase() BoolFlag {
	return salsaDatabase
}

// Inject will inject the Flagger into the context.
func Inject(ctx context.Context, flagger Flagger) context.Context {
	return feature.Inject(ctx, flagger)
}

var all = []Flag{
	aggregateTransformationTransport,
	groupTransformationGroup,
	optimizeUnionTransformation,
	narrowTransformationDifference,
	narrowTransformationFill,
	optimizeAggregateWindow,
	labelPolymorphism,
	optimizeSetTransformation,
	unusedSymbolWarnings,
	queryConcurrencyIncrease,
	strictNullLogicalOps,
	prettyError,
	salsaDatabase,
}

var byKey = map[string]Flag{
	"aggregateTransformationTransport": aggregateTransformationTransport,
	"groupTransformationGroup":         groupTransformationGroup,
	"optimizeUnionTransformation":      optimizeUnionTransformation,
	"narrowTransformationDifference":   narrowTransformationDifference,
	"narrowTransformationFill":         narrowTransformationFill,
	"optimizeAggregateWindow":          optimizeAggregateWindow,
	"labelPolymorphism":                labelPolymorphism,
	"optimizeSetTransformation":        optimizeSetTransformation,
	"unusedSymbolWarnings":             unusedSymbolWarnings,
	"queryConcurrencyIncrease":         queryConcurrencyIncrease,
	"strictNullLogicalOps":             strictNullLogicalOps,
	"prettyError":                      prettyError,
	"salsaDatabase":                    salsaDatabase,
}

// Flags returns all feature flags.
func Flags() []Flag {
	return all
}

// ByKey returns the Flag corresponding to the given key.
func ByKey(k string) (Flag, bool) {
	v, found := byKey[k]
	return v, found
}

type Metrics = feature.Metrics

// SetMetrics sets the metric store for feature flags.
func SetMetrics(m Metrics) {
	feature.SetMetrics(m)
}
