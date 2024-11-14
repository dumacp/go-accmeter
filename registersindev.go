package accmeter

var INT_REL = Register{Name: "INT_REL", Addr: 0x17, Len: 1}
var CNTL1 = Register{Name: "CNTL1", Addr: 0x18, Len: 1}
var CNTL2 = Register{Name: "CNTL2", Addr: 0x19, Len: 1}
var CNTL3 = Register{Name: "CNTL3", Addr: 0x1A, Len: 1}

var INC1 = Register{Name: "INC1", Addr: 0x1C, Len: 1}
var INC2 = Register{Name: "INC2", Addr: 0x1D, Len: 1}
var INC3 = Register{Name: "INC3", Addr: 0x1E, Len: 1}
var INC4 = Register{Name: "INC4", Addr: 0x1F, Len: 1}

var WUFC = Register{Name: "WUFC", Addr: 0x23, Len: 1}

var ATH = Register{Name: "ATH", Addr: 0x30, Len: 1}

var XHP_L = Register{Name: "XHP_L", Addr: 0x00, Len: 1}
var YHP_L = Register{Name: "YHP_L", Addr: 0x02, Len: 1}
var ZHP_L = Register{Name: "ZHP_L", Addr: 0x04, Len: 1}
var XHP_H = Register{Name: "XHP_H", Addr: 0x01, Len: 1}
var YHP_H = Register{Name: "YHP_H", Addr: 0x03, Len: 1}
var ZHP_H = Register{Name: "ZHP_H", Addr: 0x05, Len: 1}

var XHP = Register{Name: "XHP", Addr: 0x00, Len: 2}
var YHP = Register{Name: "YHP", Addr: 0x02, Len: 2}
var ZHP = Register{Name: "ZHP", Addr: 0x04, Len: 2}

var XOUT = Register{Name: "XOUT", Addr: 0x06, Len: 2}
var YOUT = Register{Name: "YOUT", Addr: 0x08, Len: 2}
var ZOUT = Register{Name: "ZOUT", Addr: 0x0A, Len: 2}

var ODCNTL = Register{Name: "ODCNTL", Addr: 0x1B, Len: 1}
