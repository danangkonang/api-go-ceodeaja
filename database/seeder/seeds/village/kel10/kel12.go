package kel10

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kel12() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
		(12001, 'DUKUHSETI', 1015, 50, 4),
		(12002, 'DUMPIL', 1015, 50, 4),
		(12003, 'GROGOLAN', 1015, 50, 4),
		(12004, 'KEMBANG', 1015, 50, 4),
		(12005, 'KENANTI', 1015, 50, 4),
		(12006, 'NGAGEL', 1015, 50, 4),
		(12007, 'PUNCEL', 1015, 50, 4),
		(12008, 'TEGALOMBO', 1015, 50, 4),
		(12009, 'WEDUSAN', 1015, 50, 4),
		(12010, 'BABALAN', 1016, 50, 4),
		(12011, 'BANJAREJO', 1016, 50, 4),
		(12012, 'BANJARSARI', 1016, 50, 4),
		(12013, 'BENDOHARJO', 1016, 50, 4),
		(12014, 'BOGOTANJUNG', 1016, 50, 4),
		(12015, 'GABUS', 1016, 50, 4),
		(12016, 'GABUS', 1016, 50, 4),
		(12017, 'GEBANG', 1016, 50, 4),
		(12018, 'GEMPOLSARI', 1016, 50, 4),
		(12019, 'KALIPANG', 1016, 50, 4),
		(12020, 'KARABAN', 1016, 50, 4),
		(12021, 'KARANGREJO', 1016, 50, 4),
		(12022, 'KEYONGAN', 1016, 50, 4),
		(12023, 'KORIPANDRIYO', 1016, 50, 4),
		(12024, 'KORYOKALANGAN', 1016, 50, 4),
		(12025, 'KOSEKAN', 1016, 50, 4),
		(12026, 'MINTOBASUKI', 1016, 50, 4),
		(12027, 'MOJOLAWARAN', 1016, 50, 4),
		(12028, 'NGLINDUK', 1016, 50, 4),
		(12029, 'PANDANHARUM', 1016, 50, 4),
		(12030, 'PANTIREJO', 1016, 50, 4),
		(12031, 'PELEM', 1016, 50, 4),
		(12032, 'PENANGGUNGAN', 1016, 50, 4),
		(12033, 'PLUMBUNGAN', 1016, 50, 4),
		(12034, 'SAMBIREJO', 1016, 50, 4),
		(12035, 'SOKO', 1016, 50, 4),
		(12036, 'SUGIHREJO', 1016, 50, 4),
		(12037, 'SULURSARI', 1016, 50, 4),
		(12038, 'SUNGGINGWARNO', 1016, 50, 4),
		(12039, 'SUWATU', 1016, 50, 4),
		(12040, 'TAHUNAN', 1016, 50, 4),
		(12041, 'TAMBAHMULYO', 1016, 50, 4),
		(12042, 'TANJANG', 1016, 50, 4),
		(12043, 'TANJUNGANOM', 1016, 50, 4),
		(12044, 'TLOGOAYU', 1016, 50, 4),
		(12045, 'TLOGOTIRTO', 1016, 50, 4),
		(12046, 'TUNGGULREJO', 1016, 50, 4),
		(12047, 'WUWUR', 1016, 50, 4),
		(12048, 'BAGENG', 1017, 50, 4),
		(12049, 'BERMI', 1017, 50, 4),
		(12050, 'GEMBONG', 1017, 50, 4),
		(12051, 'KEDUNGBULUS', 1017, 50, 4),
		(12052, 'KETANGGAN', 1017, 50, 4),
		(12053, 'KLAKAH KASIAN', 1017, 50, 4),
		(12054, 'PLUKARAN', 1017, 50, 4),
		(12055, 'POHGADING', 1017, 50, 4),
		(12056, 'SEMIREJO', 1017, 50, 4),
		(12057, 'SITILUHUR', 1017, 50, 4),
		(12058, 'WONOSEKAR', 1017, 50, 4),
		(12059, 'BANCAK', 1018, 50, 4),
		(12060, 'GADU', 1018, 50, 4),
		(12061, 'GAJIHAN', 1018, 50, 4),
		(12062, 'GILING', 1018, 50, 4),
		(12063, 'GULANGPUNGGE', 1018, 50, 4),
		(12064, 'GUNUNGWUNGKAL', 1018, 50, 4),
		(12065, 'JEMBULWUNUT', 1018, 50, 4),
		(12066, 'JEPALO', 1018, 50, 4),
		(12067, 'JRAHI', 1018, 50, 4),
		(12068, 'NGETUK', 1018, 50, 4),
		(12069, 'PERDOPO', 1018, 50, 4),
		(12070, 'PESAGEN', 1018, 50, 4),
		(12071, 'SAMPOK', 1018, 50, 4),
		(12072, 'SIDOMULYO', 1018, 50, 4),
		(12073, 'SUMBEREJO', 1018, 50, 4),
		(12074, 'AROMANIS (ARUMANIS)', 1019, 50, 4),
		(12075, 'BOTO', 1019, 50, 4),
		(12076, 'KEBONTURI', 1019, 50, 4),
		(12077, 'LUNDO', 1019, 50, 4),
		(12078, 'MANJANG', 1019, 50, 4),
		(12079, 'MANTINGAN', 1019, 50, 4),
		(12080, 'MOJOLAMPIR', 1019, 50, 4),
		(12081, 'MOJOLUHUR', 1019, 50, 4),
		(12082, 'RONGGO', 1019, 50, 4),
		(12083, 'SIDOLUHUR', 1019, 50, 4),
		(12084, 'SIDOMUKTI', 1019, 50, 4),
		(12085, 'SRIKATON', 1019, 50, 4),
		(12086, 'SRIWEDARI', 1019, 50, 4),
		(12087, 'SUKORUKUN', 1019, 50, 4),
		(12088, 'SUMBERAGUNG', 1019, 50, 4),
		(12089, 'SUMBERAN', 1019, 50, 4),
		(12090, 'SUMBERARUM', 1019, 50, 4),
		(12091, 'SUMBERREJO', 1019, 50, 4),
		(12092, 'TAMANSARI', 1019, 50, 4),
		(12093, 'TEGALARUM', 1019, 50, 4),
		(12094, 'TRIKOYO', 1019, 50, 4),
		(12095, 'BUNGASREJO', 1020, 50, 4),
		(12096, 'DUKUHMULYO', 1020, 50, 4),
		(12097, 'GLONGGONG', 1020, 50, 4),
		(12098, 'JAKENAN', 1020, 50, 4),
		(12099, 'JATISARI', 1020, 50, 4),
		(12100, 'KALIMULYO', 1020, 50, 4),
		(12101, 'KARANGREJO LOR', 1020, 50, 4),
		(12102, 'KARANGROWO', 1020, 50, 4),
		(12103, 'KEDUNGMULYO', 1020, 50, 4),
		(12104, 'MANTINGAN TENGAH', 1020, 50, 4),
		(12105, 'NGASTOREJO', 1020, 50, 4),
		(12106, 'PLOSOJENAR', 1020, 50, 4),
		(12107, 'PULUHAN TENGAH', 1020, 50, 4),
		(12108, 'SEMBATURAGUNG', 1020, 50, 4),
		(12109, 'SENDANGSOKO', 1020, 50, 4),
		(12110, 'SIDOARUM', 1020, 50, 4),
		(12111, 'SIDOMULYO', 1020, 50, 4),
		(12112, 'SONOREJO', 1020, 50, 4),
		(12113, 'TAMBAHMULYO', 1020, 50, 4),
		(12114, 'TANJUNGSARI', 1020, 50, 4),
		(12115, 'TLOGOREJO', 1020, 50, 4),
		(12116, 'TONDOKERTO', 1020, 50, 4),
		(12117, 'TONDOMULYO', 1020, 50, 4),
		(12118, 'AGUNGMULYO', 1021, 50, 4),
		(12119, 'BAJOMULYO', 1021, 50, 4),
		(12120, 'BAKARAN KULON', 1021, 50, 4),
		(12121, 'BAKARAN WETAN', 1021, 50, 4),
		(12122, 'BENDAR', 1021, 50, 4),
		(12123, 'BRINGIN', 1021, 50, 4),
		(12124, 'BUMIREJO', 1021, 50, 4),
		(12125, 'DOROPAYUNG', 1021, 50, 4),
		(12126, 'DUKUTALIT', 1021, 50, 4),
		(12127, 'GADINGREJO', 1021, 50, 4),
		(12128, 'GENENGMULYO', 1021, 50, 4),
		(12129, 'GROWONG KIDUL', 1021, 50, 4),
		(12130, 'GROWONG LOR', 1021, 50, 4),
		(12131, 'JEPURO', 1021, 50, 4),
		(12132, 'KARANG', 1021, 50, 4),
		(12133, 'KARANGREJO', 1021, 50, 4),
		(12134, 'KAUMAN', 1021, 50, 4),
		(12135, 'KEBONSAWAHAN', 1021, 50, 4),
		(12136, 'KEDUNGPANCING', 1021, 50, 4),
		(12137, 'KETIP', 1021, 50, 4),
		(12138, 'KUDUKERAS', 1021, 50, 4),
		(12139, 'LANGENHARJO', 1021, 50, 4),
		(12140, 'MARGOMULYO', 1021, 50, 4),
		(12141, 'MINTOMULYO', 1021, 50, 4),
		(12142, 'PAJEKSAN', 1021, 50, 4),
		(12143, 'PEKUWON', 1021, 50, 4),
		(12144, 'SEJOMULYO', 1021, 50, 4),
		(12145, 'TLUWAH', 1021, 50, 4),
		(12146, 'TRIMULYO', 1021, 50, 4),
		(12147, 'BEKETEL', 1022, 50, 4),
		(12148, 'BOLOAGUNG', 1022, 50, 4),
		(12149, 'BRATI', 1022, 50, 4),
		(12150, 'DURENSAWIT', 1022, 50, 4),
		(12151, 'JATIROTO', 1022, 50, 4),
		(12152, 'JIMBARAN', 1022, 50, 4),
		(12153, 'KAYEN', 1022, 50, 4),
		(12154, 'PASURUHAN', 1022, 50, 4),
		(12155, 'PESAGI', 1022, 50, 4),
		(12156, 'PURWOKERTO', 1022, 50, 4),
		(12157, 'ROGOMULYO', 1022, 50, 4),
		(12158, 'SLUNGKEP', 1022, 50, 4),
		(12159, 'SRIKATON', 1022, 50, 4),
		(12160, 'SUMBERSARI', 1022, 50, 4),
		(12161, 'SUNDOLUHUR', 1022, 50, 4),
		(12162, 'TALUN', 1022, 50, 4),
		(12163, 'TRIMULYO', 1022, 50, 4),
		(12164, 'BADEGAN', 1023, 50, 4),
		(12165, 'BANYUURIP', 1023, 50, 4),
		(12166, 'BUMIREJO', 1023, 50, 4),
		(12167, 'DADIREJO', 1023, 50, 4),
		(12168, 'JAMBEAN KIDUL', 1023, 50, 4),
		(12169, 'JIMBARAN', 1023, 50, 4),
		(12170, 'LANGENHARJO', 1023, 50, 4),
		(12171, 'LANGSE', 1023, 50, 4),
		(12172, 'MARGOREJO', 1023, 50, 4),
		(12173, 'METARAMAN', 1023, 50, 4),
		(12174, 'MUKTIHARJO', 1023, 50, 4),
		(12175, 'NGAWEN', 1023, 50, 4),
		(12176, 'PEGANDAN', 1023, 50, 4),
		(12177, 'PENAMBUHAN', 1023, 50, 4),
		(12178, 'SOKOKULON', 1023, 50, 4),
		(12179, 'SUKOBUBUK', 1023, 50, 4),
		(12180, 'SUKOHARJO', 1023, 50, 4),
		(12181, 'WANGUNREJO', 1023, 50, 4),
		(12182, 'BULUMANIS KIDUL', 1024, 50, 4),
		(12183, 'BULUMANIS LOR', 1024, 50, 4),
		(12184, 'CEBOLEK KIDUL', 1024, 50, 4),
		(12185, 'KAJEN', 1024, 50, 4),
		(12186, 'KERTOMULYO', 1024, 50, 4),
		(12187, 'LANGGENHARJO', 1024, 50, 4),
		(12188, 'MARGOTUHU KIDUL', 1024, 50, 4),
		(12189, 'MARGOYOSO', 1024, 50, 4),
		(12190, 'NGEMPLAK KIDUL', 1024, 50, 4),
		(12191, 'NGEMPLAK LOR', 1024, 50, 4),
		(12192, 'PANGKALAN', 1024, 50, 4),
		(12193, 'POHIJO', 1024, 50, 4),
		(12194, 'PURWODADI', 1024, 50, 4),
		(12195, 'PURWOREJO', 1024, 50, 4),
		(12196, 'SEKARJALAK', 1024, 50, 4),
		(12197, 'SEMERAK', 1024, 50, 4),
		(12198, 'SIDOMUKTI', 1024, 50, 4),
		(12199, 'SONEYAN', 1024, 50, 4),
		(12200, 'TANJUNGREJO', 1024, 50, 4),
		(12201, 'TEGALARUM', 1024, 50, 4),
		(12202, 'TUNJUNGREJO', 1024, 50, 4),
		(12203, 'WATUROYO', 1024, 50, 4),
		(12204, 'BLARU', 1025, 50, 4),
		(12205, 'DENGKEK', 1025, 50, 4),
		(12206, 'GAJAHMATI', 1025, 50, 4),
		(12207, 'GERITAN', 1025, 50, 4),
		(12208, 'KALIDORO', 1025, 50, 4),
		(12209, 'KUTOHARJO', 1025, 50, 4),
		(12210, 'MULYOHARJO', 1025, 50, 4),
		(12211, 'MUSTOKOHARJO', 1025, 50, 4),
		(12212, 'NGARUS', 1025, 50, 4),
		(12213, 'NGEPUNGROJO', 1025, 50, 4),
		(12214, 'PANJUNAN', 1025, 50, 4),
		(12215, 'PARENGGAN', 1025, 50, 4),
		(12216, 'PATI KIDUL', 1025, 50, 4),
		(12217, 'PATI LOR', 1025, 50, 4),
		(12218, 'PATI WETAN', 1025, 50, 4),
		(12219, 'PAYANG', 1025, 50, 4),
		(12220, 'PLANGITAN', 1025, 50, 4),
		(12221, 'PURI', 1025, 50, 4),
		(12222, 'PURWOREJO', 1025, 50, 4),
		(12223, 'SARIREJO', 1025, 50, 4),
		(12224, 'SEMAMPIR', 1025, 50, 4),
		(12225, 'SIDOHARJO', 1025, 50, 4),
		(12226, 'SIDOKERTO', 1025, 50, 4),
		(12227, 'SINOMAN', 1025, 50, 4),
		(12228, 'SUGIHARJO', 1025, 50, 4),
		(12229, 'TAMBAHARJO', 1025, 50, 4),
		(12230, 'TAMBAHSARI', 1025, 50, 4),
		(12231, 'WIDOROKANDANG', 1025, 50, 4),
		(12232, 'WINONG', 1025, 50, 4),
		(12233, 'BODEH', 1026, 50, 4),
		(12234, 'GROGOLSARI', 1026, 50, 4),
		(12235, 'JETAK', 1026, 50, 4),
		(12236, 'KARANGREJO', 1026, 50, 4),
		(12237, 'KARANGWOTAN', 1026, 50, 4),
		(12238, 'KEPOHKENCONO', 1026, 50, 4),
		(12239, 'KLETEK', 1026, 50, 4),
		(12240, 'LUMBUNGMAS', 1026, 50, 4),
		(12241, 'MENCON', 1026, 50, 4),
		(12242, 'MOJOAGUNG', 1026, 50, 4),
		(12243, 'PELEMGEDE', 1026, 50, 4),
		(12244, 'PLOSOREJO', 1026, 50, 4),
		(12245, 'PUCAKWANGI', 1026, 50, 4),
		(12246, 'SITIMULYO', 1026, 50, 4),
		(12247, 'SOKOPULUHAN', 1026, 50, 4),
		(12248, 'TANJUNGSEKAR', 1026, 50, 4),
		(12249, 'TEGALWERO', 1026, 50, 4),
		(12250, 'TERTEG', 1026, 50, 4),
		(12251, 'TRIGUNO', 1026, 50, 4),
		(12252, 'WATESHAJI', 1026, 50, 4),
		(12253, 'ANGKATAN KIDUL', 1027, 50, 4),
		(12254, 'ANGKATAN LOR', 1027, 50, 4),
		(12255, 'KARANGAWEN', 1027, 50, 4),
		(12256, 'KARANGMULYO', 1027, 50, 4),
		(12257, 'KARANGWONO', 1027, 50, 4),
		(12258, 'KEBEN', 1027, 50, 4),
		(12259, 'KEDALINGAN', 1027, 50, 4),
		(12260, 'LARANGAN', 1027, 50, 4),
		(12261, 'MAITAN', 1027, 50, 4),
		(12262, 'MANGUNREKSO', 1027, 50, 4),
		(12263, 'MOJOMULYO', 1027, 50, 4),
		(12264, 'PAKIS', 1027, 50, 4),
		(12265, 'SINOMWIDODO', 1027, 50, 4),
		(12266, 'SITIREJO', 1027, 50, 4),
		(12267, 'TAMBAHAGUNG', 1027, 50, 4),
		(12268, 'TAMBAHARJO', 1027, 50, 4),
		(12269, 'TAMBAKROMO', 1027, 50, 4),
		(12270, 'WUKIRSARI', 1027, 50, 4),
		(12271, 'BENDOKATON KIDUL', 1028, 50, 4),
		(12272, 'BULUNGAN', 1028, 50, 4),
		(12273, 'DOROREJO', 1028, 50, 4),
		(12274, 'JEPAT KIDUL', 1028, 50, 4),
		(12275, 'JEPAT LOR', 1028, 50, 4),
		(12276, 'KALIKALONG', 1028, 50, 4),
		(12277, 'KEBOROMO', 1028, 50, 4),
		(12278, 'KEDUNGBANG', 1028, 50, 4),
		(12279, 'KEDUNGSARI', 1028, 50, 4),
		(12280, 'LUWANG', 1028, 50, 4),
		(12281, 'MARGOMULYO', 1028, 50, 4),
		(12282, 'PAKIS', 1028, 50, 4),
		(12283, 'PONDOWAN', 1028, 50, 4),
		(12284, 'PUNDENREJO', 1028, 50, 4),
		(12285, 'PURWOKERTO', 1028, 50, 4),
		(12286, 'SAMBIROTO', 1028, 50, 4),
		(12287, 'SENDANGREJO', 1028, 50, 4),
		(12288, 'TAYUKULON', 1028, 50, 4),
		(12289, 'TAYUWETAN', 1028, 50, 4),
		(12290, 'TENDAS', 1028, 50, 4),
		(12291, 'TUNGGULSARI', 1028, 50, 4),
		(12292, 'CABAK', 1029, 50, 4),
		(12293, 'GUNUNGSARI', 1029, 50, 4),
		(12294, 'GUWO', 1029, 50, 4),
		(12295, 'KLUMPIT', 1029, 50, 4),
		(12296, 'LAHAR', 1029, 50, 4),
		(12297, 'PURWOSARI', 1029, 50, 4),
		(12298, 'REGALOH', 1029, 50, 4),
		(12299, 'SAMBIREJO', 1029, 50, 4),
		(12300, 'SUMBERMULYO', 1029, 50, 4),
		(12301, 'SUWATU', 1029, 50, 4),
		(12302, 'TAJUNGSARI', 1029, 50, 4),
		(12303, 'TAMANSARI', 1029, 50, 4),
		(12304, 'TLOGOREJO', 1029, 50, 4),
		(12305, 'TLOGOSARI', 1029, 50, 4),
		(12306, 'WONOREJO', 1029, 50, 4),
		(12307, 'ASEMPAPAN', 1030, 50, 4),
		(12308, 'GUYANGAN', 1030, 50, 4),
		(12309, 'KADILANGU', 1030, 50, 4),
		(12310, 'KAJAR', 1030, 50, 4),
		(12311, 'KARANGLEGI', 1030, 50, 4),
		(12312, 'KARANGWAGE', 1030, 50, 4),
		(12313, 'KERTOMULYO', 1030, 50, 4),
		(12314, 'KETANEN', 1030, 50, 4),
		(12315, 'KRANDAN', 1030, 50, 4),
		(12316, 'MOJOAGUNG', 1030, 50, 4),
		(12317, 'PASUCEN', 1030, 50, 4),
		(12318, 'REJOAGUNG', 1030, 50, 4),
		(12319, 'SAMBILAWANG', 1030, 50, 4),
		(12320, 'TEGALHARJO', 1030, 50, 4),
		(12321, 'TLUTUP', 1030, 50, 4),
		(12322, 'TRANGKIL', 1030, 50, 4),
		(12323, 'BANGSALREJO', 1031, 50, 4),
		(12324, 'BUMIAYU', 1031, 50, 4),
		(12325, 'JATIMULYO', 1031, 50, 4),
		(12326, 'JETAK', 1031, 50, 4),
		(12327, 'JONTRO', 1031, 50, 4),
		(12328, 'KEPOH', 1031, 50, 4),
		(12329, 'MARGOREJO', 1031, 50, 4),
		(12330, 'NGURENREJO', 1031, 50, 4),
		(12331, 'NGURENSITI', 1031, 50, 4),
		(12332, 'PAGERHARJO', 1031, 50, 4),
		(12333, 'PANGGUNGROYOM', 1031, 50, 4),
		(12334, 'SIDOHARJO', 1031, 50, 4),
		(12335, 'SUKOHARJO', 1031, 50, 4),
		(12336, 'SUWADUK', 1031, 50, 4),
		(12337, 'TAWANGHARJO', 1031, 50, 4),
		(12338, 'TLOGOHARUM', 1031, 50, 4),
		(12339, 'TLUWUK', 1031, 50, 4),
		(12340, 'WEDARIJAKSA', 1031, 50, 4),
		(12341, 'BLINGIJATI', 1032, 50, 4),
		(12342, 'BRINGINWARENG', 1032, 50, 4),
		(12343, 'BUMIHARJO', 1032, 50, 4),
		(12344, 'DANYANGMULYO', 1032, 50, 4),
		(12345, 'DEGAN', 1032, 50, 4),
		(12346, 'GODO', 1032, 50, 4),
		(12347, 'GUNUNGPANTI', 1032, 50, 4),
		(12348, 'GUYANGAN', 1032, 50, 4),
		(12349, 'KARANGKONANG', 1032, 50, 4),
		(12350, 'KARANGSUMBER', 1032, 50, 4),
		(12351, 'KEBOLAMPANG', 1032, 50, 4),
		(12352, 'KEBOWAN', 1032, 50, 4),
		(12353, 'KLECOREGONANG', 1032, 50, 4),
		(12354, 'KROPAK', 1032, 50, 4),
		(12355, 'KUDUR', 1032, 50, 4),
		(12356, 'MINTORAHAYU', 1032, 50, 4),
		(12357, 'PADANGAN', 1032, 50, 4),
		(12358, 'PAGENDISAN', 1032, 50, 4),
		(12359, 'PEKALONGAN', 1032, 50, 4),
		(12360, 'POHGADING', 1032, 50, 4),
		(12361, 'PULOREJO', 1032, 50, 4),
		(12362, 'SARIMULYO', 1032, 50, 4),
		(12363, 'SERUTSADANG', 1032, 50, 4),
		(12364, 'SUGIHAN', 1032, 50, 4),
		(12365, 'SUMBERMULYO', 1032, 50, 4),
		(12366, 'TANGGEL', 1032, 50, 4),
		(12367, 'TAWANGREJO', 1032, 50, 4),
		(12368, 'TLOGOREJO', 1032, 50, 4),
		(12369, 'WINONG', 1032, 50, 4),
		(12370, 'WIRUN', 1032, 50, 4),
		(12371, 'BLIGO', 1033, 51, 4),
		(12372, 'COPRAYAN', 1033, 51, 4),
		(12373, 'KERTIJAYAN', 1033, 51, 4),
		(12374, 'PAKUMBULAN', 1033, 51, 4),
		(12375, 'PAWEDEN', 1033, 51, 4),
		(12376, 'SAPUGARUT', 1033, 51, 4),
		(12377, 'SIMBANG KULON', 1033, 51, 4),
		(12378, 'SIMBANG WETAN', 1033, 51, 4),
		(12379, 'WATUSALAM', 1033, 51, 4),
		(12380, 'WONOYOSO', 1033, 51, 4),
		(12381, 'BLIGOREJO', 1034, 51, 4),
		(12382, 'DORO', 1034, 51, 4),
		(12383, 'DOROREJO', 1034, 51, 4),
		(12384, 'HARJOSARI', 1034, 51, 4),
		(12385, 'KALIMOJOSARI', 1034, 51, 4),
		(12386, 'KUTOSARI', 1034, 51, 4),
		(12387, 'LARIKAN', 1034, 51, 4),
		(12388, 'LEMAH ABANG', 1034, 51, 4),
		(12389, 'PUNGANGAN', 1034, 51, 4),
		(12390, 'RANDUSARI', 1034, 51, 4),
		(12391, 'ROGOSELO', 1034, 51, 4),
		(12392, 'SAWANGAN', 1034, 51, 4),
		(12393, 'SIDOHARJO', 1034, 51, 4),
		(12394, 'WRINGIN AGUNG', 1034, 51, 4),
		(12395, 'BRENGKOLANG', 1035, 51, 4),
		(12396, 'GANDARUM', 1035, 51, 4),
		(12397, 'GEJLIG', 1035, 51, 4),
		(12398, 'KAJEN', 1035, 51, 4),
		(12399, 'KAJONGAN', 1035, 51, 4),
		(12400, 'KALIJOYO', 1035, 51, 4),
		(12401, 'KEBON AGUNG', 1035, 51, 4),
		(12402, 'KUTOREJO', 1035, 51, 4),
		(12403, 'KUTOROJO', 1035, 51, 4),
		(12404, 'LINGGOASRI', 1035, 51, 4),
		(12405, 'NYAMOK', 1035, 51, 4),
		(12406, 'PEKIRINGAN AGENG', 1035, 51, 4),
		(12407, 'PEKIRINGAN ALIT', 1035, 51, 4),
		(12408, 'PRINGSURAT', 1035, 51, 4),
		(12409, 'ROWOLAKU', 1035, 51, 4),
		(12410, 'SABARWANGI', 1035, 51, 4),
		(12411, 'SALIT', 1035, 51, 4),
		(12412, 'SAMBIROTO', 1035, 51, 4),
		(12413, 'SANGKAN JOYO', 1035, 51, 4),
		(12414, 'SINANGOH PRENDENG', 1035, 51, 4),
		(12415, 'SOKOYOSO', 1035, 51, 4),
		(12416, 'TAMBAKROTO', 1035, 51, 4),
		(12417, 'TANJUNG KULON', 1035, 51, 4),
		(12418, 'TANJUNGSARI', 1035, 51, 4),
		(12419, 'WONOREJO', 1035, 51, 4),
		(12420, 'BODAS', 1036, 51, 4),
		(12421, 'BOJONGKONENG', 1036, 51, 4),
		(12422, 'BUBAK', 1036, 51, 4),
		(12423, 'GARUNGWIYORO', 1036, 51, 4),
		(12424, 'GEMBONG', 1036, 51, 4),
		(12425, 'KANDANGSERANG', 1036, 51, 4),
		(12426, 'KARANGGONDANG', 1036, 51, 4),
		(12427, 'KLESEM', 1036, 51, 4),
		(12428, 'LAMBUR', 1036, 51, 4),
		(12429, 'LURAGUNG', 1036, 51, 4),
		(12430, 'SUKOHARJO', 1036, 51, 4),
		(12431, 'TAJUR', 1036, 51, 4),
		(12432, 'TRAJUMAS', 1036, 51, 4),
		(12433, 'WANGKELANG', 1036, 51, 4),
		(12434, 'JREBENG KEMBANG', 1037, 51, 4),
		(12435, 'KALIGAWE', 1037, 51, 4),
		(12436, 'KALILEMBU', 1037, 51, 4),
		(12437, 'KARANGDADAP', 1037, 51, 4),
		(12438, 'KEBONROWO PUCANG', 1037, 51, 4),
		(12439, 'KEBONSARI', 1037, 51, 4),
		(12440, 'KEDUNG KEBO', 1037, 51, 4),
		(12441, 'LOGANDENG', 1037, 51, 4),
		(12442, 'PAGUMENGAN MAS', 1037, 51, 4),
		(12443, 'PANGKAH', 1037, 51, 4),
		(12444, 'PEGANDON', 1037, 51, 4),
		(12445, 'AMBOKEMBANG', 1038, 51, 4),
		(12446, 'BUGANGAN/BUNGANGAN', 1038, 51, 4),
		(12447, 'KARANGDOWO', 1038, 51, 4),
		(12448, 'KEDUNG PATANGEWU', 1038, 51, 4),
		(12449, 'KEDUNGWUNI BARAT', 1038, 51, 4),
		(12450, 'KEDUNGWUNI TIMUR', 1038, 51, 4),
		(12451, 'KWAYANGAN', 1038, 51, 4),
		(12452, 'LANGKAP', 1038, 51, 4),
		(12453, 'PAJOMBLANGAN', 1038, 51, 4),
		(12454, 'PAKISPUTIH', 1038, 51, 4),
		(12455, 'PEKAJANGAN', 1038, 51, 4),
		(12456, 'PODO', 1038, 51, 4),
		(12457, 'PROTO', 1038, 51, 4),
		(12458, 'RENGAS', 1038, 51, 4),
		(12459, 'ROWOCACING', 1038, 51, 4),
		(12460, 'SALAKBROJO', 1038, 51, 4),
		(12461, 'TANGKIL KULON', 1038, 51, 4),
		(12462, 'TANGKIL TENGAH', 1038, 51, 4),
		(12463, 'TOSARAN', 1038, 51, 4),
		(12464, 'BRONDONG', 1039, 51, 4),
		(12465, 'JAGUNG', 1039, 51, 4),
		(12466, 'KAIBAHAN', 1039, 51, 4),
		(12467, 'KALIMADE', 1039, 51, 4),
		(12468, 'KARANGREJO', 1039, 51, 4),
		(12469, 'KARYOMUKTI', 1039, 51, 4),
		(12470, 'KESESI', 1039, 51, 4),
		(12471, 'KRANDON', 1039, 51, 4),
		(12472, 'KWASEN', 1039, 51, 4),
		(12473, 'KWIGARAN', 1039, 51, 4),
		(12474, 'LANGENSARI', 1039, 51, 4),
		(12475, 'MULYOREJO', 1039, 51, 4),
		(12476, 'PANTIREJO', 1039, 51, 4),
		(12477, 'PODOSARI', 1039, 51, 4),
		(12478, 'PONOLAWEN', 1039, 51, 4),
		(12479, 'SIDOMULYO', 1039, 51, 4),
		(12480, 'SIDOSARI', 1039, 51, 4),
		(12481, 'SRINAHAN', 1039, 51, 4),
		(12482, 'SUKOREJO', 1039, 51, 4),
		(12483, 'UJUNG NEGORO', 1039, 51, 4),
		(12484, 'WATUGAJAH', 1039, 51, 4),
		(12485, 'WATUPAYUNG', 1039, 51, 4),
		(12486, 'WINDUROJO', 1039, 51, 4),
		(12487, 'BANTAR KULON', 1040, 51, 4),
		(12488, 'DEPOK', 1040, 51, 4),
		(12489, 'KAPUNDUTAN', 1040, 51, 4),
		(12490, 'KUTOREMBET', 1040, 51, 4),
		(12491, 'LEBAKBARANG', 1040, 51, 4),
		(12492, 'MENDOLO', 1040, 51, 4),
		(12493, 'PAMUTUH', 1040, 51, 4),
		(12494, 'SIDOMULYO', 1040, 51, 4),
		(12495, 'TEMBELANG GUNUNG', 1040, 51, 4),
		(12496, 'TIMBANGSARI', 1040, 51, 4),
		(12497, 'WONOSIDO', 1040, 51, 4),
		(12498, 'BEDAGUNG', 1041, 51, 4),
		(12499, 'BOTOSARI', 1041, 51, 4),
		(12500, 'DOMIYANG', 1041, 51, 4),
		(12501, 'KALIBOJA', 1041, 51, 4),
		(12502, 'KALIOMBO', 1041, 51, 4),
		(12503, 'KRANDEGAN', 1041, 51, 4),
		(12504, 'LAMBANGGELUN', 1041, 51, 4),
		(12505, 'LOMENENG', 1041, 51, 4),
		(12506, 'NOTOGIWANG', 1041, 51, 4),
		(12507, 'PANINGGARAN', 1041, 51, 4),
		(12508, 'SAWANGAN', 1041, 51, 4),
		(12509, 'TANGGERAN', 1041, 51, 4),
		(12510, 'TENOGO', 1041, 51, 4),
		(12511, 'WERDI', 1041, 51, 4),
		(12512, 'WINDUAJI', 1041, 51, 4),
		(12513, 'BENDAN', 1042, 51, 4),
		(12514, 'BUMIREJO', 1042, 51, 4),
		(12515, 'KEBULEN', 1042, 51, 4),
		(12516, 'KERGON', 1042, 51, 4),
		(12517, 'KRAMATSARI', 1042, 51, 4),
		(12518, 'KRATON KIDUL', 1042, 51, 4),
		(12519, 'MEDONO', 1042, 51, 4),
		(12520, 'PASIRSARI', 1042, 51, 4),
		(12521, 'PODOSUGIH', 1042, 51, 4),
		(12522, 'PRINGLANGU', 1042, 51, 4),
		(12523, 'SAPURO', 1042, 51, 4),
		(12524, 'TEGALREJO', 1042, 51, 4),
		(12525, 'TIRTO', 1042, 51, 4),
		(12526, 'BANYURIP AGENG', 1043, 51, 4),
		(12527, 'BANYURIP ALIT', 1043, 51, 4),
		(12528, 'BUARAN', 1043, 51, 4),
		(12529, 'DUWET', 1043, 51, 4),
		(12530, 'JENGGOT', 1043, 51, 4),
		(12531, 'KERTOHARJO', 1043, 51, 4),
		(12532, 'KRADENAN', 1043, 51, 4),
		(12533, 'KURIPAN KIDUL', 1043, 51, 4),
		(12534, 'KURIPAN LOR', 1043, 51, 4),
		(12535, 'SOKO', 1043, 51, 4),
		(12536, 'YOSOREJO', 1043, 51, 4),
		(12537, 'BAROS', 1044, 51, 4),
		(12538, 'DEKORO', 1044, 51, 4),
		(12539, 'GAMER', 1044, 51, 4),
		(12540, 'KARANGMALANG', 1044, 51, 4),
		(12541, 'KAUMAN', 1044, 51, 4),
		(12542, 'KEPUTRAN', 1044, 51, 4),
		(12543, 'KLEGO', 1044, 51, 4),
		(12544, 'LANDUNGSARI', 1044, 51, 4),
		(12545, 'NOYONTAAN', 1044, 51, 4),
		(12546, 'PONCOL', 1044, 51, 4),
		(12547, 'SAMPANGAN', 1044, 51, 4),
		(12548, 'SOKOREJO', 1044, 51, 4),
		(12549, 'SUGIHWARAS', 1044, 51, 4),
		(12550, 'BANDENGAN', 1045, 51, 4),
		(12551, 'DEGAYU', 1045, 51, 4),
		(12552, 'DUKUH', 1045, 51, 4),
		(12553, 'KANDANGPANJANG', 1045, 51, 4),
		(12554, 'KRAPYAK KIDUL', 1045, 51, 4),
		(12555, 'KRAPYAK LOR', 1045, 51, 4),
		(12556, 'KRATON LOR', 1045, 51, 4),
		(12557, 'PABEAN', 1045, 51, 4),
		(12558, 'PANJANG BARU', 1045, 51, 4),
		(12559, 'PANJANG WETAN', 1045, 51, 4),
		(12560, 'CURUGMUNCAR', 1046, 51, 4),
		(12561, 'GUMELEM', 1046, 51, 4),
		(12562, 'KASIMPAR', 1046, 51, 4),
		(12563, 'KAYUPURING', 1046, 51, 4),
		(12564, 'SIMEGO', 1046, 51, 4),
		(12565, 'SONGGODADI', 1046, 51, 4),
		(12566, 'TLOGOHENDRO', 1046, 51, 4),
		(12567, 'TLOGOPAKIS', 1046, 51, 4),
		(12568, 'YOSOREJO', 1046, 51, 4),
		(12569, 'BLACANAN', 1047, 51, 4),
		(12570, 'BLIMBING WULUH', 1047, 51, 4),
		(12571, 'BOYO TELUK', 1047, 51, 4),
		(12572, 'DEPOK', 1047, 51, 4),
		(12573, 'MEJASEM', 1047, 51, 4),
		(12574, 'PAIT', 1047, 51, 4),
		(12575, 'REMBUN', 1047, 51, 4),
		(12576, 'SIWALAN', 1047, 51, 4),
		(12577, 'TENGENG WETAN', 1047, 51, 4),
		(12578, 'TENGENGKULON', 1047, 51, 4),
		(12579, 'TUNJUNGSARI', 1047, 51, 4),
		(12580, 'WONOSARI', 1047, 51, 4),
		(12581, 'YOSOREJO', 1047, 51, 4),
		(12582, 'BAKTI RASA', 1048, 51, 4),
		(12583, 'BANDAR AGUNG', 1048, 51, 4),
		(12584, 'BULAK PELEM', 1048, 51, 4),
		(12585, 'BULAKSARI', 1048, 51, 4),
		(12586, 'GEBANGKEREP', 1048, 51, 4),
		(12587, 'KALIJAMBE', 1048, 51, 4),
		(12588, 'KEDAUNG', 1048, 51, 4),
		(12589, 'KEDUNGJARAN', 1048, 51, 4),
		(12590, 'KETANON AGENG', 1048, 51, 4),
		(12591, 'KLUNJUKAN', 1048, 51, 4),
		(12592, 'KRASAKAGENG', 1048, 51, 4),
		(12593, 'KUALA SEKAMPUNG', 1048, 51, 4),
		(12594, 'MANDALASARI', 1048, 51, 4),
		(12595, 'MARGAJASA', 1048, 51, 4),
		(12596, 'MARGASARI', 1048, 51, 4),
		(12597, 'MRICAN', 1048, 51, 4),
		(12598, 'PURWODADI', 1048, 51, 4),
		(12599, 'PURWOREJO', 1048, 51, 4),
		(12600, 'SIJERUK', 1048, 51, 4),
		(12601, 'SRAGI', 1048, 51, 4),
		(12602, 'SUKAPURA', 1048, 51, 4),
		(12603, 'SUMBER AGUNG', 1048, 51, 4),
		(12604, 'SUMBER SARI', 1048, 51, 4),
		(12605, 'SUMUB KIDUL', 1048, 51, 4),
		(12606, 'SUMUB LOR', 1048, 51, 4),
		(12607, 'TEGAL SURUH', 1048, 51, 4),
		(12608, 'TEGALONTAR', 1048, 51, 4),
		(12609, 'CURUG', 1049, 51, 4),
		(12610, 'DADIREJO', 1049, 51, 4),
		(12611, 'JERUKSARI', 1049, 51, 4),
		(12612, 'KARANGANYAR', 1049, 51, 4),
		(12613, 'KARANGJOMPO', 1049, 51, 4),
		(12614, 'MULYOREJO', 1049, 51, 4),
		(12615, 'NGALIAN', 1049, 51, 4),
		(12616, 'PACAR', 1049, 51, 4),
		(12617, 'PANDAN ARUM', 1049, 51, 4),
		(12618, 'PUCUNG', 1049, 51, 4),
		(12619, 'SAMBOREJO', 1049, 51, 4),
		(12620, 'SIDOREJO', 1049, 51, 4),
		(12621, 'SILIREJO', 1049, 51, 4),
		(12622, 'TANJUNG', 1049, 51, 4),
		(12623, 'TEGALDOWO', 1049, 51, 4),
		(12624, 'WULED', 1049, 51, 4),
		(12625, 'BENER', 1050, 51, 4),
		(12626, 'BONDANSARI', 1050, 51, 4),
		(12627, 'DELEGTUKANG', 1050, 51, 4),
		(12628, 'GUMAWANG', 1050, 51, 4),
		(12629, 'KADIPATEN', 1050, 51, 4),
		(12630, 'KAMPIL', 1050, 51, 4),
		(12631, 'KARANGJATI', 1050, 51, 4),
		(12632, 'KAUMAN', 1050, 51, 4),
		(12633, 'KEMPLONG', 1050, 51, 4),
		(12634, 'KEPATIHAN', 1050, 51, 4),
		(12635, 'MAYANGAN', 1050, 51, 4),
		(12636, 'PEKUNCEN', 1050, 51, 4),
		(12637, 'PETUKANGAN', 1050, 51, 4),
		(12638, 'WARUKIDUL', 1050, 51, 4),
		(12639, 'WARULOR', 1050, 51, 4),
		(12640, 'WIRADESA', 1050, 51, 4),
		(12641, 'API-API', 1051, 51, 4),
		(12642, 'BEBEL', 1051, 51, 4),
		(12643, 'PECAKARAN', 1051, 51, 4),
		(12644, 'PESANGGRAHAN', 1051, 51, 4),
		(12645, 'ROWOYOSO', 1051, 51, 4),
		(12646, 'SEMUT', 1051, 51, 4),
		(12647, 'SIJAMBE', 1051, 51, 4),
		(12648, 'TRATEBANG', 1051, 51, 4),
		(12649, 'WERDI', 1051, 51, 4),
		(12650, 'WONOKERTO KULON', 1051, 51, 4),
		(12651, 'WONOKERTO WETAN', 1051, 51, 4),
		(12652, 'GALANG PENGAMPON GEDE', 1052, 51, 4),
		(12653, 'GETAS', 1052, 51, 4),
		(12654, 'GONDANG', 1052, 51, 4),
		(12655, 'JETAK KIDUL', 1052, 51, 4),
		(12656, 'JETAK LENGKONG', 1052, 51, 4),
		(12657, 'KWAGEAN', 1052, 51, 4),
		(12658, 'LEGOK GUNUNG', 1052, 51, 4),
		(12659, 'PEGADEN TENGAH', 1052, 51, 4),
		(12660, 'ROWOKEMBU', 1052, 51, 4),
		(12661, 'SAMPIH', 1052, 51, 4),
		(12662, 'SASTRODIRJAN', 1052, 51, 4),
		(12663, 'SURABAYAN', 1052, 51, 4),
		(12664, 'WONOPRINGGO', 1052, 51, 4),
		(12665, 'WONOREJO', 1052, 51, 4),
		(12666, 'BANJARSARI', 1053, 52, 4),
		(12667, 'BANTARBOLANG', 1053, 52, 4),
		(12668, 'GLANDANG', 1053, 52, 4),
		(12669, 'KARANGANYAR', 1053, 52, 4),
		(12670, 'KEBON GEDE', 1053, 52, 4),
		(12671, 'KUTA', 1053, 52, 4),
		(12672, 'LENGGERONG', 1053, 52, 4),
		(12673, 'PABUARAN', 1053, 52, 4),
		(12674, 'PAGUYANGAN', 1053, 52, 4),
		(12675, 'PEDAGUNG', 1053, 52, 4),
		(12676, 'PEGIRINGAN', 1053, 52, 4),
		(12677, 'PURANA', 1053, 52, 4),
		(12678, 'SAMBENG', 1053, 52, 4),
		(12679, 'SARWODADI', 1053, 52, 4),
		(12680, 'SUMURKIDANG', 1053, 52, 4),
		(12681, 'SURU', 1053, 52, 4),
		(12682, 'WANARATA', 1053, 52, 4),
		(12683, 'BADAK', 1054, 52, 4),
		(12684, 'BELIK', 1054, 52, 4),
		(12685, 'BELUK', 1054, 52, 4),
		(12686, 'BULAKAN', 1054, 52, 4),
		(12687, 'GOMBONG', 1054, 52, 4),
		(12688, 'GUNUNGJAYA', 1054, 52, 4),
		(12689, 'GUNUNGTIGA', 1054, 52, 4),
		(12690, 'KALISALEH', 1054, 52, 4),
		(12691, 'KUTA', 1054, 52, 4),
		(12692, 'MENDELEM', 1054, 52, 4),
		(12693, 'SIKASUR', 1054, 52, 4),
		(12694, 'SIMPUR', 1054, 52, 4),
		(12695, 'BABAKAN', 1055, 52, 4),
		(12696, 'BODEH', 1055, 52, 4),
		(12697, 'CANGAK', 1055, 52, 4),
		(12698, 'GUNUNGBATU', 1055, 52, 4),
		(12699, 'JATINGARANG', 1055, 52, 4),
		(12700, 'JATIROYOM', 1055, 52, 4),
		(12701, 'JRAGANAN', 1055, 52, 4),
		(12702, 'KARANGBRAI', 1055, 52, 4),
		(12703, 'KEBANDARAN', 1055, 52, 4),
		(12704, 'KEBANDUNGAN', 1055, 52, 4),
		(12705, 'KELANGDEPOK', 1055, 52, 4),
		(12706, 'KESESIREJO', 1055, 52, 4),
		(12707, 'KWASEN', 1055, 52, 4),
		(12708, 'LONGKEYANG', 1055, 52, 4),
		(12709, 'MUNCANG', 1055, 52, 4),
		(12710, 'PARUNGGALIH', 1055, 52, 4),
		(12711, 'PASIR', 1055, 52, 4),
		(12712, 'PAYUNG', 1055, 52, 4),
		(12713, 'PENDOWO', 1055, 52, 4),
		(12714, 'AMBOKULON', 1056, 52, 4),
		(12715, 'GANDU', 1056, 52, 4),
		(12716, 'GEDEG', 1056, 52, 4),
		(12717, 'GINTUNG', 1056, 52, 4),
		(12718, 'KANDANG', 1056, 52, 4),
		(12719, 'KAUMAN', 1056, 52, 4),
		(12720, 'KEBOJONGAN', 1056, 52, 4),
		(12721, 'KLEGEN', 1056, 52, 4),
		(12722, 'LOWA', 1056, 52, 4),
		(12723, 'PECANGAKAN', 1056, 52, 4),
		(12724, 'PURWOHARJO', 1056, 52, 4),
		(12725, 'PURWOSARI', 1056, 52, 4),
		(12726, 'SARWODADI', 1056, 52, 4),
		(12727, 'SIDOREJO', 1056, 52, 4),
		(12728, 'SIKAYU', 1056, 52, 4),
		(12729, 'SUSUKAN', 1056, 52, 4),
		(12730, 'TUMBAL', 1056, 52, 4),
		(12731, 'WONOKROMO', 1056, 52, 4),
		(12732, 'BANYUMUDAL', 1057, 52, 4),
		(12733, 'GENDOWANG', 1057, 52, 4),
		(12734, 'KEBANGGAN', 1057, 52, 4),
		(12735, 'MANDIRAJA', 1057, 52, 4),
		(12736, 'MOGA', 1057, 52, 4),
		(12737, 'PEPEDAN', 1057, 52, 4),
		(12738, 'PLAKARAN', 1057, 52, 4),
		(12739, 'SIMA', 1057, 52, 4),
		(12740, 'WALANGSANGA', 1057, 52, 4),
		(12741, 'WANGKELANG', 1057, 52, 4),
		(12742, 'BANJARMULYA', 1058, 52, 4),
		(12743, 'BOJONGBATA', 1058, 52, 4),
		(12744, 'BOJONGNANGKA', 1058, 52, 4),
		(12745, 'DANASARI', 1058, 52, 4),
		(12746, 'KEBONDALEM', 1058, 52, 4),
		(12747, 'KRAMAT', 1058, 52, 4),
		(12748, 'LAWANGREJO', 1058, 52, 4),
		(12749, 'MENGORI', 1058, 52, 4),
		(12750, 'MULYOHARJO', 1058, 52, 4),
		(12751, 'PADURAKSA', 1058, 52, 4),
		(12752, 'PAGONGSORAN', 1058, 52, 4),
		(12753, 'PELUTAN', 1058, 52, 4),
		(12754, 'SARADAN', 1058, 52, 4),
		(12755, 'SEWAKA', 1058, 52, 4),
		(12756, 'SUGIHWARAS', 1058, 52, 4),
		(12757, 'SUNGAPAN', 1058, 52, 4),
		(12758, 'SURAJAYA', 1058, 52, 4),
		(12759, 'TAMBAKREJO', 1058, 52, 4),
		(12760, 'WANAMULYA', 1058, 52, 4),
		(12761, 'WIDURI', 1058, 52, 4),
		(12762, 'BULU', 1059, 52, 4),
		(12763, 'ISER', 1059, 52, 4),
		(12764, 'KALIRANDU', 1059, 52, 4),
		(12765, 'KARANGASEM', 1059, 52, 4),
		(12766, 'KENDALDOYONG', 1059, 52, 4),
		(12767, 'KENDALREJO', 1059, 52, 4),
		(12768, 'KENDALSARI', 1059, 52, 4),
		(12769, 'KLAREYAN', 1059, 52, 4),
		(12770, 'LONING', 1059, 52, 4),
		(12771, 'NYAMPLUNG SARI', 1059, 52, 4),
		(12772, 'PANJUNAN', 1059, 52, 4),
		(12773, 'PEGUNDAN', 1059, 52, 4),
		(12774, 'PESUCEN', 1059, 52, 4),
		(12775, 'PETANJUNGAN', 1059, 52, 4),
		(12776, 'PETARUKAN', 1059, 52, 4),
		(12777, 'SERANG', 1059, 52, 4),
		(12778, 'SIRANGKANG', 1059, 52, 4),
		(12779, 'TEGALMLATI', 1059, 52, 4),
		(12780, 'TEMUIRENG', 1059, 52, 4),
		(12781, 'WIDODAREN', 1059, 52, 4),
		(12782, 'BANJARNEGARA', 1060, 52, 4),
		(12783, 'BANJARWANGI', 1060, 52, 4),
		(12784, 'BATURSARI', 1060, 52, 4),
		(12785, 'CIKENDUNG', 1060, 52, 4),
		(12786, 'CILENTUNG', 1060, 52, 4),
		(12787, 'CLEKATAKAN', 1060, 52, 4),
		(12788, 'GAMBUHAN', 1060, 52, 4),
		(12789, 'GUNUNGSARI', 1060, 52, 4),
		(12790, 'JURANGMANGU', 1060, 52, 4),
		(12791, 'KADUHEJO', 1060, 52, 4),
		(12792, 'KARANGSARI', 1060, 52, 4),
		(12793, 'KARYAWANGI', 1060, 52, 4),
		(12794, 'KORANJI', 1060, 52, 4),
		(12795, 'NYALEMBENG', 1060, 52, 4),
		(12796, 'PAGENTERAN', 1060, 52, 4),
		(12797, 'PENAKIR', 1060, 52, 4),
		(12798, 'PULOSARI', 1060, 52, 4),
		(12799, 'SANGHIANGDENGDEK', 1060, 52, 4),
		(12800, 'SIREMENG', 1060, 52, 4),
		(12801, 'SUKARAJA', 1060, 52, 4),
		(12802, 'SUKASARI', 1060, 52, 4),
		(12803, 'BANJARANYAR', 1061, 52, 4),
		(12804, 'GEMBYANG', 1061, 52, 4),
		(12805, 'GONGSENG', 1061, 52, 4),
		(12806, 'KALIMAS', 1061, 52, 4),
		(12807, 'KALITORONG', 1061, 52, 4),
		(12808, 'KARANGMONCOL', 1061, 52, 4),
		(12809, 'KECEPIT', 1061, 52, 4),
		(12810, 'KEJENE', 1061, 52, 4),
		(12811, 'KREYO', 1061, 52, 4),
		(12812, 'LODAYA', 1061, 52, 4),
		(12813, 'MANGLI', 1061, 52, 4),
		(12814, 'MEJAGONG', 1061, 52, 4),
		(12815, 'PENUSUPAN', 1061, 52, 4),
		(12816, 'RANDUDONGKAL', 1061, 52, 4),
		(12817, 'REMBUL', 1061, 52, 4),
		(12818, 'SEMAYA', 1061, 52, 4),
		(12819, 'SEMINGKIR', 1061, 52, 4),
		(12820, 'TANAHBAYA', 1061, 52, 4),
		(12821, 'AMBOWETAN', 1062, 52, 4),
		(12822, 'BLENDUNG', 1062, 52, 4),
		(12823, 'BOTEKAN', 1062, 52, 4),
		(12824, 'BUMIREJO', 1062, 52, 4),
		(12825, 'KALIPRAU', 1062, 52, 4),
		(12826, 'KERTOSARI', 1062, 52, 4),
		(12827, 'KETAPANG', 1062, 52, 4),
		(12828, 'LIMBANGAN', 1062, 52, 4),
		(12829, 'MOJO', 1062, 52, 4),
		(12830, 'PADEK', 1062, 52, 4),
		(12831, 'PAGERGUNUNG', 1062, 52, 4),
		(12832, 'PAMUTIH', 1062, 52, 4),
		(12833, 'PESANTREN', 1062, 52, 4),
		(12834, 'ROWOSARI', 1062, 52, 4),
		(12835, 'SAMONG', 1062, 52, 4),
		(12836, 'SUKOREJO', 1062, 52, 4),
		(12837, 'TASIKREJO', 1062, 52, 4),
		(12838, 'WIYOROWETAN', 1062, 52, 4),
		(12839, 'CIBUYUR', 1063, 52, 4),
		(12840, 'DATAR', 1063, 52, 4),
		(12841, 'KARANGDAWA', 1063, 52, 4),
		(12842, 'MERENG', 1063, 52, 4),
		(12843, 'PAKEMBARAN', 1063, 52, 4),
		(12844, 'WARUNGPRING', 1063, 52, 4),
		(12845, 'BODAS', 1064, 52, 4),
		(12846, 'BONGAS', 1064, 52, 4),
		(12847, 'CAWET', 1064, 52, 4),
		(12848, 'CIKADU', 1064, 52, 4),
		(12849, 'GAPURA', 1064, 52, 4),
		(12850, 'JOJOGAN', 1064, 52, 4),
		(12851, 'MAJAKERTA', 1064, 52, 4),
		(12852, 'MAJALANGU', 1064, 52, 4),
		(12853, 'MEDAYU', 1064, 52, 4),
		(12854, 'PAGELARAN', 1064, 52, 4),
		(12855, 'TAMBI', 1064, 52, 4),
		(12856, 'TLAGASANA', 1064, 52, 4),
		(12857, 'TUNDAGAN', 1064, 52, 4),
		(12858, 'WATUKUMPUL', 1064, 52, 4),
		(12859, 'WISNU', 1064, 52, 4),
		(12860, 'BANJARSARI', 1065, 53, 4),
		(12861, 'BOBOTSARI', 1065, 53, 4),
		(12862, 'DAGAN', 1065, 53, 4),
		(12863, 'GANDASULI', 1065, 53, 4),
		(12864, 'GUNUNGKARANG', 1065, 53, 4),
		(12865, 'KALAPACUNG', 1065, 53, 4),
		(12866, 'KARANGDUREN', 1065, 53, 4),
		(12867, 'KARANGMALANG', 1065, 53, 4),
		(12868, 'KARANGTALUN', 1065, 53, 4),
		(12869, 'LIMBASARI', 1065, 53, 4),
		(12870, 'MAJAPURA', 1065, 53, 4),
		(12871, 'PAKUNCEN', 1065, 53, 4),
		(12872, 'PALUMBUNGAN KULON', 1065, 53, 4),
		(12873, 'PALUMBUNGAN WETAN', 1065, 53, 4),
		(12874, 'TALAGENING', 1065, 53, 4),
		(12875, 'TLAGAYASA', 1065, 53, 4),
		(12876, 'BANJARAN', 1066, 53, 4),
		(12877, 'BEJI', 1066, 53, 4),
		(12878, 'BOJONG SARI (LAMA)', 1066, 53, 4),
		(12879, 'BOJONG SARI BARU', 1066, 53, 4),
		(12880, 'BOJONGSARI', 1066, 53, 4),
		(12881, 'BROBOT', 1066, 53, 4),
		(12882, 'BUMISARI', 1066, 53, 4),
		(12883, 'CURUG', 1066, 53, 4),
		(12884, 'DUREN MEKAR', 1066, 53, 4),
		(12885, 'DUREN SERIBU', 1066, 53, 4),
		(12886, 'GALUH', 1066, 53, 4),
		(12887, 'GEMBONG', 1066, 53, 4),
		(12888, 'KAJONGAN', 1066, 53, 4),
		(12889, 'KARANGBANJAR', 1066, 53, 4),
		(12890, 'METENGGENG', 1066, 53, 4),
		(12891, 'PAGEDANGAN', 1066, 53, 4),
		(12892, 'PATEMON', 1066, 53, 4),
		(12893, 'PEKALONGAN', 1066, 53, 4),
		(12894, 'PONDOK PETIR', 1066, 53, 4),
		(12895, 'SERUA', 1066, 53, 4),
		(12896, 'BAJONG', 1067, 53, 4),
		(12897, 'BUKATEJA', 1067, 53, 4),
		(12898, 'CIPAWON', 1067, 53, 4),
		(12899, 'KARANGCENGIS', 1067, 53, 4),
		(12900, 'KARANGGEDANG', 1067, 53, 4),
		(12901, 'KARANGNANGKA', 1067, 53, 4),
		(12902, 'KEBUTUH', 1067, 53, 4),
		(12903, 'KEDUNGJATI', 1067, 53, 4),
		(12904, 'KEMBANGAN', 1067, 53, 4),
		(12905, 'KUTAWIS', 1067, 53, 4),
		(12906, 'MAJASARI', 1067, 53, 4),
		(12907, 'PENARUBAN', 1067, 53, 4),
		(12908, 'TIDU', 1067, 53, 4),
		(12909, 'WIRASABA', 1067, 53, 4),
		(12910, 'ARENAN', 1068, 53, 4),
		(12911, 'BRECEK', 1068, 53, 4),
		(12912, 'CILAPAR', 1068, 53, 4),
		(12913, 'KALIGONDANG', 1068, 53, 4),
		(12914, 'KALIKAJAR', 1068, 53, 4),
		(12915, 'KEMBARAN WETAN', 1068, 53, 4),
		(12916, 'LAMONGAN', 1068, 53, 4),
		(12917, 'PAGERANDONG', 1068, 53, 4),
		(12918, 'PENARUBAN', 1068, 53, 4),
		(12919, 'PENOLIH', 1068, 53, 4),
		(12920, 'SELAKAMBANG', 1068, 53, 4),
		(12921, 'SELANEGARA', 1068, 53, 4),
		(12922, 'SEMPOR LOR', 1068, 53, 4),
		(12923, 'SIDANEGARA', 1068, 53, 4),
		(12924, 'SIDAREJA', 1068, 53, 4),
		(12925, 'SINDURAJA', 1068, 53, 4),
		(12926, 'SLINGA', 1068, 53, 4),
		(12927, 'TEJASARI', 1068, 53, 4),
		(12928, 'BABAKAN', 1069, 53, 4),
		(12929, 'BLATER', 1069, 53, 4),
		(12930, 'GRECOL', 1069, 53, 4),
		(12931, 'JOMPO', 1069, 53, 4),
		(12932, 'KALIKABONG', 1069, 53, 4),
		(12933, 'KALIMANAH KULON', 1069, 53, 4),
		(12934, 'KALIMANAH WETAN', 1069, 53, 4),
		(12935, 'KARANGMANYAR', 1069, 53, 4),
		(12936, 'KARANGPETIR', 1069, 53, 4),
		(12937, 'KARANGSARI', 1069, 53, 4),
		(12938, 'KEDUNGWULUH', 1069, 53, 4),
		(12939, 'KLAPASAWIT', 1069, 53, 4),
		(12940, 'MANDURAGA', 1069, 53, 4),
		(12941, 'MEWEK', 1069, 53, 4),
		(12942, 'RABAK', 1069, 53, 4),
		(12943, 'SELABAYA', 1069, 53, 4),
		(12944, 'SIDAKANGEN', 1069, 53, 4),
		(12945, 'DANASARI', 1070, 53, 4),
		(12946, 'JINGKANG', 1070, 53, 4),
		(12947, 'KARANGJAMBU', 1070, 53, 4),
		(12948, 'PURBASARI', 1070, 53, 4),
		(12949, 'SANGUWATANG', 1070, 53, 4),
		(12950, 'SIRANDU', 1070, 53, 4),
		(12951, 'BALERAKSA', 1071, 53, 4),
		(12952, 'GRANTUNG', 1071, 53, 4),
		(12953, 'KARANGSARI', 1071, 53, 4),
		(12954, 'KRAMAT', 1071, 53, 4),
		(12955, 'PEKIRINGAN', 1071, 53, 4),
		(12956, 'PEPEDAN', 1071, 53, 4),
		(12957, 'RAJAWANA', 1071, 53, 4),
		(12958, 'SIRAU', 1071, 53, 4),
		(12959, 'TAJUG', 1071, 53, 4),
		(12960, 'TAMANSARI', 1071, 53, 4),
		(12961, 'TUNJUNGMULI', 1071, 53, 4),
		(12962, 'GONDANG', 1072, 53, 4),
		(12963, 'KARANGREJA', 1072, 53, 4),
		(12964, 'KUTABAWA', 1072, 53, 4),
		(12965, 'SERANG', 1072, 53, 4),
		(12966, 'SIWARAK', 1072, 53, 4),
		(12967, 'TLAHAB KIDUL', 1072, 53, 4),
		(12968, 'TLAHAB LOR', 1072, 53, 4),
		(12969, 'BANDINGAN', 1073, 53, 4),
		(12970, 'GUMIWANG', 1073, 53, 4),
		(12971, 'KEDARPAN', 1073, 53, 4),
		(12972, 'KEJOBONG', 1073, 53, 4),
		(12973, 'KRENCENG', 1073, 53, 4),
		(12974, 'LAMUK', 1073, 53, 4),
		(12975, 'LANGGAR', 1073, 53, 4),
		(12976, 'NANGKASAWIT', 1073, 53, 4),
		(12977, 'NANGKOD', 1073, 53, 4),
		(12978, 'PANDANSARI', 1073, 53, 4),
		(12979, 'PANGEMPON', 1073, 53, 4),
		(12980, 'SOKANEGARA', 1073, 53, 4),
		(12981, 'TIMBANG', 1073, 53, 4),
		(12982, 'BAKULAN', 1074, 53, 4),
		(12983, 'BOKOL', 1074, 53, 4),
		(12984, 'GAMBARSARI', 1074, 53, 4),
		(12985, 'JETIS', 1074, 53, 4),
		(12986, 'KALIALANG', 1074, 53, 4),
		(12987, 'KARANGKEMIRI', 1074, 53, 4),
		(12988, 'KARANGTENGAH', 1074, 53, 4),
		(12989, 'KEDUNGBENDA', 1074, 53, 4),
		(12990, 'KEDUNGLEGOK', 1074, 53, 4),
		(12991, 'KEMANGKON', 1074, 53, 4),
		(12992, 'MAJASEM', 1074, 53, 4),
		(12993, 'MAJATENGAH', 1074, 53, 4),
		(12994, 'MUNTANG', 1074, 53, 4),
		(12995, 'PANICAN', 1074, 53, 4),
		(12996, 'PEGANDEKAN', 1074, 53, 4),
		(12997, 'PLUMUTAN', 1074, 53, 4),
		(12998, 'SENON', 1074, 53, 4),
		(12999, 'SUMILIR', 1074, 53, 4),
		(13000, 'TOYAREKA', 1074, 53, 4),
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 12")
}
