package objects

import (
	"fmt"
	"github.com/mmsuo/vcalender/objects/property"
	"github.com/mmsuo/vcalender/objects/property/components"
	"github.com/mmsuo/vcalender/objects/property/components/properties/changemanage"
	"github.com/mmsuo/vcalender/objects/property/components/properties/datetime"
	"github.com/mmsuo/vcalender/objects/property/components/properties/descriptive"
	"github.com/mmsuo/vcalender/objects/property/components/properties/relationship"
	"testing"
)

func TestCalendar_Calendar(t *testing.T) {
	//       BEGIN:VCALENDAR
	//       VERSION:2.0
	//       PRODID:-//hacksw/handcal//NONSGML v1.0//EN
	//       BEGIN:VEVENT
	//       UID:19970610T172345Z-AF23B2@example.com
	//       DTSTAMP:19970610T172345Z
	//       DTSTART:19970714T170000Z
	//       DTEND:19970715T040000Z
	//       SUMMARY:Bastille Day Party
	//       END:VEVENT
	//       END:VCALENDAR
	c := Calendar{
		ProdId:  property.NewProductIdentifier("-//hacksw/handcal//NONSGML v1.0//EN"),
		Version: &property.Version2,
		Components: []components.Component{
			&components.Event{
				DtStamp: changemanage.NewDtStamp(1997, 6, 10, 17, 23, 45),
				Uid:     relationship.NewUid("19970610T172345Z-AF23B2@example.com"),
				DtStart: datetime.NewDateStartWithDatetime(1997, 7, 14, 17, 0, 0),
				Summary: descriptive.NewSummary("Bastille Day Party"),
				DtEnd:   datetime.NewDateTimeDateEnd(1997, 7, 15, 4, 0, 0),
			},
		},
	}
	calendar, _ := c.Calendar()
	fmt.Println(calendar)
}

func BenchmarkCalendar_Calendar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := Calendar{
			ProdId:  property.NewProductIdentifier("-//hacksw/handcal//NONSGML v1.0//EN"),
			Version: &property.Version2,
			Components: []components.Component{
				&components.Event{
					DtStamp: changemanage.NewDtStamp(1997, 6, 10, 17, 23, 45),
					Uid:     relationship.NewUid("19970610T172345Z-AF23B2@example.com"),
					DtStart: datetime.NewDateStartWithDatetime(1997, 7, 14, 17, 0, 0),
					Summary: descriptive.NewSummary("Bastille Day Party"),
					DtEnd:   datetime.NewDateTimeDateEnd(1997, 7, 15, 4, 0, 0),
				},
			},
		}
		_, _ = c.Calendar()
	}
}
