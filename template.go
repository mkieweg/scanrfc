package scanrfc

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
)

var t *template.Template
var writer io.Writer

func init() {
	t = template.New("rfc")
	t := t.Delims("<<", ">>")
	_, err := t.Parse(entry)
	if err != nil {
		log.Fatal(err)
	}
}

const entry = `
@misc{rfc<< .Number >>,
	series =	{Request for Comments},
	number =	<< .Number >>,
	howpublished =	{RFC << .Number >>},
	publisher =	{RFC Editor},
	doi =		{10.17487/RFC<< .Number >>},
	url =		{https://rfc-editor.org/rfc/rfc<< .Number >>.txt},
        author =	{<<range $element := .Authors>><<$element.Name>> <<end>>},
	title =		{{<< .Title >>}},
	pagetotal =	<< .Pages >>,
	year =		<< .Year >>,
	month =		<< .Month >>,
	abstract =	{ << .Abstract >>},
}
`

type Author struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Affiliation string `json:"affiliation"`
}

type RfcEntry struct {
	Number   string
	Authors  []Author `json:"authors"`
	Title    string   `json:"title"`
	Pages    int      `json:"pages"`
	Year     int
	Month    string
	Abstract string `json:"abstract"`

	Date JsonDate `json:"time"`
}

type JsonDate time.Time

func (jd *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	s = strings.Split(s, " ")[0]
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*jd = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

func CreateEntry(n string, b []byte) (*RfcEntry, error) {
	entry := &RfcEntry{
		Date: JsonDate{},
	}
	if err := json.Unmarshal(b, entry); err != nil {
		return nil, err
	}
	date := time.Time(entry.Date)
	entry.Year = date.Year()
	entry.Month = date.Month().String()
	entry.Number = n
	return entry, nil
}

func ParseJsonResponse(rfc string, resp []byte) ([]byte, error) {
	n := strings.Trim(rfc, "rfc")
	out := bytes.NewBuffer(make([]byte, 0))
	writer = out
	e, err := CreateEntry(n, resp)
	if err != nil {
		return nil, err
	}
	if err := t.Execute(writer, e); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
