package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// Surah names with English & Arabic titles
var surahNames = map[int]struct {
	English string
	Arabic  string
}{
	1:   {"Al-Fathiha", "الفاتحة"},
	2:   {"Al-Baqarah", "البقرة"},
	3:   {"Ali-Imran", "آل عمران"},
	4:   {"An-Nisa", "النساء"},
	5:   {"Al-Maeda", "المائدة"},
	6:   {"Al-An'am", "الأنعام"},
	7:   {"Al-A'raf", "الأعراف"},
	8:   {"Al-Anfal", "الأنفال"},
	9:   {"Ath-Taubah", "التوبة"},
	10:  {"Younus", "يونس"},
	11:  {"Hud", "هود"},
	12:  {"Yousuf", "يوسف"},
	13:  {"Ar-R'ad", "الرعد"},
	14:  {"Ibrahim", "إبراهيم"},
	15:  {"Al-Hijr", "الحجر"},
	16:  {"An-Nahl", "النحل"},
	17:  {"Al-Isra", "الإسراء"},
	18:  {"Al-Kahf", "الكهف"},
	19:  {"Maryam", "مريم"},
	20:  {"Taha", "طه"},
	21:  {"Al-Anbiya", "الأنبياء"},
	22:  {"Al-Hajj", "الحج"},
	23:  {"Al-Mu'minoon", "المؤمنون"},
	24:  {"An-Noor", "النور"},
	25:  {"Al-Furqan", "الفرقان"},
	26:  {"Ash-Shuara", "الشعراء"},
	27:  {"An-Naml", "النمل"},
	28:  {"Al-Qasas", "القصص"},
	29:  {"Al-Ankabut", "العنكبوت"},
	30:  {"Ar-Room", "الروم"},
	31:  {"Luqman", "لقمان"},
	32:  {"As-Sajdah", "السجدة"},
	33:  {"Al-Ahzab", "الأحزاب"},
	34:  {"Saba", "سبأ"},
	35:  {"Fathir", "فاطر"},
	36:  {"Yaseen", "يس"},
	37:  {"As-Saffat", "الصافات"},
	38:  {"Saad", "ص"},
	39:  {"Az-Zumar", "الزمر"},
	40:  {"Gafir", "غافر"},
	41:  {"Fussilat", "فصلت"},
	42:  {"Ash-Shoora", "الشورى"},
	43:  {"Az-Zukruf", "الزخرف"},
	44:  {"Ad-Dukhan", "الدخان"},
	45:  {"Al-Jathiyah", "الجاثية"},
	46:  {"Al-Ahqaf", "الأحقاف"},
	47:  {"Muhammad", "محمد"},
	48:  {"Al-Fath", "الفتح"},
	49:  {"Al-Hujurat", "الحجرات"},
	50:  {"Qaf", "ق"},
	51:  {"Adh-Dhariyat", "الذاريات"},
	52:  {"At-Toor", "الطور"},
	53:  {"An-Najm", "النجم"},
	54:  {"Al-Qamar", "القمر"},
	55:  {"Ar-Rahman", "الرحمن"},
	56:  {"Al-Waqiah", "الواقعة"},
	57:  {"Al-Hadeed", "الحديد"},
	58:  {"Al-Mujadalah", "المجادلة"},
	59:  {"Al-Hashr", "الحشر"},
	60:  {"Al-Mumthahinah", "الممتحنة"},
	61:  {"Al-Saff", "الصف"},
	62:  {"Al-Jumuah", "الجمعة"},
	63:  {"Al-Munafiqun", "المنافقون"},
	64:  {"At-Thagabun", "التغابن"},
	65:  {"At-Talaq", "الطلاق"},
	66:  {"At-Tahreem", "التحريم"},
	67:  {"Al-Mulk", "الملك"},
	68:  {"Al-Qalam", "القلم"},
	69:  {"Al-Haqqah", "الحاقة"},
	70:  {"Al-Ma'arij", "المعارج"},
	71:  {"Nuh", "نوح"},
	72:  {"Al-Jinn", "الجن"},
	73:  {"Al-Muzzammil", "المزمل"},
	74:  {"Al-Muddassir", "المدثر"},
	75:  {"Al-Qiamah", "القيامة"},
	76:  {"Al-Insan", "الإنسان"},
	77:  {"Al-Mursalat", "المرسلات"},
	78:  {"An-Naba", "النبأ"},
	79:  {"An-Naziath", "النازعات"},
	80:  {"Abasa", "عبس"},
	81:  {"At-Takwir", "التكوير"},
	82:  {"Al-Infitar", "الإنفطار"},
	83:  {"Al-Mutaffifeen", "المطففين"},
	84:  {"Al-Inshiqaq", "الإنشقاق"},
	85:  {"Al-Burooj", "البروج"},
	86:  {"At-Taariq", "الطارق"},
	87:  {"Al-A'la", "الأعلى"},
	88:  {"Al-Ghashiya", "الغاشية"},
	89:  {"Al-Fajr", "الفجر"},
	90:  {"Al-Balad", "البلد"},
	91:  {"Ash-Shams", "الشمس"},
	92:  {"Al-Lail", "الليل"},
	93:  {"Ad-Dhuha", "الضحى"},
	94:  {"Al-Inshirah", "الشرح"},
	95:  {"At-Teen", "التين"},
	96:  {"Al-Alaq", "العلق"},
	97:  {"Al-Qadr", "القدر"},
	98:  {"Al-Bayyinah", "البينة"},
	99:  {"Al-Zalzalah", "الزلزلة"},
	100: {"Al-Aadiyat", "العاديات"},
	101: {"Al-Qariah", "القارعة"},
	102: {"At-Thakathur", "التكاثر"},
	103: {"Al-Asr", "العصر"},
	104: {"Al-Humazah", "الهمزة"},
	105: {"Al-Fil", "الفيل"},
	106: {"Quraish", "قريش"},
	107: {"Al-Maun", "الماعون"},
	108: {"Al-Kauthar", "الكوثر"},
	109: {"Al-Kafiroon", "الكافرون"},
	110: {"An-Nasr", "النصر"},
	111: {"Al-Masad", "المسد"},
	112: {"Al-Ikhlas", "الإخلاص"},
	113: {"Al-Falaq", "الفلق"},
	114: {"An-Nas", "الناس"},
}

// Rename function using CLI parameters

func extractSurahNumber(filename string) (int, error) {
	re := regexp.MustCompile(`\d+`) // Find first number in filename
	match := re.FindString(filename)
	return strconv.Atoi(match)
}

// Renames files based on extracted Surah number
func RenameFiles(pattern string) {
	files, _ := filepath.Glob(pattern)
	for _, file := range files {
		index, err := extractSurahNumber(filepath.Base(file))
		if err != nil {
			log.Printf("Skipping %s (no valid number found)\n", file)
			continue
		}

		if surah, found := surahNames[index]; found {
			newName := fmt.Sprintf("%03d - %s - %s%s", index, surah.English, surah.Arabic, filepath.Ext(file))
			err := os.Rename(file, newName)
			if err != nil {
				log.Printf("Error renaming %s: %v\n", file, err)
			} else {
				fmt.Printf("Renamed: %s → %s\n", file, newName)
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run renamer.go '*.mp3'")
	}
	RenameFiles(os.Args[1])
}
