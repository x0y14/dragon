package tokenizer

type Kind int

const (
	_ Kind = iota
	Eof

	TagBegin    // <
	TagEnd      // >
	Exclamation // !
	Assign      // =
	Hyphen      // -
	Slash       // /

	Identifier
	String
	Integer
	Decimal

	KWDocType
	KWHtml

	KWBase
	KWHead
	KWLink
	KWMeta
	KWStyle
	KWTitle

	KWBody

	KWAddress
	KWArticle
	KWAside
	KWFooter
	KWHeader
	KWH1
	KWH2
	KWH3
	KWH4
	KWH5
	KWH6
	KWMain
	KWNav
	KWSection

	KWBlockquote
	KWDd
	KWDiv
	KWDl
	KWDt
	KWFigcaption
	KWFigure
	KWHr
	KWLi
	KWMenu
	KWOl
	KWP
	KWPre
	KWUl

	KWA
	KWAbbr
	KWB
	KWBdi
	KWBdo
	KWBr
	KWCite
	KWCode
	KWData
	KWDfn
	KWEm
	KWI
	KWKbd
	KWMark
	KWQ
	KWRt
	KWRuby
	KWS
	KWSamp
	KWSmall
	KWSpan
	KWStrong
	KWSub
	KWSup
	KWTime
	KWU
	KWVar
	KWWbr

	KWArea
	KWAudio
	KWImg
	KWMap
	KWTrack
	KWVideo

	KWEmbed
	KWIframe
	KWObject
	KWPicture
	KWPortal
	KWSource

	KWSvg
	KWMath

	KWCanvas
	KWNoscript
	KWScript

	KWDel
	KWIns

	KWCaption
	KWCol
	KWColgroup
	KWTable
	KWTbody
	KWTd
	KWTfoot
	KWTh
	KWThread
	KWTr

	KWButton
	KWDatalist
	KWFieldset
	KWForm
	KWInput
	LWLabel
	KWLegend
	KWMeter
	KWOptgroup
	KWOption
	KWOutput
	KWProgress
	KWSelect
	KWTextarea

	KWDetails
	KWDialog
	KWSummary

	KWSlot
	KWTemplate
)

var kinds = [...]string{
	Eof:          "Eof",
	TagBegin:     "TagBegin",
	TagEnd:       "TagEnd",
	Exclamation:  "Exclamation",
	Assign:       "Assign",
	Hyphen:       "Hyphen",
	Slash:        "Slash",
	Identifier:   "Identifier",
	String:       "String",
	Integer:      "Integer",
	Decimal:      "Decimal",
	KWDocType:    "KWDocType",
	KWHtml:       "KWHtml",
	KWBase:       "KWBase",
	KWHead:       "KWHead",
	KWLink:       "KWLink",
	KWMeta:       "KWMeta",
	KWStyle:      "KWStyle",
	KWTitle:      "KWTitle",
	KWBody:       "KWBody",
	KWAddress:    "KWAddress",
	KWArticle:    "KWArticle",
	KWAside:      "KWAside",
	KWFooter:     "KWFooter",
	KWHeader:     "KWHeader",
	KWH1:         "KWH1",
	KWH2:         "KWH2",
	KWH3:         "KWH3",
	KWH4:         "KWH4",
	KWH5:         "KWH5",
	KWH6:         "KWH6",
	KWMain:       "KWMain",
	KWNav:        "KWNav",
	KWSection:    "KWSection",
	KWBlockquote: "KWBlockquote",
	KWDd:         "KWDd",
	KWDiv:        "KWDiv",
	KWDl:         "KWDl",
	KWDt:         "KWDt",
	KWFigcaption: "KWFigcaption",
	KWFigure:     "KWFigure",
	KWHr:         "KWHr",
	KWLi:         "KWLi",
	KWMenu:       "KWMenu",
	KWOl:         "KWOl",
	KWP:          "KWP",
	KWPre:        "KWPre",
	KWUl:         "KWUl",
	KWA:          "KWA",
	KWAbbr:       "KWAbbr",
	KWB:          "KWB",
	KWBdi:        "KWBdi",
	KWBdo:        "KWBdo",
	KWBr:         "KWBr",
	KWCite:       "KWCite",
	KWCode:       "KWCode",
	KWData:       "KWData",
	KWDfn:        "KWDfn",
	KWEm:         "KWEm",
	KWI:          "KWI",
	KWKbd:        "KWKbd",
	KWMark:       "KWMark",
	KWQ:          "KWQ",
	KWRt:         "KWRt",
	KWRuby:       "KWRuby",
	KWS:          "KWS",
	KWSamp:       "KWSamp",
	KWSmall:      "KWSmall",
	KWSpan:       "KWSpan",
	KWStrong:     "KWStrong",
	KWSub:        "KWSub",
	KWSup:        "KWSup",
	KWTime:       "KWTime",
	KWU:          "KWU",
	KWVar:        "KWVar",
	KWWbr:        "KWWbr",
	KWArea:       "KWArea",
	KWAudio:      "KWAudio",
	KWImg:        "KWImg",
	KWMap:        "KWMap",
	KWTrack:      "KWTrack",
	KWVideo:      "KWVideo",
	KWEmbed:      "KWEmbed",
	KWIframe:     "KWIframe",
	KWObject:     "KWObject",
	KWPicture:    "KWPicture",
	KWPortal:     "KWPortal",
	KWSource:     "KWSource",
	KWSvg:        "KWSvg",
	KWMath:       "KWMath",
	KWCanvas:     "KWCanvas",
	KWNoscript:   "KWNoscript",
	KWScript:     "KWScript",
	KWDel:        "KWDel",
	KWIns:        "KWIns",
	KWCaption:    "KWCaption",
	KWCol:        "KWCol",
	KWColgroup:   "KWColgroup",
	KWTable:      "KWTable",
	KWTbody:      "KWTbody",
	KWTd:         "KWTd",
	KWTfoot:      "KWTfoot",
	KWTh:         "KWTh",
	KWThread:     "KWThread",
	KWTr:         "KWTr",
	KWButton:     "KWButton",
	KWDatalist:   "KWDatalist",
	KWFieldset:   "KWFieldset",
	KWForm:       "KWForm",
	KWInput:      "KWInput",
	LWLabel:      "LWLabel",
	KWLegend:     "KWLegend",
	KWMeter:      "KWMeter",
	KWOptgroup:   "KWOptgroup",
	KWOption:     "KWOption",
	KWOutput:     "KWOutput",
	KWProgress:   "KWProgress",
	KWSelect:     "KWSelect",
	KWTextarea:   "KWTextarea",
	KWDetails:    "KWDetails",
	KWDialog:     "KWDialog",
	KWSummary:    "KWSummary",
	KWSlot:       "KWSlot",
	KWTemplate:   "KWTemplate",
}

func (k Kind) String() string {
	return kinds[k]
}
