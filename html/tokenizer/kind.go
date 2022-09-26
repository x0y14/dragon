package tokenizer

type Kind int

const (
	_ Kind = iota
	Illegal
	Eof
	Whitespace

	TagBegin    // <
	TagEnd      // >
	Exclamation // !
	Assign      // =
	Hyphen      // -
	Slash       // /
	Amp         // &

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
	KWLabel
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
	Illegal:      "Illegal",
	Eof:          "Eof",
	Whitespace:   "Whitespace",
	TagBegin:     "TagBegin",
	TagEnd:       "TagEnd",
	Exclamation:  "Exclamation",
	Assign:       "Assign",
	Hyphen:       "Hyphen",
	Slash:        "Slash",
	Amp:          "Amp",
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
	KWLabel:      "KWLabel",
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

func symbolKind(symbol string) Kind {
	var kind Kind
	switch symbol {
	case "<":
		kind = TagBegin
	case ">":
		kind = TagEnd
	case "!":
		kind = Exclamation
	case "=":
		kind = Assign
	case "-":
		kind = Hyphen
	case "/":
		kind = Slash
	case "&":
		kind = Amp
	default:
		kind = Illegal
	}
	return kind
}

func identKind(ident string) Kind {
	switch ident {
	case "docType":
		return KWDocType
	case "html":
		return KWHtml
	case "base":
		return KWBase
	case "head":
		return KWHead
	case "link":
		return KWLink
	case "meta":
		return KWMeta
	case "style":
		return KWStyle
	case "title":
		return KWTitle
	case "body":
		return KWBody
	case "address":
		return KWAddress
	case "article":
		return KWArticle
	case "aside":
		return KWAside
	case "footer":
		return KWFooter
	case "header":
		return KWHeader
	case "h1":
		return KWH1
	case "h2":
		return KWH2
	case "h3":
		return KWH3
	case "h4":
		return KWH4
	case "h5":
		return KWH5
	case "h6":
		return KWH6
	case "main":
		return KWMain
	case "nav":
		return KWNav
	case "section":
		return KWSection
	case "blockquote":
		return KWBlockquote
	case "dd":
		return KWDd
	case "div":
		return KWDiv
	case "dl":
		return KWDl
	case "dt":
		return KWDt
	case "figcaption":
		return KWFigcaption
	case "figure":
		return KWFigure
	case "hr":
		return KWHr
	case "li":
		return KWLi
	case "menu":
		return KWMenu
	case "ol":
		return KWOl
	case "p":
		return KWP
	case "pre":
		return KWPre
	case "ul":
		return KWUl
	case "a":
		return KWA
	case "abbr":
		return KWAbbr
	case "b":
		return KWB
	case "bdi":
		return KWBdi
	case "bdo":
		return KWBdo
	case "br":
		return KWBr
	case "cite":
		return KWCite
	case "code":
		return KWCode
	case "data":
		return KWData
	case "dfn":
		return KWDfn
	case "em":
		return KWEm
	case "i":
		return KWI
	case "kbd":
		return KWKbd
	case "mark":
		return KWMark
	case "q":
		return KWQ
	case "rt":
		return KWRt
	case "ruby":
		return KWRuby
	case "s":
		return KWS
	case "samp":
		return KWSamp
	case "small":
		return KWSmall
	case "span":
		return KWSpan
	case "strong":
		return KWStrong
	case "sub":
		return KWSub
	case "sup":
		return KWSup
	case "time":
		return KWTime
	case "u":
		return KWU
	case "var":
		return KWVar
	case "wbr":
		return KWWbr
	case "area":
		return KWArea
	case "audio":
		return KWAudio
	case "img":
		return KWImg
	case "map":
		return KWMap
	case "track":
		return KWTrack
	case "video":
		return KWVideo
	case "embed":
		return KWEmbed
	case "iframe":
		return KWIframe
	case "object":
		return KWObject
	case "picture":
		return KWPicture
	case "portal":
		return KWPortal
	case "source":
		return KWSource
	case "svg":
		return KWSvg
	case "math":
		return KWMath
	case "canvas":
		return KWCanvas
	case "noscript":
		return KWNoscript
	case "script":
		return KWScript
	case "del":
		return KWDel
	case "ins":
		return KWIns
	case "caption":
		return KWCaption
	case "col":
		return KWCol
	case "colgroup":
		return KWColgroup
	case "table":
		return KWTable
	case "tbody":
		return KWTbody
	case "td":
		return KWTd
	case "tfoot":
		return KWTfoot
	case "th":
		return KWTh
	case "thread":
		return KWThread
	case "tr":
		return KWTr
	case "button":
		return KWButton
	case "datalist":
		return KWDatalist
	case "fieldset":
		return KWFieldset
	case "form":
		return KWForm
	case "input":
		return KWInput
	case "label":
		return KWLabel
	case "legend":
		return KWLegend
	case "meter":
		return KWMeter
	case "optgroup":
		return KWOptgroup
	case "option":
		return KWOption
	case "output":
		return KWOutput
	case "progress":
		return KWProgress
	case "select":
		return KWSelect
	case "textarea":
		return KWTextarea
	case "details":
		return KWDetails
	case "dialog":
		return KWDialog
	case "summary":
		return KWSummary
	case "slot":
		return KWSlot
	case "template":
		return KWTemplate
	}
	return Identifier
}
