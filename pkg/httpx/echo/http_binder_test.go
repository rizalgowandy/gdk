package echo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	ID    int      `query:"id"`
	Name  string   `query:"name"`
	Tags  []string `query:"tags"`
	Score float64  `query:"score"`
	Flag  bool     `query:"flag"`
}

type embeddedStruct struct {
	Value int `query:"value"`
}

type testStructWithEmbed struct {
	embeddedStruct
	Name string `query:"name"`
}

func newTestContext(method, target string, query url.Values) echo.Context {
	e := echo.New()
	req := httptest.NewRequest(method, target, strings.NewReader(""))
	req.URL.RawQuery = query.Encode()
	return e.NewContext(req, nil)
}

func TestCustomBinder_Bind_PrimitiveFields(t *testing.T) {
	cb := &CustomBinder{}
	q := url.Values{}
	q.Set("id", "42")
	q.Set("name", "foo")
	q.Set("score", "3.14")
	q.Set("flag", "true")
	ctx := newTestContext(http.MethodGet, "/", q)
	var s testStruct
	err := cb.Bind(&s, ctx)
	require.NoError(t, err)
	assert.Equal(t, 42, s.ID)
	assert.Equal(t, "foo", s.Name)
	assert.InEpsilon(t, 3.14, s.Score, 0.0001)
	assert.True(t, s.Flag)
}

func TestCustomBinder_Bind_SliceField(t *testing.T) {
	cb := &CustomBinder{}
	q := url.Values{}
	q.Set("tags", "a,b,c")
	ctx := newTestContext(http.MethodGet, "/", q)
	var s testStruct
	err := cb.Bind(&s, ctx)
	require.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c"}, s.Tags)
}

func TestCustomBinder_Bind_EmbeddedStruct(t *testing.T) {
	cb := &CustomBinder{}
	q := url.Values{}
	q.Set("value", "99")
	q.Set("name", "bar")
	ctx := newTestContext(http.MethodGet, "/", q)
	var s testStructWithEmbed
	err := cb.Bind(&s, ctx)
	require.NoError(t, err)
	assert.Equal(t, 99, s.Value)
	assert.Equal(t, "bar", s.Name)
}

func TestCustomBinder_Bind_NonPointer(t *testing.T) {
	cb := &CustomBinder{}
	ctx := newTestContext(http.MethodGet, "/", url.Values{})
	var s testStruct
	err := cb.Bind(s, ctx)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "pointer to struct")
}

func TestCustomBinder_Bind_NonStructPointer(t *testing.T) {
	cb := &CustomBinder{}
	ctx := newTestContext(http.MethodGet, "/", url.Values{})
	var s int
	err := cb.Bind(&s, ctx)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "expects struct")
}

func TestCustomBinder_bindValue_Pointer(t *testing.T) {
	cb := &CustomBinder{}
	var v *int
	val := reflect.ValueOf(&v).Elem()
	err := cb.bindValue(val, val.Type(), "123", "id")
	require.NoError(t, err)
	assert.NotNil(t, v)
	assert.Equal(t, 123, *v)
}

func TestCustomBinder_bindSlice_Primitive(t *testing.T) {
	cb := &CustomBinder{}
	var s []int
	val := reflect.ValueOf(&s).Elem()
	err := cb.bindSlice(val, val.Type(), "1,2,3", "id")
	require.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, s)
}

func TestCustomBinder_bindSlice_PointerElem(t *testing.T) {
	cb := &CustomBinder{}
	var s []*int
	val := reflect.ValueOf(&s).Elem()
	err := cb.bindSlice(val, val.Type(), "1,2", "id")
	require.NoError(t, err)
	assert.Len(t, s, 2)
	assert.Equal(t, 1, *s[0])
	assert.Equal(t, 2, *s[1])
}

func TestCustomBinder_bindPrimitive_Errors(t *testing.T) {
	cb := &CustomBinder{}
	var i int
	val := reflect.ValueOf(&i).Elem()
	err := cb.bindPrimitive(val, "notanint", "id")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid int")
}
