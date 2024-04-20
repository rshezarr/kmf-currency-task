package currency

// Название таблицы
const CurrencyTableName = "R_CURRENCY"

// Запросы
const (
	SaveRates = `INSERT INTO %s (TITLE, CODE, VALUE, A_DATE) VALUES (@Title, @Code, @Value, @A_DATE);`

	GetRates = `SELECT TITLE, CODE, VALUE, A_DATE FROM %s WHERE A_DATE=@A_DATE;`
)
