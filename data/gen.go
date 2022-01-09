package data

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
)

var (
	// chinese               = []int64{19968, 40869}
	chinese               = []int64{19968, 20000}
	seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	cnameList = strings.Split(strings.ReplaceAll(cnames, `\n`, ""), " ")
	cnameLen  = len(cnameList)
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const cnames = `趙 錢 孫 李 周 吳 鄭 王 馮 陳 褚 衞 蔣 沈 韓 楊 朱 秦 尤 許 何 呂 施 張 孔 曹 嚴 華 金 魏 陶 姜 戚 謝 鄒 喻
柏 水 竇 章 雲 蘇 潘 葛 奚 範 彭 郎 魯 韋 昌 馬 苗 鳳 花 方 俞 任 袁 柳 酆 鮑 史 唐 費 廉 岑 薛 雷 賀 倪 湯 滕 殷 羅 畢 郝 鄔 安 常 樂 於 時 傅
皮 卞 齊 康 伍 餘 元 卜 顧 孟 平 黃 和 穆 蕭 尹 姚 邵 湛 汪 祁 毛 禹 狄 米 貝 明 臧 計 伏 成 戴 談 宋 茅 龐 熊 紀 舒 屈 項 祝 董 梁 杜 阮 藍 閔
席 季 麻 強 賈 路 婁 危 江 童 顏 郭 梅 盛 林 刁 鍾 徐 邱 駱 高 夏 蔡 田 樊 胡 凌 霍 虞 萬 支 柯 昝 管 盧 莫 柯 房 裘 繆 幹 解 應 宗 丁 宣 賁 鄧
鬱 單 杭 洪 包 諸 左 石 崔 吉 鈕 龔 程 嵇 邢 滑 裴 陸 榮 翁 荀 羊 於 惠 甄 曲 家 封 芮 羿 儲 靳 汲 邴 糜 松 井 段 富 巫 烏 焦 巴 弓 牧 隗 山 谷 車 侯 宓 蓬
全 郗 班 仰 秋 仲 伊 宮 寧 仇 欒 暴 甘 鈄 歷 戎 祖 武 符 劉 景 詹 束 龍 葉 幸 司 韶 郜 黎 薊 溥 印 宿 白 懷 蒲 邰 從 鄂 索 鹹 籍 賴 卓 藺 屠 蒙 池 喬 陽 鬱
胥 能 蒼 雙 聞 莘 黨 翟 譚 貢 勞 逄 姬 申 扶 堵 冉 宰 酈 雍 卻 璩 桑 桂 濮 牛 壽 通 邊 扈 燕 冀 浦 尚 農 温 別 莊 晏 柴 瞿 閻 充 慕 連 茹 習 宦 艾 魚 容
向 古 易 慎 戈 廖 庾 終 暨 居 衡 步 都 耿 滿 弘 匡 國 文 寇 廣 祿 闕 東 歐 殳 沃 利 蔚 越 夔 隆 師 鞏 厙 聶 晁 勾 敖 融 冷 訾 辛 闞 那 簡 饒 空
曾 毋 沙 乜 養 鞠 須 豐 巢 關 蒯 相 查 後 荊 紅 遊 竺 權 逮 盍 益 桓 公 万俟 司馬 上官 歐陽 夏侯 諸葛 聞人 東方 赫連 皇甫 尉遲 公羊 澹台 公冶
宗政 濮陽 淳于 單于 太叔 申屠 公孫 仲孫 軒轅 令狐 徐離 宇文 長孫 慕容`

type Employee struct {
	Account       string
	Chinesename   string
	Address       string
	Englishname   string
	DepartmentEng string
	DeptID        string
	Phone         string
}

//generate 100000 random Employee records
func GenerateEmployee(qty int) []Employee {
	var employees []Employee
	for i := 0; i < qty; i++ {
		employees = append(employees, Employee{
			Account:       petname.Generate(2, "."),
			Chinesename:   cnameList[RandInt(0, int64(cnameLen))] + generateRandomRune(2, chinese[0], chinese[1]),
			Address:       petname.Generate(4, "."),
			Englishname:   petname.Generate(2, "-"),
			DepartmentEng: petname.Generate(1, "-"),
			DeptID:        randString(6) + "-" + stringWithCharset(2, "1234567890"),
			Phone:         stringWithCharset(7, "1234567890-"),
		})
	}
	return employees
}

func BigHungString(emps []Employee) string {
	var buffer bytes.Buffer

	for _, emp := range emps {
		str := fmt.Sprintf("%-30s%-10s%-50s%-30s%-20s%-10s%-10s",
			emp.Account, []byte(emp.Chinesename), emp.Address, emp.Englishname,
			emp.DepartmentEng, emp.DeptID, emp.Phone)
		buffer.WriteString(str)
		for i := 0; i < 200-len(str); i++ {
			buffer.Write([]byte(" "))
		}
	}
	return buffer.String()
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randString(length int) string {
	return stringWithCharset(length, charset)
}

/// chinese name
func generateRandomRune(size int, start, end int64) string {
	randRune := make([]rune, size)
	for i := range randRune {
		randRune[i] = rune(RandInt(start, end))
	}
	return string(randRune)
}

func RandInt(start, end int64) int64 {
	return (start + rand.Int63n(end-start))
}
