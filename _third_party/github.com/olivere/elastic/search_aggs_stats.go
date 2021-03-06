// Copyright 2012-2014 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

// StatsAggregation is a multi-value metrics aggregation that computes stats
// over numeric values extracted from the aggregated documents.
// These values can be extracted either from specific numeric fields
// in the documents, or be generated by a provided script.
// See: http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-aggregations-metrics-stats-aggregation.html
type StatsAggregation struct {
	field           string
	script          string
	lang            string
	params          map[string]interface{}
	subAggregations map[string]Aggregation
}

func NewStatsAggregation() StatsAggregation {
	a := StatsAggregation{
		params:          make(map[string]interface{}),
		subAggregations: make(map[string]Aggregation),
	}
	return a
}

func (a StatsAggregation) Field(field string) StatsAggregation {
	a.field = field
	return a
}

func (a StatsAggregation) Script(script string) StatsAggregation {
	a.script = script
	return a
}

func (a StatsAggregation) Lang(lang string) StatsAggregation {
	a.lang = lang
	return a
}

func (a StatsAggregation) Param(name string, value interface{}) StatsAggregation {
	a.params[name] = value
	return a
}

func (a StatsAggregation) SubAggregation(name string, subAggregation Aggregation) StatsAggregation {
	a.subAggregations[name] = subAggregation
	return a
}

func (a StatsAggregation) Source() interface{} {
	// Example:
	//	{
	//    "aggs" : {
	//      "grades_stats" : { "stats" : { "field" : "grade" } }
	//    }
	//	}
	// This method returns only the { "stats" : { "field" : "grade" } } part.

	source := make(map[string]interface{})
	opts := make(map[string]interface{})
	source["stats"] = opts

	// ValuesSourceAggregationBuilder
	if a.field != "" {
		opts["field"] = a.field
	}
	if a.script != "" {
		opts["script"] = a.script
	}
	if a.lang != "" {
		opts["lang"] = a.lang
	}
	if len(a.params) > 0 {
		opts["params"] = a.params
	}

	// AggregationBuilder (SubAggregations)
	if len(a.subAggregations) > 0 {
		aggsMap := make(map[string]interface{})
		source["aggregations"] = aggsMap
		for name, aggregate := range a.subAggregations {
			aggsMap[name] = aggregate.Source()
		}
	}

	return source
}
