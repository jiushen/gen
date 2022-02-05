package dbmeta

import (
	"gorm.io/gorm"
	"strings"
)

func LoadGormobjMeta(stmt gorm.Statement, tableName string) (DbTableMeta, error) {
	m := &dbTableMeta{
		sqlType:     "gormobj",
		sqlDatabase: "dummydb",
		tableName:   tableName,
	}

	//cols, err := schema.ColumnTypes(db, sqlDatabase, tableName)
	//if err != nil {
	//	return nil, err
	//}

	//ddl, err := mysqlLoadDDL(db, tableName)
	//if err != nil {
	//	return nil, fmt.Errorf("mysqlLoadDDL - unable to load ddl from mysql: %v", err)
	//}
	ddl := "dummy ddl"

	m.ddl = ddl
	//colsDDL, primaryKeys := mysqlParseDDL(ddl)
	//
	//infoSchema, err := LoadTableInfoFromMSSqlInformationSchema(db, tableName)
	//if err != nil {
	//	fmt.Printf("error calling LoadTableInfoFromMSSqlInformationSchema table: %s error: %v\n", tableName, err)
	//}
	m.columns = make([]*columnMeta, len(stmt.Schema.Fields))

	for i, v := range stmt.Schema.Fields {
		notes := ""
		nullable := !v.NotNull
		//nullable, ok := v.Nullable()
		//if !ok {
		//	nullable = false
		//}

		//colDDL := colsDDL[v.Name()]
		//isAutoIncrement := strings.Index(colDDL, "AUTO_INCREMENT") > -1
		//isUnsigned := strings.Index(colDDL, " unsigned ") > -1 || strings.Index(colDDL, " UNSIGNED ") > -1
		isAutoIncrement := v.AutoIncrement
		tagStr := string(v.Tag)
		isUnsigned := strings.Index(tagStr, " unsigned ") > -1 || strings.Index(tagStr, " UNSIGNED ") > -1

		//_, isPrimaryKey := find(primaryKeys, v.Name())
		isPrimaryKey := v.PrimaryKey
		//defaultVal := ""
		defaultVal := v.DefaultValue
		//columnType, columnLen := ParseSQLType(v.DatabaseTypeName())
		columnType := string(v.DataType)
		columnLen := int64(v.Size)

		if isUnsigned {
			notes = notes + " column is set for unsigned"
			columnType = "u" + columnType
		}

		//comment := ""
		//commentIdx := strings.Index(colDDL, "COMMENT '")
		//if commentIdx > -1 {
		//	re := regexp.MustCompile("COMMENT '(.*?)'")
		//	match := re.FindStringSubmatch(colDDL)
		//	if len(match) > 0 {
		//		comment = match[1]
		//	}
		//}
		comment := v.Comment

		//if infoSchema != nil {
		//	infoSchemaColInfo, ok := infoSchema[v.Name()]
		//	if ok {
		//		if infoSchemaColInfo.ColumnDefault != nil {
		//			defaultVal = BytesToString(infoSchemaColInfo.ColumnDefault.([]uint8))
		//			defaultVal = cleanupDefault(defaultVal)
		//		}
		//	}
		//}

		colMeta := &columnMeta{
			index:            i,
			name:             v.DBName,
			databaseTypeName: columnType,
			nullable:         nullable,
			isPrimaryKey:     isPrimaryKey,
			isAutoIncrement:  isAutoIncrement,
			colDDL:           tagStr,
			defaultVal:       defaultVal,
			columnType:       columnType,
			columnLen:        columnLen,
			notes:            strings.Trim(notes, " "),
			comment:          comment,
		}

		//dbType := strings.ToLower(colMeta.DatabaseTypeName())
		// fmt.Printf("dbType: %s\n", dbType)

		//if strings.Contains(dbType, "char") || strings.Contains(dbType, "text") {
		//	columnLen, err := GetFieldLenFromInformationSchema(db, sqlDatabase, tableName, v.Name())
		//	if err == nil {
		//		colMeta.columnLen = columnLen
		//	}
		//}

		m.columns[i] = colMeta
	}

	m = updateDefaultPrimaryKey(m)
	return m, nil
}
