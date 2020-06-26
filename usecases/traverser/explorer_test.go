//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package traverser

import (
	"context"
	"testing"

	"github.com/semi-technologies/weaviate/entities/filters"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema/kind"
	"github.com/semi-technologies/weaviate/entities/search"
	libprojector "github.com/semi-technologies/weaviate/usecases/projector"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Explorer_GetClass(t *testing.T) {
	t.Run("when an explore param is set", func(t *testing.T) {
		params := GetParams{
			Kind:      kind.Thing,
			ClassName: "BestClass",
			Explore: &ExploreParams{
				Values: []string{"foo"},
			},
			Pagination: &filters.Pagination{Limit: 100},
			Filters:    nil,
		}

		searchResults := []search.Result{
			{
				Kind: kind.Thing,
				ID:   "id1",
				Schema: map[string]interface{}{
					"name": "Foo",
				},
			},
			{
				Kind: kind.Action,
				ID:   "id2",
				Schema: map[string]interface{}{
					"age": 200,
				},
			},
		}

		search := &fakeVectorSearcher{}
		vectorizer := &fakeVectorizer{}
		extender := &fakeExtender{}
		log, _ := test.NewNullLogger()
		projector := &fakeProjector{}
		explorer := NewExplorer(search, vectorizer, newFakeDistancer(), log, extender, projector)
		expectedParamsToSearch := params
		expectedParamsToSearch.SearchVector = []float32{1, 2, 3}
		search.
			On("VectorClassSearch", expectedParamsToSearch).
			Return(searchResults, nil)

		res, err := explorer.GetClass(context.Background(), params)

		t.Run("vector search must be called with right params", func(t *testing.T) {
			assert.Nil(t, err)
			search.AssertExpectations(t)
		})

		t.Run("response must contain concepts", func(t *testing.T) {
			require.Len(t, res, 2)
			assert.Equal(t,
				map[string]interface{}{
					"name": "Foo",
				}, res[0])
			assert.Equal(t,
				map[string]interface{}{
					"age": 200,
				}, res[1])
		})
	})

	t.Run("when an explore param is set and the required certainty not met", func(t *testing.T) {
		params := GetParams{
			Kind:      kind.Thing,
			ClassName: "BestClass",
			Explore: &ExploreParams{
				Values:    []string{"foo"},
				Certainty: 0.8,
			},
			Pagination: &filters.Pagination{Limit: 100},
			Filters:    nil,
		}

		searchResults := []search.Result{
			{
				Kind: kind.Thing,
				ID:   "id1",
			},
			{
				Kind: kind.Action,
				ID:   "id2",
			},
		}

		search := &fakeVectorSearcher{}
		vectorizer := &fakeVectorizer{}
		extender := &fakeExtender{}
		log, _ := test.NewNullLogger()

		projector := &fakeProjector{}
		explorer := NewExplorer(search, vectorizer, newFakeDistancer(), log, extender, projector)
		expectedParamsToSearch := params
		expectedParamsToSearch.SearchVector = []float32{1, 2, 3}
		search.
			On("VectorClassSearch", expectedParamsToSearch).
			Return(searchResults, nil)

		res, err := explorer.GetClass(context.Background(), params)

		t.Run("vector search must be called with right params", func(t *testing.T) {
			assert.Nil(t, err)
			search.AssertExpectations(t)
		})

		t.Run("no concept met the required certainty", func(t *testing.T) {
			assert.Len(t, res, 0)
		})
	})

	t.Run("when no explore param is set", func(t *testing.T) {
		params := GetParams{
			Kind:       kind.Thing,
			ClassName:  "BestClass",
			Pagination: &filters.Pagination{Limit: 100},
			Filters:    nil,
		}

		searchResults := []search.Result{
			{
				Kind: kind.Thing,
				ID:   "id1",
				Schema: map[string]interface{}{
					"name": "Foo",
				},
			},
			{
				Kind: kind.Action,
				ID:   "id2",
				Schema: map[string]interface{}{
					"age": 200,
				},
			},
		}

		search := &fakeVectorSearcher{}
		vectorizer := &fakeVectorizer{}
		extender := &fakeExtender{}
		log, _ := test.NewNullLogger()
		projector := &fakeProjector{}
		explorer := NewExplorer(search, vectorizer, newFakeDistancer(), log, extender, projector)
		expectedParamsToSearch := params
		expectedParamsToSearch.SearchVector = nil
		search.
			On("ClassSearch", expectedParamsToSearch).
			Return(searchResults, nil)

		res, err := explorer.GetClass(context.Background(), params)

		t.Run("class search must be called with right params", func(t *testing.T) {
			assert.Nil(t, err)
			search.AssertExpectations(t)
		})

		t.Run("response must contain concepts", func(t *testing.T) {
			require.Len(t, res, 2)
			assert.Equal(t,
				map[string]interface{}{
					"name": "Foo",
				}, res[0])
			assert.Equal(t,
				map[string]interface{}{
					"age": 200,
				}, res[1])
		})
	})

	t.Run("when the _classification prop is set", func(t *testing.T) {
		params := GetParams{
			Kind:       kind.Thing,
			ClassName:  "BestClass",
			Pagination: &filters.Pagination{Limit: 100},
			Filters:    nil,
			UnderscoreProperties: UnderscoreProperties{
				Classification: true,
			},
		}

		searchResults := []search.Result{
			{
				Kind: kind.Thing,
				ID:   "id1",
				Schema: map[string]interface{}{
					"name": "Foo",
				},
				UnderscoreProperties: &models.UnderscoreProperties{
					Classification: nil,
				},
			},
			{
				Kind: kind.Action,
				ID:   "id2",
				Schema: map[string]interface{}{
					"age": 200,
				},
				UnderscoreProperties: &models.UnderscoreProperties{
					Classification: &models.UnderscorePropertiesClassification{
						ID: "1234",
					},
				},
			},
		}

		search := &fakeVectorSearcher{}
		vectorizer := &fakeVectorizer{}
		extender := &fakeExtender{}
		log, _ := test.NewNullLogger()
		projector := &fakeProjector{}
		explorer := NewExplorer(search, vectorizer, newFakeDistancer(), log, extender, projector)
		expectedParamsToSearch := params
		expectedParamsToSearch.SearchVector = nil
		search.
			On("ClassSearch", expectedParamsToSearch).
			Return(searchResults, nil)

		res, err := explorer.GetClass(context.Background(), params)

		t.Run("class search must be called with right params", func(t *testing.T) {
			assert.Nil(t, err)
			search.AssertExpectations(t)
		})

		t.Run("response must contain concepts", func(t *testing.T) {
			require.Len(t, res, 2)
			assert.Equal(t,
				map[string]interface{}{
					"name": "Foo",
				}, res[0])
			assert.Equal(t,
				map[string]interface{}{
					"age": 200,
					"_classification": &models.UnderscorePropertiesClassification{
						ID: "1234",
					},
				}, res[1])
		})
	})

	t.Run("when the _interpretation prop is set", func(t *testing.T) {
		params := GetParams{
			Kind:       kind.Thing,
			ClassName:  "BestClass",
			Pagination: &filters.Pagination{Limit: 100},
			Filters:    nil,
			UnderscoreProperties: UnderscoreProperties{
				Interpretation: true,
			},
		}

		searchResults := []search.Result{
			{
				Kind: kind.Thing,
				ID:   "id1",
				Schema: map[string]interface{}{
					"name": "Foo",
				},
				UnderscoreProperties: &models.UnderscoreProperties{
					Interpretation: nil,
				},
			},
			{
				Kind: kind.Action,
				ID:   "id2",
				Schema: map[string]interface{}{
					"age": 200,
				},
				UnderscoreProperties: &models.UnderscoreProperties{
					Interpretation: &models.Interpretation{
						Source: []*models.InterpretationSource{
							&models.InterpretationSource{
								Concept:    "foo",
								Weight:     0.123,
								Occurrence: 123,
							},
						},
					},
				},
			},
		}

		search := &fakeVectorSearcher{}
		vectorizer := &fakeVectorizer{}
		extender := &fakeExtender{}
		log, _ := test.NewNullLogger()
		projector := &fakeProjector{}
		explorer := NewExplorer(search, vectorizer, newFakeDistancer(), log, extender, projector)
		expectedParamsToSearch := params
		expectedParamsToSearch.SearchVector = nil
		search.
			On("ClassSearch", expectedParamsToSearch).
			Return(searchResults, nil)

		res, err := explorer.GetClass(context.Background(), params)

		t.Run("class search must be called with right params", func(t *testing.T) {
			assert.Nil(t, err)
			search.AssertExpectations(t)
		})

		t.Run("response must contain concepts", func(t *testing.T) {
			require.Len(t, res, 2)
			assert.Equal(t,
				map[string]interface{}{
					"name": "Foo",
				}, res[0])
			assert.Equal(t,
				map[string]interface{}{
					"age": 200,
					"_interpretation": &models.Interpretation{
						Source: []*models.InterpretationSource{
							&models.InterpretationSource{
								Concept:    "foo",
								Weight:     0.123,
								Occurrence: 123,
							},
						},
					},
				}, res[1])
		})
	})

	t.Run("when the _nearestNeighbors prop is set", func(t *testing.T) {
		params := GetParams{
			Kind:       kind.Thing,
			ClassName:  "BestClass",
			Pagination: &filters.Pagination{Limit: 100},
			Filters:    nil,
			UnderscoreProperties: UnderscoreProperties{
				NearestNeighbors: true,
			},
		}

		searchResults := []search.Result{
			{
				Kind: kind.Thing,
				ID:   "id1",
				Schema: map[string]interface{}{
					"name": "Foo",
				},
			},
			{
				Kind: kind.Action,
				ID:   "id2",
				Schema: map[string]interface{}{
					"name": "Bar",
				},
			},
		}

		searcher := &fakeVectorSearcher{}
		vectorizer := &fakeVectorizer{}
		log, _ := test.NewNullLogger()
		extender := &fakeExtender{
			returnArgs: []search.Result{
				{
					Kind: kind.Thing,
					ID:   "id1",
					Schema: map[string]interface{}{
						"name": "Foo",
					},
					UnderscoreProperties: &models.UnderscoreProperties{
						NearestNeighbors: &models.NearestNeighbors{
							Neighbors: []*models.NearestNeighbor{
								&models.NearestNeighbor{
									Concept:  "foo",
									Distance: 0.1,
								},
							},
						},
					},
				},
				{
					Kind: kind.Action,
					ID:   "id2",
					Schema: map[string]interface{}{
						"name": "Bar",
					},
					UnderscoreProperties: &models.UnderscoreProperties{
						NearestNeighbors: &models.NearestNeighbors{
							Neighbors: []*models.NearestNeighbor{
								&models.NearestNeighbor{
									Concept:  "bar",
									Distance: 0.1,
								},
							},
						},
					},
				},
			},
		}
		projector := &fakeProjector{}
		explorer := NewExplorer(searcher, vectorizer, newFakeDistancer(), log, extender, projector)
		expectedParamsToSearch := params
		expectedParamsToSearch.SearchVector = nil
		searcher.
			On("ClassSearch", expectedParamsToSearch).
			Return(searchResults, nil)

		res, err := explorer.GetClass(context.Background(), params)

		t.Run("class search must be called with right params", func(t *testing.T) {
			assert.Nil(t, err)
			searcher.AssertExpectations(t)
		})

		t.Run("response must contain concepts", func(t *testing.T) {
			require.Len(t, res, 2)
			assert.Equal(t,
				map[string]interface{}{
					"name": "Foo",
					"_nearestNeighbors": &models.NearestNeighbors{
						Neighbors: []*models.NearestNeighbor{
							&models.NearestNeighbor{
								Concept:  "foo",
								Distance: 0.1,
							},
						},
					},
				}, res[0])
			assert.Equal(t,
				map[string]interface{}{
					"name": "Bar",
					"_nearestNeighbors": &models.NearestNeighbors{
						Neighbors: []*models.NearestNeighbor{
							&models.NearestNeighbor{
								Concept:  "bar",
								Distance: 0.1,
							},
						},
					},
				}, res[1])
		})
	})

	t.Run("when the _featureProjection prop is set", func(t *testing.T) {
		params := GetParams{
			Kind:       kind.Thing,
			ClassName:  "BestClass",
			Pagination: &filters.Pagination{Limit: 100},
			Filters:    nil,
			UnderscoreProperties: UnderscoreProperties{
				FeatureProjection: &libprojector.Params{},
			},
		}

		searchResults := []search.Result{
			{
				Kind: kind.Thing,
				ID:   "id1",
				Schema: map[string]interface{}{
					"name": "Foo",
				},
			},
			{
				Kind: kind.Action,
				ID:   "id2",
				Schema: map[string]interface{}{
					"name": "Bar",
				},
			},
		}

		searcher := &fakeVectorSearcher{}
		vectorizer := &fakeVectorizer{}
		log, _ := test.NewNullLogger()
		extender := &fakeExtender{}
		projector := &fakeProjector{
			returnArgs: []search.Result{
				{
					Kind: kind.Thing,
					ID:   "id1",
					Schema: map[string]interface{}{
						"name": "Foo",
					},
					UnderscoreProperties: &models.UnderscoreProperties{
						FeatureProjection: &models.FeatureProjection{
							Vector: []float32{0, 1},
						},
					},
				},
				{
					Kind: kind.Action,
					ID:   "id2",
					Schema: map[string]interface{}{
						"name": "Bar",
					},
					UnderscoreProperties: &models.UnderscoreProperties{
						FeatureProjection: &models.FeatureProjection{
							Vector: []float32{1, 0},
						},
					},
				},
			},
		}
		explorer := NewExplorer(searcher, vectorizer, newFakeDistancer(), log, extender, projector)
		expectedParamsToSearch := params
		expectedParamsToSearch.SearchVector = nil
		searcher.
			On("ClassSearch", expectedParamsToSearch).
			Return(searchResults, nil)

		res, err := explorer.GetClass(context.Background(), params)

		t.Run("class search must be called with right params", func(t *testing.T) {
			assert.Nil(t, err)
			searcher.AssertExpectations(t)
		})

		t.Run("response must contain concepts", func(t *testing.T) {
			require.Len(t, res, 2)
			assert.Equal(t,
				map[string]interface{}{
					"name": "Foo",
					"_featureProjection": &models.FeatureProjection{
						Vector: []float32{0, 1},
					},
				}, res[0])
			assert.Equal(t,
				map[string]interface{}{
					"name": "Bar",
					"_featureProjection": &models.FeatureProjection{
						Vector: []float32{1, 0},
					},
				}, res[1])
		})
	})
}

func newFakeDistancer() func(a, b []float32) (float32, error) {
	return func(source, target []float32) (float32, error) {
		return 0.5, nil
	}
}
