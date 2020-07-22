package kel80

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kel81() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
		(81001, 'NENOMANI', 6638, 475, 34),
		(81002, 'NILBO', 6638, 475, 34),
		(81003, 'NOFUALMA', 6638, 475, 34),
		(81004, 'NOHOLUOK', 6638, 475, 34),
		(81005, 'OBABIN', 6638, 475, 34),
		(81006, 'OBABIN SATU', 6638, 475, 34),
		(81007, 'PAMUMU', 6638, 475, 34),
		(81008, 'PIRIKALEM', 6638, 475, 34),
		(81009, 'SABILIRONGGO', 6638, 475, 34),
		(81010, 'SAHIKMA', 6638, 475, 34),
		(81011, 'SALO', 6638, 475, 34),
		(81012, 'SALOHE', 6638, 475, 34),
		(81013, 'SAPIWAREK', 6638, 475, 34),
		(81014, 'SEBI', 6638, 475, 34),
		(81015, 'SENGIKMA', 6638, 475, 34),
		(81016, 'SOHI', 6638, 475, 34),
		(81017, 'SOHOMBUNU', 6638, 475, 34),
		(81018, 'SOMBULE', 6638, 475, 34),
		(81019, 'SUAHE', 6638, 475, 34),
		(81020, 'SUELE', 6638, 475, 34),
		(81021, 'SUMBULE SATU', 6638, 475, 34),
		(81022, 'SUMINAIKMA', 6638, 475, 34),
		(81023, 'TAGABAGA', 6638, 475, 34),
		(81024, 'ULUHUFUK', 6638, 475, 34),
		(81025, 'ULUN', 6638, 475, 34),
		(81026, 'ULUSI', 6638, 475, 34),
		(81027, 'USABIYE', 6638, 475, 34),
		(81028, 'UWAMBO SATU', 6638, 475, 34),
		(81029, 'WABUHUK', 6638, 475, 34),
		(81030, 'WAGHASILIMO', 6638, 475, 34),
		(81031, 'WALAGIMA', 6638, 475, 34),
		(81032, 'WAMBAL', 6638, 475, 34),
		(81033, 'WAMBALFAK', 6638, 475, 34),
		(81034, 'WAMBO', 6638, 475, 34),
		(81035, 'WAMHOR', 6638, 475, 34),
		(81036, 'WANDIKBONGKION', 6638, 475, 34),
		(81037, 'WANGGIBO', 6638, 475, 34),
		(81038, 'WANGKUN', 6638, 475, 34),
		(81039, 'WAROHAM', 6638, 475, 34),
		(81040, 'WENEMIKMA', 6638, 475, 34),
		(81041, 'WILAK', 6638, 475, 34),
		(81042, 'WILEROMA', 6638, 475, 34),
		(81043, 'WILTLANGGO', 6638, 475, 34),
		(81044, 'WOROL', 6638, 475, 34),
		(81045, 'WURUPIKMA', 6638, 475, 34),
		(81046, 'WUTLARIN', 6638, 475, 34),
		(81047, 'YABUSUP', 6638, 475, 34),
		(81048, 'YAHATMA', 6638, 475, 34),
		(81049, 'YAMA', 6638, 475, 34),
		(81050, 'YAMBAIKMA', 6638, 475, 34),
		(81051, 'YANAMIK', 6638, 475, 34),
		(81052, 'YANGKIKALMA', 6638, 475, 34),
		(81053, 'YAWAN', 6638, 475, 34),
		(81054, 'YUTANGGO', 6638, 475, 34),
		(81055, 'ALIMUHUK', 6639, 475, 34),
		(81056, 'APALAPSILI', 6639, 475, 34),
		(81057, 'ASILIKMA', 6639, 475, 34),
		(81058, 'BABLIK', 6639, 475, 34),
		(81059, 'BENYAM', 6639, 475, 34),
		(81060, 'EAL', 6639, 475, 34),
		(81061, 'FALUKWALILO', 6639, 475, 34),
		(81062, 'FARI', 6639, 475, 34),
		(81063, 'FILIYAHIK', 6639, 475, 34),
		(81064, 'HABALO', 6639, 475, 34),
		(81065, 'HINONGKOAMBUT', 6639, 475, 34),
		(81066, 'HOLIK', 6639, 475, 34),
		(81067, 'HOLUAOGALMA', 6639, 475, 34),
		(81068, 'HOLUOKALEM', 6639, 475, 34),
		(81069, 'HUBLIKI', 6639, 475, 34),
		(81070, 'HUKALOPUNU', 6639, 475, 34),
		(81071, 'HUMALEN', 6639, 475, 34),
		(81072, 'IFKINDI', 6639, 475, 34),
		(81073, 'KAYOMA', 6639, 475, 34),
		(81074, 'KELAMPURIM', 6639, 475, 34),
		(81075, 'KENGGGEMBUN', 6639, 475, 34),
		(81076, 'KILAL', 6639, 475, 34),
		(81077, 'KINHE', 6639, 475, 34),
		(81078, 'KULET', 6639, 475, 34),
		(81079, 'KUNDIKELE', 6639, 475, 34),
		(81080, 'LAMBUKMU', 6639, 475, 34),
		(81081, 'LILINSALUK', 6639, 475, 34),
		(81082, 'MAKRIG', 6639, 475, 34),
		(81083, 'MALINWAREK', 6639, 475, 34),
		(81084, 'MOLIYINGGI', 6639, 475, 34),
		(81085, 'NANAHORUK', 6639, 475, 34),
		(81086, 'NASINENA', 6639, 475, 34),
		(81087, 'NASWER', 6639, 475, 34),
		(81088, 'NATOKSILI', 6639, 475, 34),
		(81089, 'NOHONIL', 6639, 475, 34),
		(81090, 'PIPISIM', 6639, 475, 34),
		(81091, 'PONG', 6639, 475, 34),
		(81092, 'SABILIKALEM', 6639, 475, 34),
		(81093, 'SABUALO', 6639, 475, 34),
		(81094, 'SIEN', 6639, 475, 34),
		(81095, 'SOBIKAMBUT', 6639, 475, 34),
		(81096, 'SUERELIHIM', 6639, 475, 34),
		(81097, 'TEMBUT', 6639, 475, 34),
		(81098, 'TEMBUT', 6639, 475, 34),
		(81099, 'TIKANO', 6639, 475, 34),
		(81100, 'WANAMALO', 6639, 475, 34),
		(81101, 'WASALALO', 6639, 475, 34),
		(81102, 'WIRALESILI', 6639, 475, 34),
		(81103, 'WIYUKWILIK', 6639, 475, 34),
		(81104, 'YAREMA', 6639, 475, 34),
		(81105, 'YAREMA SATU', 6639, 475, 34),
		(81106, 'AMPERA', 6640, 475, 34),
		(81107, 'BENAWA SATU', 6640, 475, 34),
		(81108, 'BURUKUKU', 6640, 475, 34),
		(81109, 'DUKUNASI', 6640, 475, 34),
		(81110, 'GILIKA', 6640, 475, 34),
		(81111, 'HIMI', 6640, 475, 34),
		(81112, 'IKON', 6640, 475, 34),
		(81113, 'IYAP', 6640, 475, 34),
		(81114, 'KAMIKA', 6640, 475, 34),
		(81115, 'KAPAWA', 6640, 475, 34),
		(81116, 'KARAMINA', 6640, 475, 34),
		(81117, 'KEY', 6640, 475, 34),
		(81118, 'KONOBUN', 6640, 475, 34),
		(81119, 'KUKDOMOL', 6640, 475, 34),
		(81120, 'KUTAKURUK', 6640, 475, 34),
		(81121, 'LAWE', 6640, 475, 34),
		(81122, 'LULUM', 6640, 475, 34),
		(81123, 'MAKRIP', 6640, 475, 34),
		(81124, 'MAKSARAM', 6640, 475, 34),
		(81125, 'MERSENG', 6640, 475, 34),
		(81126, 'NAIRA', 6640, 475, 34),
		(81127, 'NARAU', 6640, 475, 34),
		(81128, 'NIYAWI', 6640, 475, 34),
		(81129, 'PALUKE', 6640, 475, 34),
		(81130, 'PENSALE', 6640, 475, 34),
		(81131, 'PEPERA', 6640, 475, 34),
		(81132, 'PINAMPI', 6640, 475, 34),
		(81133, 'PISOL', 6640, 475, 34),
		(81134, 'PRITITI', 6640, 475, 34),
		(81135, 'SIWA', 6640, 475, 34),
		(81136, 'THAMAKSIN', 6640, 475, 34),
		(81137, 'TRIKORA', 6640, 475, 34),
		(81138, 'WELAP', 6640, 475, 34),
		(81139, 'WERMAS', 6640, 475, 34),
		(81140, 'WIBI', 6640, 475, 34),
		(81141, 'WIRSA', 6640, 475, 34),
		(81142, 'YAKWA', 6640, 475, 34),
		(81143, 'ALUIS', 6641, 475, 34),
		(81144, 'DUWONG', 6641, 475, 34),
		(81145, 'ELELIM SATU', 6641, 475, 34),
		(81146, 'ELILEM', 6641, 475, 34),
		(81147, 'EMON', 6641, 475, 34),
		(81148, 'EREGI', 6641, 475, 34),
		(81149, 'HABIYEAMBUT', 6641, 475, 34),
		(81150, 'HELABU', 6641, 475, 34),
		(81151, 'HESMAT', 6641, 475, 34),
		(81152, 'HOBAKMA', 6641, 475, 34),
		(81153, 'HOLANI', 6641, 475, 34),
		(81154, 'HONITA', 6641, 475, 34),
		(81155, 'ISILA', 6641, 475, 34),
		(81156, 'KERON', 6641, 475, 34),
		(81157, 'KWIKMA', 6641, 475, 34),
		(81158, 'LAPORA', 6641, 475, 34),
		(81159, 'MARIBU', 6641, 475, 34),
		(81160, 'MOLINGGI', 6641, 475, 34),
		(81161, 'MOMONHUSI', 6641, 475, 34),
		(81162, 'OHOBAM', 6641, 475, 34),
		(81163, 'OHONIAM', 6641, 475, 34),
		(81164, 'PANGKIK', 6641, 475, 34),
		(81165, 'PASIMBOLO', 6641, 475, 34),
		(81166, 'PIHEL', 6641, 475, 34),
		(81167, 'PIRIP', 6641, 475, 34),
		(81168, 'PISIREG', 6641, 475, 34),
		(81169, 'PUNGKAHIK', 6641, 475, 34),
		(81170, 'SILI', 6641, 475, 34),
		(81171, 'SIPSOI', 6641, 475, 34),
		(81172, 'SIPSON', 6641, 475, 34),
		(81173, 'SIRA', 6641, 475, 34),
		(81174, 'SOWI', 6641, 475, 34),
		(81175, 'SUBUIRUINGGAMA', 6641, 475, 34),
		(81176, 'TANABASI', 6641, 475, 34),
		(81177, 'UBI', 6641, 475, 34),
		(81178, 'ULO', 6641, 475, 34),
		(81179, 'WAGAGU', 6641, 475, 34),
		(81180, 'WALKEP', 6641, 475, 34),
		(81181, 'WARA', 6641, 475, 34),
		(81182, 'WARIKMA', 6641, 475, 34),
		(81183, 'WASUA', 6641, 475, 34),
		(81184, 'WEREKMA', 6641, 475, 34),
		(81185, 'YABEMA', 6641, 475, 34),
		(81186, 'YAKIKMA', 6641, 475, 34),
		(81187, 'AMBILIKI', 6642, 475, 34),
		(81188, 'AMPOLONGSILI', 6642, 475, 34),
		(81189, 'AMULUK', 6642, 475, 34),
		(81190, 'DIKOHOBARI', 6642, 475, 34),
		(81191, 'FEINGKAMA', 6642, 475, 34),
		(81192, 'FIKFAK', 6642, 475, 34),
		(81193, 'FOLONGSILI', 6642, 475, 34),
		(81194, 'FUNUI', 6642, 475, 34),
		(81195, 'HALIALO', 6642, 475, 34),
		(81196, 'HALISEK', 6642, 475, 34),
		(81197, 'HELEBOL', 6642, 475, 34),
		(81198, 'HILARIKI', 6642, 475, 34),
		(81199, 'HINDALIMUKUK', 6642, 475, 34),
		(81200, 'HOBUT', 6642, 475, 34),
		(81201, 'HOLOWI', 6642, 475, 34),
		(81202, 'IRAREK', 6642, 475, 34),
		(81203, 'KAMPOL', 6642, 475, 34),
		(81204, 'KAYALEM', 6642, 475, 34),
		(81205, 'KAYO', 6642, 475, 34),
		(81206, 'KOUM', 6642, 475, 34),
		(81207, 'LANGAM', 6642, 475, 34),
		(81208, 'LASIK', 6642, 475, 34),
		(81209, 'LOHOLAN', 6642, 475, 34),
		(81210, 'MABUALEM', 6642, 475, 34),
		(81211, 'MAMIOAN', 6642, 475, 34),
		(81212, 'MENEPINI', 6642, 475, 34),
		(81213, 'MOHOBIYE', 6642, 475, 34),
		(81214, 'MOHONU', 6642, 475, 34),
		(81215, 'MONTEK', 6642, 475, 34),
		(81216, 'NOHONIL', 6642, 475, 34),
		(81217, 'PAMI', 6642, 475, 34),
		(81218, 'PANAL', 6642, 475, 34),
		(81219, 'PANALULUN', 6642, 475, 34),
		(81220, 'PIRANG', 6642, 475, 34),
		(81221, 'PISANGGO', 6642, 475, 34),
		(81222, 'POHOLANGGEN', 6642, 475, 34),
		(81223, 'POIK', 6642, 475, 34),
		(81224, 'POIK SATU', 6642, 475, 34),
		(81225, 'SAKAM', 6642, 475, 34),
		(81226, 'SALEMA', 6642, 475, 34),
		(81227, 'SALY', 6642, 475, 34),
		(81228, 'SAMARIA', 6642, 475, 34),
		(81229, 'SEHEREK', 6642, 475, 34),
		(81230, 'SELEBI', 6642, 475, 34),
		(81231, 'SELEK', 6642, 475, 34),
		(81232, 'SILFAL', 6642, 475, 34),
		(81233, 'SINAHAL', 6642, 475, 34),
		(81234, 'SOHARAM', 6642, 475, 34),
		(81235, 'SUKALIMI', 6642, 475, 34),
		(81236, 'SUNTAM', 6642, 475, 34),
		(81237, 'TAHAMAK', 6642, 475, 34),
		(81238, 'TINMUHUK', 6642, 475, 34),
		(81239, 'UBALIHI', 6642, 475, 34),
		(81240, 'ULUM', 6642, 475, 34),
		(81241, 'WALINGKAPMA', 6642, 475, 34),
		(81242, 'WANGKULAMSAFIANG', 6642, 475, 34),
		(81243, 'WASUPAHIK', 6642, 475, 34),
		(81244, 'WELAREK', 6642, 475, 34),
		(81245, 'WELEAREKPUNU', 6642, 475, 34),
		(81246, 'WERENGGIK', 6642, 475, 34),
		(81247, 'WOMPOLI', 6642, 475, 34),
		(81248, 'YAHAMER', 6642, 475, 34)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 81")
}