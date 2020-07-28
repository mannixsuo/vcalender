RFC 5445 simple client implication in go

see https://tools.ietf.org/html/rfc5545

todo list:

- [x] Value Types
  - [x] Binary
  - [x] Boolean
  - [x] Calendar User Address
  - [x] Date
  - [x] Date-Time
  - [x] Duration
  - [x] Float
  - [x] Integer
  - [x] Period of Time
  - [ ] Recurrence Rule
  - [x] Text
  - [x] Time
  - [x] URI
  - [x] UTC Offset
- [x] Property Parameters
     - [x] Alternate Text Representation
     - [x] Common Name
     - [x] Calendar User Type
     - [x] Delegators  
     - [x] Delegatees  
     - [x] Directory Entry Reference
     - [x] Inline Encoding
     - [x] Format Type
     - [x] Free/Busy Time Type
     - [x] Language  
     - [x] Group or List Membership  
     - [x] Participation Status  
     - [x] Recurrence Identifier Range  
     - [x] Alarm Trigger Relationship  
     - [x] Relationship Type  
     - [x] Participation Role  
     - [x] RSVP Expectation  
     - [x] Sent By  
     - [x] Time Zone Identifier  
     - [x] Value Data Types
- [x] Component Properties
    - [x] Descriptive Component Properties
        - [x] Attachment
        - [x] Categories
        - [x] Classification
        - [x] Comment 
        - [x] Description 
        - [x] Geographic Position 
        - [x] Location  
        - [x] Percent Complete 
        - [x] Priority  
        - [x] Resources 
        - [x] Status  
    - [x] Date and Time Component Properties
        - [x] Date-Time Completed  
        - [x] Date-Time End  
        - [x] Date-Time Due  
        - [x] Date-Time Start  
        - [x] Duration  
        - [x] Free/Busy Time   
        - [x] Time Transparency
    - [x] Time Zone Component Properties
        - [x]  Time Zone Identifier
        - [x]  Time Zone Name
        - [x]  Time Zone Offset From
        - [x]  Time Zone Offset To
        - [x]  Time Zone URL
    - [x]  Relationship Component Properties
        - [x] Attendee  
        - [x]  Contact   
        - [x]  Organizer    
        - [x]  Recurrence ID   
        - [x]  Related To 
        - [x]  Uniform Resource Locator
        - [x]  Unique Identifier
    - [x]  Recurrence Component Properties
        - [x]  Exception Date-Times
        - [x]  Recurrence Date-Times
        - [x]  Recurrence Rule
    - [x]  Alarm Component Properties
        - [x]  Action  
        - [x]  Repeat Count
        - [x]  Trigger 
    - [x]   Change Management Component Properties
        - [x]  Date-Time Created
        - [x]  Date-Time Stamp
        - [x]  Last Modified
        - [x]  Sequence Number
    - [x]  Miscellaneous Component Properties
        - [x]  IANA Properties
        - [x]  Non-Standard Properties
        - [x]  Request Status
- [x] EVENT
- [x] TO-DO
- [x] JOURNAL
- [x] FREE/BUSY
- [x] TIME ZONE
- [x] Calender
- [ ] Error-check

## Usage:

```go
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
```