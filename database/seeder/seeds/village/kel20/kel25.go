package kel20

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kel25() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	(25001, 'KARANG', 1905, 99, 6),
	(25002, 'KOWANG', 1905, 99, 6),
	(25003, 'NGINO', 1905, 99, 6),
	(25004, 'PENAMBANGAN', 1905, 99, 6),
	(25005, 'PRUNGGAHAN KULON', 1905, 99, 6),
	(25006, 'PRUNGGAHAN WETAN', 1905, 99, 6),
	(25007, 'SAMBONGREJO', 1905, 99, 6),
	(25008, 'SEMANDING', 1905, 99, 6),
	(25009, 'TEGALAGUNG', 1905, 99, 6),
	(25010, 'TUNAH', 1905, 99, 6),
	(25011, 'BANYUURIP', 1906, 99, 6),
	(25012, 'JATISARI', 1906, 99, 6),
	(25013, 'KALIGEDE', 1906, 99, 6),
	(25014, 'KATERBAN', 1906, 99, 6),
	(25015, 'LERAN', 1906, 99, 6),
	(25016, 'MEDALEM', 1906, 99, 6),
	(25017, 'RAYUNG', 1906, 99, 6),
	(25018, 'SENDANG', 1906, 99, 6),
	(25019, 'SIDOHARJO', 1906, 99, 6),
	(25020, 'WANGLU KULON', 1906, 99, 6),
	(25021, 'WANGLU WETAN', 1906, 99, 6),
	(25022, 'WONOSARI', 1906, 99, 6),
	(25023, 'BINANGUN', 1907, 99, 6),
	(25024, 'KEDUNGJAMBE', 1907, 99, 6),
	(25025, 'LAJO KIDUL', 1907, 99, 6),
	(25026, 'LAJO LOR', 1907, 99, 6),
	(25027, 'MERGOSARI', 1907, 99, 6),
	(25028, 'MULYOAGUNG', 1907, 99, 6),
	(25029, 'MULYOREJO', 1907, 99, 6),
	(25030, 'SARINGEMBAT', 1907, 99, 6),
	(25031, 'TANGGIR', 1907, 99, 6),
	(25032, 'TANJUNGREJO', 1907, 99, 6),
	(25033, 'TINGKIS', 1907, 99, 6),
	(25034, 'TUNGGULREJO', 1907, 99, 6),
	(25035, 'BANGUNREJO', 1908, 99, 6),
	(25036, 'CEKALANG', 1908, 99, 6),
	(25037, 'GLAGAHSARI', 1908, 99, 6),
	(25038, 'GUNUNGANYAR', 1908, 99, 6),
	(25039, 'JATI', 1908, 99, 6),
	(25040, 'JEGULO', 1908, 99, 6),
	(25041, 'KENDALREJO', 1908, 99, 6),
	(25042, 'KENONGOSARI', 1908, 99, 6),
	(25043, 'KLUMPIT', 1908, 99, 6),
	(25044, 'MENILO', 1908, 99, 6),
	(25045, 'MENTORO', 1908, 99, 6),
	(25046, 'MOJOAGUNG', 1908, 99, 6),
	(25047, 'NGURUAN', 1908, 99, 6),
	(25048, 'PANDANAGUNG', 1908, 99, 6),
	(25049, 'PANDANWANGI', 1908, 99, 6),
	(25050, 'PRAMBONTERGAYANG', 1908, 99, 6),
	(25051, 'RAHAYU', 1908, 99, 6),
	(25052, 'SANDINGROWO', 1908, 99, 6),
	(25053, 'SIMO', 1908, 99, 6),
	(25054, 'SOKOSARI', 1908, 99, 6),
	(25055, 'SUMURCINDE', 1908, 99, 6),
	(25056, 'TLUWE', 1908, 99, 6),
	(25057, 'WADUNG', 1908, 99, 6),
	(25058, 'BELIKANGET', 1909, 99, 6),
	(25059, 'COKROWATI', 1909, 99, 6),
	(25060, 'DASIN', 1909, 99, 6),
	(25061, 'DIKIR', 1909, 99, 6),
	(25062, 'GADON', 1909, 99, 6),
	(25063, 'GLONDONGGEDE', 1909, 99, 6),
	(25064, 'KENANTI', 1909, 99, 6),
	(25065, 'KLUTUK', 1909, 99, 6),
	(25066, 'MANDER', 1909, 99, 6),
	(25067, 'MERKAWANG', 1909, 99, 6),
	(25068, 'NGULAHAN', 1909, 99, 6),
	(25069, 'PABEYAN', 1909, 99, 6),
	(25070, 'PLAJAN', 1909, 99, 6),
	(25071, 'PULOGEDE', 1909, 99, 6),
	(25072, 'SAWIR', 1909, 99, 6),
	(25073, 'SOBONTORO', 1909, 99, 6),
	(25074, 'SOTANG', 1909, 99, 6),
	(25075, 'TAMBAKBOYO', 1909, 99, 6),
	(25076, 'BATURETNO', 1910, 99, 6),
	(25077, 'DOROMUKTI', 1910, 99, 6),
	(25078, 'KARANGSARI', 1910, 99, 6),
	(25079, 'KEBONSARI', 1910, 99, 6),
	(25080, 'KEMBANGBILO', 1910, 99, 6),
	(25081, 'KINGKING', 1910, 99, 6),
	(25082, 'KUTOREJO', 1910, 99, 6),
	(25083, 'LATSARI', 1910, 99, 6),
	(25084, 'MONDOKAN', 1910, 99, 6),
	(25085, 'PERBON', 1910, 99, 6),
	(25086, 'RONGGOMULYO', 1910, 99, 6),
	(25087, 'SENDANGHARJO', 1910, 99, 6),
	(25088, 'SIDOMULYO', 1910, 99, 6),
	(25089, 'SIDOREJO', 1910, 99, 6),
	(25090, 'SUGIHARJO', 1910, 99, 6),
	(25091, 'SUKOLILO', 1910, 99, 6),
	(25092, 'SUMURGUNG', 1910, 99, 6),
	(25093, 'BANJAR', 1911, 99, 6),
	(25094, 'BUNUT', 1911, 99, 6),
	(25095, 'COMPRENG', 1911, 99, 6),
	(25096, 'KEDUNGHARJO', 1911, 99, 6),
	(25097, 'KUJUNG', 1911, 99, 6),
	(25098, 'MINOHOREJO', 1911, 99, 6),
	(25099, 'MLANGI', 1911, 99, 6),
	(25100, 'MRUTUK', 1911, 99, 6),
	(25101, 'NGADIPURO', 1911, 99, 6),
	(25102, 'NGADIREJO', 1911, 99, 6),
	(25103, 'PATIHAN', 1911, 99, 6),
	(25104, 'SIMOREJO', 1911, 99, 6),
	(25105, 'SUMBEREJO', 1911, 99, 6),
	(25106, 'TEGALREJO', 1911, 99, 6),
	(25107, 'TEGALSARI', 1911, 99, 6),
	(25108, 'WIDANG', 1911, 99, 6),
	(25109, 'BABAKAN', 1912, 100, 6),
	(25110, 'BANDUNG', 1912, 100, 6),
	(25111, 'BANDUNG', 1912, 100, 6),
	(25112, 'BANTENGAN', 1912, 100, 6),
	(25113, 'BLOKANG', 1912, 100, 6),
	(25114, 'BULUS', 1912, 100, 6),
	(25115, 'GANDONG', 1912, 100, 6),
	(25116, 'KEDUNGWILUT', 1912, 100, 6),
	(25117, 'KESAMBI', 1912, 100, 6),
	(25118, 'MALABAR', 1912, 100, 6),
	(25119, 'MANDER', 1912, 100, 6),
	(25120, 'MERGAYU', 1912, 100, 6),
	(25121, 'NGEPEH', 1912, 100, 6),
	(25122, 'NGLAMPIR', 1912, 100, 6),
	(25123, 'NGUNGGAHAN', 1912, 100, 6),
	(25124, 'PANAMPING', 1912, 100, 6),
	(25125, 'PANGAWINAN', 1912, 100, 6),
	(25126, 'PRINGWULUNG', 1912, 100, 6),
	(25127, 'SEBALOR', 1912, 100, 6),
	(25128, 'SINGGIT', 1912, 100, 6),
	(25129, 'SOKO', 1912, 100, 6),
	(25130, 'SUKOHARJO', 1912, 100, 6),
	(25131, 'SURUHAN KIDUL', 1912, 100, 6),
	(25132, 'SURUHAN LOR', 1912, 100, 6),
	(25133, 'SUWARU', 1912, 100, 6),
	(25134, 'TALUNKULON', 1912, 100, 6),
	(25135, 'BEJI', 1913, 100, 6),
	(25136, 'BONO', 1913, 100, 6),
	(25137, 'BOYOLANGU', 1913, 100, 6),
	(25138, 'GEDANGSEWU', 1913, 100, 6),
	(25139, 'KARANGREJO', 1913, 100, 6),
	(25140, 'KENDALBULUR', 1913, 100, 6),
	(25141, 'KEPUH', 1913, 100, 6),
	(25142, 'MOYOKETEN', 1913, 100, 6),
	(25143, 'NGRANTI', 1913, 100, 6),
	(25144, 'PUCUNG KIDUL', 1913, 100, 6),
	(25145, 'SANGGRAHAN', 1913, 100, 6),
	(25146, 'SERUT', 1913, 100, 6),
	(25147, 'SOBONTORO', 1913, 100, 6),
	(25148, 'TANJUNGSARI', 1913, 100, 6),
	(25149, 'WAJAK KIDUL', 1913, 100, 6),
	(25150, 'WAJAK LOR', 1913, 100, 6),
	(25151, 'WAUNG', 1913, 100, 6),
	(25152, 'CAMPURDARAT', 1914, 100, 6),
	(25153, 'GAMPING', 1914, 100, 6),
	(25154, 'GEDANGAN', 1914, 100, 6),
	(25155, 'NGENTRONG', 1914, 100, 6),
	(25156, 'PELEM', 1914, 100, 6),
	(25157, 'POJOK', 1914, 100, 6),
	(25158, 'SAWO', 1914, 100, 6),
	(25159, 'TANGGUNG', 1914, 100, 6),
	(25160, 'WATES', 1914, 100, 6),
	(25161, 'BAKALAN', 1915, 100, 6),
	(25162, 'BALONG GEBANG', 1915, 100, 6),
	(25163, 'BEGAGANLIMO', 1915, 100, 6),
	(25164, 'BENDO', 1915, 100, 6),
	(25165, 'BENDUNGAN', 1915, 100, 6),
	(25166, 'BENING', 1915, 100, 6),
	(25167, 'BLENDIS', 1915, 100, 6),
	(25168, 'BUMIAJI', 1915, 100, 6),
	(25169, 'CAMPUR', 1915, 100, 6),
	(25170, 'CENTONG', 1915, 100, 6),
	(25171, 'DILEM', 1915, 100, 6),
	(25172, 'DUKUH', 1915, 100, 6),
	(25173, 'GLONGGONG', 1915, 100, 6),
	(25174, 'GONDANG', 1915, 100, 6),
	(25175, 'GONDANG', 1915, 100, 6),
	(25176, 'GONDANG', 1915, 100, 6),
	(25177, 'GONDANG', 1915, 100, 6),
	(25178, 'GONDANG KULON', 1915, 100, 6),
	(25179, 'GONDOSULI', 1915, 100, 6),
	(25180, 'GUMENG', 1915, 100, 6),
	(25181, 'JAAN', 1915, 100, 6),
	(25182, 'JARAKAN', 1915, 100, 6),
	(25183, 'JARI', 1915, 100, 6),
	(25184, 'JATIDUKUH', 1915, 100, 6),
	(25185, 'KALIKATIR', 1915, 100, 6),
	(25186, 'KALIWEDI', 1915, 100, 6),
	(25187, 'KARANGKUTEN', 1915, 100, 6),
	(25188, 'KARANGSEMI', 1915, 100, 6),
	(25189, 'KEBONTUNGGUL', 1915, 100, 6),
	(25190, 'KEDUNG GLUGU', 1915, 100, 6),
	(25191, 'KEMASANTANI', 1915, 100, 6),
	(25192, 'KENDAL', 1915, 100, 6),
	(25193, 'KETAWANG', 1915, 100, 6),
	(25194, 'KIPING', 1915, 100, 6),
	(25195, 'KRONDONAN', 1915, 100, 6),
	(25196, 'LOSARI', 1915, 100, 6),
	(25197, 'MACANBANG', 1915, 100, 6),
	(25198, 'MOJOARUM', 1915, 100, 6),
	(25199, 'MOJOSETO', 1915, 100, 6),
	(25200, 'NGEMBAT', 1915, 100, 6),
	(25201, 'NGLINGGO', 1915, 100, 6),
	(25202, 'NGRENDENG', 1915, 100, 6),
	(25203, 'NGUJUNG', 1915, 100, 6),
	(25204, 'NOTOREJO', 1915, 100, 6),
	(25205, 'PADI', 1915, 100, 6),
	(25206, 'PAJENG', 1915, 100, 6),
	(25207, 'PANDEAN', 1915, 100, 6),
	(25208, 'PLOSOREJO', 1915, 100, 6),
	(25209, 'POHJEJER', 1915, 100, 6),
	(25210, 'PRAGELAN', 1915, 100, 6),
	(25211, 'PUGERAN', 1915, 100, 6),
	(25212, 'REJOSARI', 1915, 100, 6),
	(25213, 'SAMBONGREJO', 1915, 100, 6),
	(25214, 'SANGGRAHAN', 1915, 100, 6),
	(25215, 'SENGATEN', 1915, 100, 6),
	(25216, 'SENGGOWAR', 1915, 100, 6),
	(25217, 'SENJAYAN', 1915, 100, 6),
	(25218, 'SEPATAN', 1915, 100, 6),
	(25219, 'SIDEM', 1915, 100, 6),
	(25220, 'SIDOMULYO', 1915, 100, 6),
	(25221, 'SRIMULYO', 1915, 100, 6),
	(25222, 'SUMBER AGUNG', 1915, 100, 6),
	(25223, 'SUMBERJO', 1915, 100, 6),
	(25224, 'TAWAR', 1915, 100, 6),
	(25225, 'TAWING', 1915, 100, 6),
	(25226, 'TEGALREJO', 1915, 100, 6),
	(25227, 'TIUDAN', 1915, 100, 6),
	(25228, 'TUNGGUL', 1915, 100, 6),
	(25229, 'WONOKROMO', 1915, 100, 6),
	(25230, 'WONOPLOSO', 1915, 100, 6),
	(25231, 'WONOTOLO', 1915, 100, 6),
	(25232, 'BANYUURIP', 1916, 100, 6),
	(25233, 'BETAK', 1916, 100, 6),
	(25234, 'DOMASAN', 1916, 100, 6),
	(25235, 'JABON', 1916, 100, 6),
	(25236, 'JOHO', 1916, 100, 6),
	(25237, 'KALIBATUR', 1916, 100, 6),
	(25238, 'KALIDAWIR', 1916, 100, 6),
	(25239, 'KARANGTALUN', 1916, 100, 6),
	(25240, 'NGUBALAN', 1916, 100, 6),
	(25241, 'PAGERSARI', 1916, 100, 6),
	(25242, 'PAKISAJI', 1916, 100, 6),
	(25243, 'REJOSARI', 1916, 100, 6),
	(25244, 'SALAK KEMBANG', 1916, 100, 6),
	(25245, 'SUKOREJO KULON', 1916, 100, 6),
	(25246, 'TANJUNG', 1916, 100, 6),
	(25247, 'TUNGGANGRI', 1916, 100, 6),
	(25248, 'WINONG', 1916, 100, 6),
	(25249, 'BABADAN', 1917, 100, 6),
	(25250, 'BUNGUR', 1917, 100, 6),
	(25251, 'GEDANGAN', 1917, 100, 6),
	(25252, 'JELI', 1917, 100, 6),
	(25253, 'KARANGREJO', 1917, 100, 6),
	(25254, 'PUNJUL', 1917, 100, 6),
	(25255, 'SEMBON', 1917, 100, 6),
	(25256, 'SUKODONO', 1917, 100, 6),
	(25257, 'SUKOREJO', 1917, 100, 6),
	(25258, 'SUKOWIDODO', 1917, 100, 6),
	(25259, 'SUKOWIYONO', 1917, 100, 6),
	(25260, 'TANJUNGSARI', 1917, 100, 6),
	(25261, 'TULUNGREJO', 1917, 100, 6),
	(25262, 'BALEREJO', 1918, 100, 6),
	(25263, 'BANARAN', 1918, 100, 6),
	(25264, 'BATANGSAREN', 1918, 100, 6),
	(25265, 'BOLOREJO', 1918, 100, 6),
	(25266, 'BRINGIN', 1918, 100, 6),
	(25267, 'CARAT', 1918, 100, 6),
	(25268, 'CILUK', 1918, 100, 6),
	(25269, 'GABEL', 1918, 100, 6),
	(25270, 'JATIMULYO', 1918, 100, 6),
	(25271, 'KALANGBRET', 1918, 100, 6),
	(25272, 'KARANGANOM', 1918, 100, 6),
	(25273, 'KATES', 1918, 100, 6),
	(25274, 'KAUMAN', 1918, 100, 6),
	(25275, 'KAUMAN', 1918, 100, 6),
	(25276, 'MARON', 1918, 100, 6),
	(25277, 'MOJOSARI', 1918, 100, 6),
	(25278, 'NGLARANGAN (NLARANGAN)', 1918, 100, 6),
	(25279, 'NGRANDU', 1918, 100, 6),
	(25280, 'NONGKODONO', 1918, 100, 6),
	(25281, 'PANGGUNGREJO', 1918, 100, 6),
	(25282, 'PENGKOL', 1918, 100, 6),
	(25283, 'PLOSOJENAR', 1918, 100, 6),
	(25284, 'PUCANGAN', 1918, 100, 6),
	(25285, 'SEMANDING', 1918, 100, 6),
	(25286, 'SIDOREJO', 1918, 100, 6),
	(25287, 'SOMOROTO', 1918, 100, 6),
	(25288, 'SUKOSARI', 1918, 100, 6),
	(25289, 'TEGALOMBO', 1918, 100, 6),
	(25290, 'TOSANAN', 1918, 100, 6),
	(25291, 'BANGOAN', 1919, 100, 6),
	(25292, 'BORO', 1919, 100, 6),
	(25293, 'BULUSARI', 1919, 100, 6),
	(25294, 'GENDINGAN', 1919, 100, 6),
	(25295, 'KEDUNGWARU', 1919, 100, 6),
	(25296, 'KETANON', 1919, 100, 6),
	(25297, 'LODERESAN', 1919, 100, 6),
	(25298, 'MAJAN', 1919, 100, 6),
	(25299, 'MANGUNSARI', 1919, 100, 6),
	(25300, 'NGUJANG', 1919, 100, 6),
	(25301, 'PLANDAAN', 1919, 100, 6),
	(25302, 'PLOSOKANDANG', 1919, 100, 6),
	(25303, 'REJOAGUNG', 1919, 100, 6),
	(25304, 'RINGINPITU', 1919, 100, 6),
	(25305, 'SIMO', 1919, 100, 6),
	(25306, 'TAPAN', 1919, 100, 6),
	(25307, 'TAWANGSARI', 1919, 100, 6),
	(25308, 'TUNGGULSARI', 1919, 100, 6),
	(25309, 'WINONG', 1919, 100, 6),
	(25310, 'BANJARSARI', 1920, 100, 6),
	(25311, 'BATOKAN', 1920, 100, 6),
	(25312, 'BENDOSARI', 1920, 100, 6),
	(25313, 'KEPUHREJO', 1920, 100, 6),
	(25314, 'MOJOAGUNG', 1920, 100, 6),
	(25315, 'NGANTRU', 1920, 100, 6),
	(25316, 'PADANGAN', 1920, 100, 6),
	(25317, 'PAKEL', 1920, 100, 6),
	(25318, 'PINGGIRSARI', 1920, 100, 6),
	(25319, 'POJOK', 1920, 100, 6),
	(25320, 'PUCUNG LOR', 1920, 100, 6),
	(25321, 'PULEREJO', 1920, 100, 6),
	(25322, 'SRIKATON', 1920, 100, 6),
	(25323, 'BALESONO', 1921, 100, 6),
	(25324, 'GILANG', 1921, 100, 6),
	(25325, 'KACANGAN', 1921, 100, 6),
	(25326, 'KALANGAN', 1921, 100, 6),
	(25327, 'KALIWUNGU', 1921, 100, 6),
	(25328, 'KARANGSONO', 1921, 100, 6),
	(25329, 'KROMASAN', 1921, 100, 6),
	(25330, 'NGUNUT', 1921, 100, 6),
	(25331, 'PANDANSARI', 1921, 100, 6),
	(25332, 'PULOSARI', 1921, 100, 6),
	(25333, 'PULOTONDO', 1921, 100, 6),
	(25334, 'PURWOREJO', 1921, 100, 6),
	(25335, 'SAMIR', 1921, 100, 6),
	(25336, 'SELOREJO', 1921, 100, 6),
	(25337, 'SUMBEREJO KULON', 1921, 100, 6),
	(25338, 'SUMBEREJO WETAN', 1921, 100, 6),
	(25339, 'SUMBERINGIN KIDUL', 1921, 100, 6),
	(25340, 'SUMBERINGIN KULON', 1921, 100, 6),
	(25341, 'GAMBIRAN', 1922, 100, 6),
	(25342, 'GONDANGGUNUNG', 1922, 100, 6),
	(25343, 'KEDUNGCANGKRING', 1922, 100, 6),
	(25344, 'KRADINAN', 1922, 100, 6),
	(25345, 'MULYOSARI', 1922, 100, 6),
	(25346, 'PAGERWOJO', 1922, 100, 6),
	(25347, 'PENJOR', 1922, 100, 6),
	(25348, 'SAMAR', 1922, 100, 6),
	(25349, 'SEGAWE', 1922, 100, 6),
	(25350, 'SIDOMULYO', 1922, 100, 6),
	(25351, 'WONOREJO', 1922, 100, 6),
	(25352, 'BANGUNJAYA', 1923, 100, 6),
	(25353, 'BANGUNMULYO', 1923, 100, 6),
	(25354, 'BONO', 1923, 100, 6),
	(25355, 'DUWET', 1923, 100, 6),
	(25356, 'GEBANG', 1923, 100, 6),
	(25357, 'GEMPOLAN', 1923, 100, 6),
	(25358, 'GESIKAN', 1923, 100, 6),
	(25359, 'GOMBANG', 1923, 100, 6),
	(25360, 'KASREMAN', 1923, 100, 6),
	(25361, 'NGEBONG', 1923, 100, 6),
	(25362, 'NGRANCE', 1923, 100, 6),
	(25363, 'PAKEL', 1923, 100, 6),
	(25364, 'PECUK', 1923, 100, 6),
	(25365, 'SAMBITAN', 1923, 100, 6),
	(25366, 'SANAN', 1923, 100, 6),
	(25367, 'SODO', 1923, 100, 6),
	(25368, 'SUKOANYAR', 1923, 100, 6),
	(25369, 'SUWALUH', 1923, 100, 6),
	(25370, 'TAMBAN', 1923, 100, 6),
	(25371, 'DEMUK', 1924, 100, 6),
	(25372, 'KALIDAWE', 1924, 100, 6),
	(25373, 'KALIGENTONG', 1924, 100, 6),
	(25374, 'MANDING', 1924, 100, 6),
	(25375, 'PANGGUNGKALAK', 1924, 100, 6),
	(25376, 'PANGGUNGUNI', 1924, 100, 6),
	(25377, 'PUCANGLABAN', 1924, 100, 6),
	(25378, 'SUMBERBENDO', 1924, 100, 6),
	(25379, 'SUMBERDADAP', 1924, 100, 6),
	(25380, 'ARYOJEDING', 1925, 100, 6),
	(25381, 'BANJAREJO', 1925, 100, 6),
	(25382, 'BLIMBING', 1925, 100, 6),
	(25383, 'BUNTARAN', 1925, 100, 6),
	(25384, 'JATIDOWO', 1925, 100, 6),
	(25385, 'KARANGSARI', 1925, 100, 6),
	(25386, 'PAKISREJO', 1925, 100, 6),
	(25387, 'PANJEREJO', 1925, 100, 6),
	(25388, 'REJOTANGAN', 1925, 100, 6),
	(25389, 'SUKOREJO WETAN', 1925, 100, 6),
	(25390, 'SUMBERAGUNG', 1925, 100, 6),
	(25391, 'TANEN', 1925, 100, 6),
	(25392, 'TEGALREJO', 1925, 100, 6),
	(25393, 'TENGGONG', 1925, 100, 6),
	(25394, 'TENGGUR', 1925, 100, 6),
	(25395, 'TUGU', 1925, 100, 6),
	(25396, 'DONO', 1926, 100, 6),
	(25397, 'GEGER', 1926, 100, 6),
	(25398, 'KEDOYO', 1926, 100, 6),
	(25399, 'KROSOK', 1926, 100, 6),
	(25400, 'NGLURUP', 1926, 100, 6),
	(25401, 'NGLUTUNG', 1926, 100, 6),
	(25402, 'NYAWANGAN', 1926, 100, 6),
	(25403, 'PICISAN', 1926, 100, 6),
	(25404, 'SENDANG', 1926, 100, 6),
	(25405, 'TALANG', 1926, 100, 6),
	(25406, 'TUGU', 1926, 100, 6),
	(25407, 'BENDILJATI KULON', 1927, 100, 6),
	(25408, 'BENDILJATI WETAN', 1927, 100, 6),
	(25409, 'BENDILWUNGU', 1927, 100, 6),
	(25410, 'BUKUR', 1927, 100, 6),
	(25411, 'DOROAMPEL', 1927, 100, 6),
	(25412, 'JABALSARI', 1927, 100, 6),
	(25413, 'JUNJUNG', 1927, 100, 6),
	(25414, 'MIRIGAMBAR', 1927, 100, 6),
	(25415, 'PODOREJO', 1927, 100, 6),
	(25416, 'SAMBIDOPLANG', 1927, 100, 6),
	(25417, 'SAMBIJAJAR', 1927, 100, 6),
	(25418, 'SAMBIROBYONG', 1927, 100, 6),
	(25419, 'SUMBERDADI', 1927, 100, 6),
	(25420, 'TAMBAKREJO', 1927, 100, 6),
	(25421, 'TRENCENG', 1927, 100, 6),
	(25422, 'WATES', 1927, 100, 6),
	(25423, 'WONOREJO', 1927, 100, 6),
	(25424, 'JENGGLUNGHARJO', 1928, 100, 6),
	(25425, 'KRESIKAN', 1928, 100, 6),
	(25426, 'NGEPOH', 1928, 100, 6),
	(25427, 'NGREJO', 1928, 100, 6),
	(25428, 'PAKISREJO', 1928, 100, 6),
	(25429, 'TANGGUNGGUNUNG', 1928, 100, 6),
	(25430, 'TENGGAREJO', 1928, 100, 6),
	(25431, 'BAGO', 1929, 100, 6),
	(25432, 'BOTORAN', 1929, 100, 6),
	(25433, 'JEPUN', 1929, 100, 6),
	(25434, 'KAMPUNGDALEM', 1929, 100, 6),
	(25435, 'KARANGWARU', 1929, 100, 6),
	(25436, 'KAUMAN', 1929, 100, 6),
	(25437, 'KEDUNGSOKO', 1929, 100, 6),
	(25438, 'KENAYAN', 1929, 100, 6),
	(25439, 'KEPATIHAN', 1929, 100, 6),
	(25440, 'KUTOANYAR', 1929, 100, 6),
	(25441, 'PANGGUNGREJO', 1929, 100, 6),
	(25442, 'SEMBUNG', 1929, 100, 6),
	(25443, 'TAMANAN', 1929, 100, 6),
	(25444, 'TERTEK', 1929, 100, 6),
	(25445, 'ABIANSEMAL', 1930, 101, 7),
	(25446, 'ANGANTAKA', 1930, 101, 7),
	(25447, 'AYUNAN', 1930, 101, 7),
	(25448, 'BLAHKIUH', 1930, 101, 7),
	(25449, 'BONGKASA', 1930, 101, 7),
	(25450, 'BONGKASA PERTIWI', 1930, 101, 7),
	(25451, 'DARMASABA', 1930, 101, 7),
	(25452, 'DAUH YEH CANI', 1930, 101, 7),
	(25453, 'JAGAPATI', 1930, 101, 7),
	(25454, 'MAMBAL', 1930, 101, 7),
	(25455, 'MEKAR BHUWANA', 1930, 101, 7),
	(25456, 'PUNGGUL', 1930, 101, 7),
	(25457, 'SANGEH', 1930, 101, 7),
	(25458, 'SEDANG', 1930, 101, 7),
	(25459, 'SELAT', 1930, 101, 7),
	(25460, 'SIBANG GEDE', 1930, 101, 7),
	(25461, 'SIBANG KAJA', 1930, 101, 7),
	(25462, 'TAMAN', 1930, 101, 7),
	(25463, 'KEDONGANAN', 1931, 101, 7),
	(25464, 'KUTA', 1931, 101, 7),
	(25465, 'LEGIAN', 1931, 101, 7),
	(25466, 'SEMINYAK', 1931, 101, 7),
	(25467, 'TUBAN', 1931, 101, 7),
	(25468, 'BENOA', 1932, 101, 7),
	(25469, 'JIMBARAN', 1932, 101, 7),
	(25470, 'KUTUH', 1932, 101, 7),
	(25471, 'PECATU', 1932, 101, 7),
	(25472, 'TANJUNG BENOA', 1932, 101, 7),
	(25473, 'UNGASAN', 1932, 101, 7),
	(25474, 'CANGGU', 1933, 101, 7),
	(25475, 'DALUNG', 1933, 101, 7),
	(25476, 'KEROBOKAN', 1933, 101, 7),
	(25477, 'KEROBOKAN KAJA', 1933, 101, 7),
	(25478, 'KEROBOKAN KELOD', 1933, 101, 7),
	(25479, 'TIBUBENENG', 1933, 101, 7),
	(25480, 'ABIANBASE', 1934, 101, 7),
	(25481, 'BAHA', 1934, 101, 7),
	(25482, 'BUDUK', 1934, 101, 7),
	(25483, 'CEMAGI', 1934, 101, 7),
	(25484, 'GULINGAN', 1934, 101, 7),
	(25485, 'KAPAL', 1934, 101, 7),
	(25486, 'KEKERAN', 1934, 101, 7),
	(25487, 'KUWUM', 1934, 101, 7),
	(25488, 'LUKLUK', 1934, 101, 7),
	(25489, 'MENGWI', 1934, 101, 7),
	(25490, 'MENGWITANI', 1934, 101, 7),
	(25491, 'MUNGGU', 1934, 101, 7),
	(25492, 'PENARUNGAN', 1934, 101, 7),
	(25493, 'PERERENAN', 1934, 101, 7),
	(25494, 'SADING', 1934, 101, 7),
	(25495, 'SEMBUNG', 1934, 101, 7),
	(25496, 'SEMPIDI', 1934, 101, 7),
	(25497, 'SOBANGAN', 1934, 101, 7),
	(25498, 'TUMBAK BAYUH', 1934, 101, 7),
	(25499, 'WERDI BHUWANA', 1934, 101, 7),
	(25500, 'BELOK', 1935, 101, 7),
	(25501, 'CARANGSARI', 1935, 101, 7),
	(25502, 'GETASAN', 1935, 101, 7),
	(25503, 'PANGSAN', 1935, 101, 7),
	(25504, 'PELAGA', 1935, 101, 7),
	(25505, 'PETANG', 1935, 101, 7),
	(25506, 'SULANGAI', 1935, 101, 7),
	(25507, 'BEBALANG', 1936, 102, 7),
	(25508, 'BUNUTIN', 1936, 102, 7),
	(25509, 'CEMPAGA', 1936, 102, 7),
	(25510, 'KAWAN', 1936, 102, 7),
	(25511, 'KAYUBIHI', 1936, 102, 7),
	(25512, 'KUBU', 1936, 102, 7),
	(25513, 'LANDIH', 1936, 102, 7),
	(25514, 'PENGOTAN', 1936, 102, 7),
	(25515, 'TAMAN BALI', 1936, 102, 7),
	(25516, 'ABANG SONGAN', 1937, 102, 7),
	(25517, 'ABUAN', 1937, 102, 7),
	(25518, 'AWAN', 1937, 102, 7),
	(25519, 'BANTANG', 1937, 102, 7),
	(25520, 'BANUA', 1937, 102, 7),
	(25521, 'BATU DINDING', 1937, 102, 7),
	(25522, 'BATUKAANG', 1937, 102, 7),
	(25523, 'BATUR SELATAN', 1937, 102, 7),
	(25524, 'BATUR TENGAH', 1937, 102, 7),
	(25525, 'BATUR UTARA', 1937, 102, 7),
	(25526, 'BAYUNGCERIK', 1937, 102, 7),
	(25527, 'BAYUNGGEDE', 1937, 102, 7),
	(25528, 'BELANCAN', 1937, 102, 7),
	(25529, 'BELANDINGAN', 1937, 102, 7),
	(25530, 'BELANGA', 1937, 102, 7),
	(25531, 'BELANTIH', 1937, 102, 7),
	(25532, 'BINYAN', 1937, 102, 7),
	(25533, 'BONYOH', 1937, 102, 7),
	(25534, 'BUAHAN', 1937, 102, 7),
	(25535, 'BUNUTIN', 1937, 102, 7),
	(25536, 'CATUR', 1937, 102, 7),
	(25537, 'DAUP', 1937, 102, 7),
	(25538, 'DAUSA', 1937, 102, 7),
	(25539, 'GUNUNGBAU', 1937, 102, 7),
	(25540, 'KATUNG', 1937, 102, 7),
	(25541, 'KEDISAN', 1937, 102, 7),
	(25542, 'KINTAMANI', 1937, 102, 7),
	(25543, 'KUTUH', 1937, 102, 7),
	(25544, 'LANGGAHAN', 1937, 102, 7),
	(25545, 'LEMBEAN', 1937, 102, 7),
	(25546, 'MANGGUH', 1937, 102, 7),
	(25547, 'MANIKLIYU', 1937, 102, 7),
	(25548, 'MENGANI', 1937, 102, 7),
	(25549, 'PENGEJARAN', 1937, 102, 7),
	(25550, 'PINGGAN', 1937, 102, 7),
	(25551, 'SATRA', 1937, 102, 7),
	(25552, 'SEKAAN', 1937, 102, 7),
	(25553, 'SEKARDADI', 1937, 102, 7),
	(25554, 'SELULUNG', 1937, 102, 7),
	(25555, 'SERAHI', 1937, 102, 7),
	(25556, 'SIYAKIN', 1937, 102, 7),
	(25557, 'SONGAN A', 1937, 102, 7),
	(25558, 'SONGAN B', 1937, 102, 7),
	(25559, 'SUBAYA', 1937, 102, 7),
	(25560, 'SUKAWANA', 1937, 102, 7),
	(25561, 'SUTER', 1937, 102, 7),
	(25562, 'TERUNYAN', 1937, 102, 7),
	(25563, 'ULIAN', 1937, 102, 7),
	(25564, 'ABUAN', 1938, 102, 7),
	(25565, 'APUAN', 1938, 102, 7),
	(25566, 'DEMULIH', 1938, 102, 7),
	(25567, 'PENGIANGAN', 1938, 102, 7),
	(25568, 'PENGLUMBARAN', 1938, 102, 7),
	(25569, 'SELAT', 1938, 102, 7),
	(25570, 'SULAHAN', 1938, 102, 7),
	(25571, 'SUSUT', 1938, 102, 7),
	(25572, 'TIGA', 1938, 102, 7),
	(25573, 'BANGBANG', 1939, 102, 7),
	(25574, 'JEHEM', 1939, 102, 7),
	(25575, 'PENINJOAN', 1939, 102, 7),
	(25576, 'TEMBUKU', 1939, 102, 7),
	(25577, 'UNDISAN', 1939, 102, 7),
	(25578, 'YANGAPI', 1939, 102, 7),
	(25579, 'BALOKANG', 1940, 103, 7),
	(25580, 'BANDUNG', 1940, 103, 7),
	(25581, 'BANJAR', 1940, 103, 7),
	(25582, 'BANJAR', 1940, 103, 7),
	(25583, 'BANJAR', 1940, 103, 7),
	(25584, 'BANJAR TEGEHA', 1940, 103, 7),
	(25585, 'BANYUATIS', 1940, 103, 7),
	(25586, 'BANYUSRI', 1940, 103, 7),
	(25587, 'CEMPAGA', 1940, 103, 7),
	(25588, 'CIBEUREUM', 1940, 103, 7),
	(25589, 'CIBEUREUM', 1940, 103, 7),
	(25590, 'CIBODAS', 1940, 103, 7),
	(25591, 'CITALAHAB', 1940, 103, 7),
	(25592, 'DENCARIK', 1940, 103, 7),
	(25593, 'GESING', 1940, 103, 7),
	(25594, 'GOBLEG', 1940, 103, 7),
	(25595, 'GUNUNGPUTRI', 1940, 103, 7),
	(25596, 'JAJAWAR', 1940, 103, 7),
	(25597, 'KADUBALE', 1940, 103, 7),
	(25598, 'KADULIMUS', 1940, 103, 7),
	(25599, 'KADUMANEUH', 1940, 103, 7),
	(25600, 'KALIASEM', 1940, 103, 7),
	(25601, 'KAYUPUTIH', 1940, 103, 7),
	(25602, 'MEKARSARI', 1940, 103, 7),
	(25603, 'MOGANA', 1940, 103, 7),
	(25604, 'MUNDUK', 1940, 103, 7),
	(25605, 'NEGLASARI', 1940, 103, 7),
	(25606, 'PASIRAWI', 1940, 103, 7),
	(25607, 'PEDAWA', 1940, 103, 7),
	(25608, 'SIDETAPA', 1940, 103, 7),
	(25609, 'SITUBATU', 1940, 103, 7),
	(25610, 'TAMPEKAN', 1940, 103, 7),
	(25611, 'TEMUKUS', 1940, 103, 7),
	(25612, 'TIGAWASA', 1940, 103, 7),
	(25613, 'TIRTASARI', 1940, 103, 7),
	(25614, 'ALASANGKER', 1941, 103, 7),
	(25615, 'ANTURAN', 1941, 103, 7),
	(25616, 'ASTINA', 1941, 103, 7),
	(25617, 'BANJAR BALI', 1941, 103, 7),
	(25618, 'BANJAR JAWA', 1941, 103, 7),
	(25619, 'BANJAR TEGAL', 1941, 103, 7),
	(25620, 'BANYUASRI', 1941, 103, 7),
	(25621, 'BANYUNING', 1941, 103, 7),
	(25622, 'BERATAN', 1941, 103, 7),
	(25623, 'BHAKTI SERAGA (BAKTISERAGA)', 1941, 103, 7),
	(25624, 'JINENGDALEM', 1941, 103, 7),
	(25625, 'KALIBUKBUK', 1941, 103, 7),
	(25626, 'KALIUNTU', 1941, 103, 7),
	(25627, 'KAMPUNG ANYAR', 1941, 103, 7),
	(25628, 'KAMPUNG BARU', 1941, 103, 7),
	(25629, 'KAMPUNG BUGIS', 1941, 103, 7),
	(25630, 'KAMPUNG KAJANAN', 1941, 103, 7),
	(25631, 'KAMPUNG SINGARAJA', 1941, 103, 7),
	(25632, 'KENDRAN', 1941, 103, 7),
	(25633, 'LILIGUNDI', 1941, 103, 7),
	(25634, 'NAGASEPAHA', 1941, 103, 7),
	(25635, 'PAKET AGUNG', 1941, 103, 7),
	(25636, 'PEMARON', 1941, 103, 7),
	(25637, 'PENARUKAN', 1941, 103, 7),
	(25638, 'PENGLATAN', 1941, 103, 7),
	(25639, 'PETANDAKAN', 1941, 103, 7),
	(25640, 'POH BERGONG', 1941, 103, 7),
	(25641, 'SARI MEKAR', 1941, 103, 7),
	(25642, 'TUKADMUNGGA', 1941, 103, 7),
	(25643, 'BENGKEL', 1942, 103, 7),
	(25644, 'BONGANCINA', 1942, 103, 7),
	(25645, 'BUSUNGBIU', 1942, 103, 7),
	(25646, 'KEDIA (KEDIS)', 1942, 103, 7),
	(25647, 'KEKERAN', 1942, 103, 7),
	(25648, 'PELAPUAN', 1942, 103, 7),
	(25649, 'PUCAKSARI', 1942, 103, 7),
	(25650, 'SEPANG', 1942, 103, 7),
	(25651, 'SEPANG KELOD', 1942, 103, 7),
	(25652, 'SUBUK', 1942, 103, 7),
	(25653, 'TELAGA', 1942, 103, 7),
	(25654, 'TINGGARSARI', 1942, 103, 7),
	(25655, 'TISTA', 1942, 103, 7),
	(25656, 'TITAB', 1942, 103, 7),
	(25657, 'UMEJERO', 1942, 103, 7),
	(25658, 'BANYUPOH', 1943, 103, 7),
	(25659, 'CELUKAN BAWANG', 1943, 103, 7),
	(25660, 'GEROKGAK', 1943, 103, 7),
	(25661, 'MUSI', 1943, 103, 7),
	(25662, 'PATAS', 1943, 103, 7),
	(25663, 'PEJARAKAN', 1943, 103, 7),
	(25664, 'PEMUTERAN', 1943, 103, 7),
	(25665, 'PENGULON', 1943, 103, 7),
	(25666, 'PENYABANGAN', 1943, 103, 7),
	(25667, 'SANGGALANGIT', 1943, 103, 7),
	(25668, 'SUMBER KLAMPOK', 1943, 103, 7),
	(25669, 'SUMBERKIMA', 1943, 103, 7),
	(25670, 'TINGA TINGA', 1943, 103, 7),
	(25671, 'TUKAD SUMAGA', 1943, 103, 7),
	(25672, 'BENGKALA', 1944, 103, 7),
	(25673, 'BILA', 1944, 103, 7),
	(25674, 'BONTIHING', 1944, 103, 7),
	(25675, 'BUKTI', 1944, 103, 7),
	(25676, 'BULIAN', 1944, 103, 7),
	(25677, 'DEPEHA', 1944, 103, 7),
	(25678, 'KUBUTAMBAHAN', 1944, 103, 7),
	(25679, 'MENGENING', 1944, 103, 7),
	(25680, 'PAKISAN', 1944, 103, 7),
	(25681, 'TAJUN', 1944, 103, 7),
	(25682, 'TAMBAKAN', 1944, 103, 7),
	(25683, 'TAMBLANG', 1944, 103, 7),
	(25684, 'TUNJUNG', 1944, 103, 7),
	(25685, 'BEBETIN', 1945, 103, 7),
	(25686, 'BUNGKULAN', 1945, 103, 7),
	(25687, 'GALUNGAN', 1945, 103, 7),
	(25688, 'GIRI EMAS', 1945, 103, 7),
	(25689, 'JAGARAGA', 1945, 103, 7),
	(25690, 'KEROBOKAN', 1945, 103, 7),
	(25691, 'LEMUKIH', 1945, 103, 7),
	(25692, 'MENYALI', 1945, 103, 7),
	(25693, 'SANGSIT', 1945, 103, 7),
	(25694, 'SAWAN', 1945, 103, 7),
	(25695, 'SEKUMPUL', 1945, 103, 7),
	(25696, 'SINABUN', 1945, 103, 7),
	(25697, 'SUDAJI', 1945, 103, 7),
	(25698, 'SUWUG', 1945, 103, 7),
	(25699, 'BANJAR ASEM', 1946, 103, 7),
	(25700, 'BESTALA', 1946, 103, 7),
	(25701, 'BUBUNAN', 1946, 103, 7),
	(25702, 'GUNUNGSARI', 1946, 103, 7),
	(25703, 'JOANYAR', 1946, 103, 7),
	(25704, 'KALIANGET', 1946, 103, 7),
	(25705, 'KALISADA', 1946, 103, 7),
	(25706, 'LOKAPAKSA', 1946, 103, 7),
	(25707, 'MAYONG', 1946, 103, 7),
	(25708, 'MUNDUK BESTALA', 1946, 103, 7),
	(25709, 'PANGKUNGPARUK', 1946, 103, 7),
	(25710, 'PATEMOH (PATEMON)', 1946, 103, 7),
	(25711, 'PENGASTULAN', 1946, 103, 7),
	(25712, 'RANGDU', 1946, 103, 7),
	(25713, 'RINGDIKIT', 1946, 103, 7),
	(25714, 'SERIRIT', 1946, 103, 7),
	(25715, 'SULANYAH', 1946, 103, 7),
	(25716, 'TANGGUWISIA', 1946, 103, 7),
	(25717, 'ULARAN', 1946, 103, 7),
	(25718, 'UMEANYAR', 1946, 103, 7),
	(25719, 'UNGGAHAN', 1946, 103, 7),
	(25720, 'AMBENGAN', 1947, 103, 7),
	(25721, 'GITGIT', 1947, 103, 7),
	(25722, 'KAYUPUTIH', 1947, 103, 7),
	(25723, 'PADANGBULIA', 1947, 103, 7),
	(25724, 'PANCASARI', 1947, 103, 7),
	(25725, 'PANJI', 1947, 103, 7),
	(25726, 'PANJI ANOM', 1947, 103, 7),
	(25727, 'PEGADUNGAN', 1947, 103, 7),
	(25728, 'PEGAYAMAN', 1947, 103, 7),
	(25729, 'SAMBANGAN', 1947, 103, 7),
	(25730, 'SELAT', 1947, 103, 7),
	(25731, 'SILANGJANA', 1947, 103, 7),
	(25732, 'SUKASADA', 1947, 103, 7),
	(25733, 'TEGAL LINGGAH (TEGALINGGAH)', 1947, 103, 7),
	(25734, 'WANAGIRI', 1947, 103, 7),
	(25735, 'BONDALEM', 1948, 103, 7),
	(25736, 'JULAH', 1948, 103, 7),
	(25737, 'LES', 1948, 103, 7),
	(25738, 'MADENAN', 1948, 103, 7),
	(25739, 'PACUNG', 1948, 103, 7),
	(25740, 'PENUKTUKAN', 1948, 103, 7),
	(25741, 'SAMBIRENTENG', 1948, 103, 7),
	(25742, 'SEMBIRAN', 1948, 103, 7),
	(25743, 'TEJAKULA', 1948, 103, 7),
	(25744, 'TEMBOK', 1948, 103, 7),
	(25745, 'DAUH PURI', 1949, 104, 7),
	(25746, 'DAUH PURI KANGIN', 1949, 104, 7),
	(25747, 'DAUH PURI KAUH', 1949, 104, 7),
	(25748, 'DAUH PURI KLOD/KELOD', 1949, 104, 7),
	(25749, 'PADANGSAMBIAN', 1949, 104, 7),
	(25750, 'PADANGSAMBIAN KAJA', 1949, 104, 7),
	(25751, 'PADANGSAMBIAN KLOD/KELOD', 1949, 104, 7),
	(25752, 'PEMECUTAN', 1949, 104, 7),
	(25753, 'PEMECUTAN KLOD/KELOD', 1949, 104, 7),
	(25754, 'TEGAL HARUM', 1949, 104, 7),
	(25755, 'TEGAL KERTHA', 1949, 104, 7),
	(25756, 'PANJER', 1950, 104, 7),
	(25757, 'PEDUNGAN', 1950, 104, 7),
	(25758, 'PEMOGAN', 1950, 104, 7),
	(25759, 'RENON', 1950, 104, 7),
	(25760, 'SANUR', 1950, 104, 7),
	(25761, 'SANUR KAJA', 1950, 104, 7),
	(25762, 'SANUR KAUH', 1950, 104, 7),
	(25763, 'SERANGAN', 1950, 104, 7),
	(25764, 'SESETAN', 1950, 104, 7),
	(25765, 'SIDAKARYA', 1950, 104, 7),
	(25766, 'DANGIN PURI', 1951, 104, 7),
	(25767, 'DANGIN PURI KLOD', 1951, 104, 7),
	(25768, 'KESIMAN', 1951, 104, 7),
	(25769, 'KESIMAN KERTALANGU', 1951, 104, 7),
	(25770, 'KESIMAN PETILAN', 1951, 104, 7),
	(25771, 'PENATIH', 1951, 104, 7),
	(25772, 'PENATIH DANGIN PURI', 1951, 104, 7),
	(25773, 'SUMERTA', 1951, 104, 7),
	(25774, 'SUMERTA KAJA', 1951, 104, 7),
	(25775, 'SUMERTA KAUH', 1951, 104, 7),
	(25776, 'SUMERTA KELOD/KLOD', 1951, 104, 7),
	(25777, 'DANGIN PURI KAJA', 1952, 104, 7),
	(25778, 'DANGIN PURI KANGIN', 1952, 104, 7),
	(25779, 'DANGIN PURI KAUH', 1952, 104, 7),
	(25780, 'DAUH PURI KAJA', 1952, 104, 7),
	(25781, 'PEGUYANGAN', 1952, 104, 7),
	(25782, 'PEGUYANGAN KAJA', 1952, 104, 7),
	(25783, 'PEGUYANGAN KANGIN', 1952, 104, 7),
	(25784, 'PEMECUTAN KAJA', 1952, 104, 7),
	(25785, 'TONJA', 1952, 104, 7),
	(25786, 'UBUNG', 1952, 104, 7),
	(25787, 'UBUNG KAJA', 1952, 104, 7),
	(25788, 'BEDULU', 1953, 105, 7),
	(25789, 'BELEGA', 1953, 105, 7),
	(25790, 'BLAHBATUH', 1953, 105, 7),
	(25791, 'BONA', 1953, 105, 7),
	(25792, 'BURUAN', 1953, 105, 7),
	(25793, 'KERAMAS', 1953, 105, 7),
	(25794, 'MEDAHAN', 1953, 105, 7),
	(25795, 'PERING', 1953, 105, 7),
	(25796, 'SABA', 1953, 105, 7),
	(25797, 'ABIANBASE', 1954, 105, 7),
	(25798, 'BAKBAKAN', 1954, 105, 7),
	(25799, 'BENG', 1954, 105, 7),
	(25800, 'BITERA', 1954, 105, 7),
	(25801, 'GIANYAR', 1954, 105, 7),
	(25802, 'LEBIH', 1954, 105, 7),
	(25803, 'PETAK', 1954, 105, 7),
	(25804, 'PETAK KAJA', 1954, 105, 7),
	(25805, 'SAMPLANGAN', 1954, 105, 7),
	(25806, 'SERONGGA', 1954, 105, 7),
	(25807, 'SIANGAN', 1954, 105, 7),
	(25808, 'SIDAN', 1954, 105, 7),
	(25809, 'SUMITA', 1954, 105, 7),
	(25810, 'SUWAT', 1954, 105, 7),
	(25811, 'TEGAL TUGU', 1954, 105, 7),
	(25812, 'TEMESI', 1954, 105, 7),
	(25813, 'TULIKUP', 1954, 105, 7),
	(25814, 'BRESELA (BERESELA)', 1955, 105, 7),
	(25815, 'BUAHAN', 1955, 105, 7),
	(25816, 'BUAHAN KAJA', 1955, 105, 7),
	(25817, 'BUKIAN', 1955, 105, 7),
	(25818, 'KELUSA', 1955, 105, 7),
	(25819, 'KERTA', 1955, 105, 7),
	(25820, 'MELINGGIH', 1955, 105, 7),
	(25821, 'MELINGGIH KELOD', 1955, 105, 7),
	(25822, 'PUHU', 1955, 105, 7),
	(25823, 'BATUAN', 1956, 105, 7),
	(25824, 'BATUAN KALER', 1956, 105, 7),
	(25825, 'BATUBULAN', 1956, 105, 7),
	(25826, 'BATUBULAN KANGIN', 1956, 105, 7),
	(25827, 'CELUK', 1956, 105, 7),
	(25828, 'GUWANG', 1956, 105, 7),
	(25829, 'KEMENUH', 1956, 105, 7),
	(25830, 'KETEWEL', 1956, 105, 7),
	(25831, 'SINGAPADU', 1956, 105, 7),
	(25832, 'SINGAPADU KALER', 1956, 105, 7),
	(25833, 'SINGAPADU TENGAH', 1956, 105, 7),
	(25834, 'SUKAWATI', 1956, 105, 7),
	(25835, 'MANUKAYA', 1957, 105, 7),
	(25836, 'PEJENG', 1957, 105, 7),
	(25837, 'PEJENG KAJA', 1957, 105, 7),
	(25838, 'PEJENG KANGIN', 1957, 105, 7),
	(25839, 'PEJENG KAWAN', 1957, 105, 7),
	(25840, 'PEJENG KLOD (KELOD)', 1957, 105, 7),
	(25841, 'SANDING', 1957, 105, 7),
	(25842, 'TAMPAKSIRING', 1957, 105, 7),
	(25843, 'KEDISAN', 1958, 105, 7),
	(25844, 'KELIKI', 1958, 105, 7),
	(25845, 'KENDERAN', 1958, 105, 7),
	(25846, 'PUPUAN', 1958, 105, 7),
	(25847, 'SEBATU', 1958, 105, 7),
	(25848, 'TARO', 1958, 105, 7),
	(25849, 'TEGALLALANG', 1958, 105, 7),
	(25850, 'KEDEWATAN', 1959, 105, 7),
	(25851, 'LODTUNDUH', 1959, 105, 7),
	(25852, 'MAS', 1959, 105, 7),
	(25853, 'PELIATAN', 1959, 105, 7),
	(25854, 'PETULU', 1959, 105, 7),
	(25855, 'SAYAN', 1959, 105, 7),
	(25856, 'SINGAKERTA (SINGEKERTA)', 1959, 105, 7),
	(25857, 'UBUD', 1959, 105, 7),
	(25858, 'AIR KUNING', 1960, 106, 7),
	(25859, 'BATUAGUNG', 1960, 106, 7),
	(25860, 'BUDENG', 1960, 106, 7),
	(25861, 'DANGIN TUKADAYA', 1960, 106, 7),
	(25862, 'DAUHWARU', 1960, 106, 7),
	(25863, 'LOLOAN TIMUR', 1960, 106, 7),
	(25864, 'PENDEM', 1960, 106, 7),
	(25865, 'PERANCAK', 1960, 106, 7),
	(25866, 'SANGKARAGUNG', 1960, 106, 7),
	(25867, 'YEH KUNING', 1960, 106, 7),
	(25868, 'BLIMBINGSARI (BELIMBINGSARI)', 1961, 106, 7),
	(25869, 'CANDIKUSUMA', 1961, 106, 7),
	(25870, 'EKASARI', 1961, 106, 7),
	(25871, 'GILIMANUK', 1961, 106, 7),
	(25872, 'MANISTUTU', 1961, 106, 7),
	(25873, 'MELAYA', 1961, 106, 7),
	(25874, 'NUSA SARI', 1961, 106, 7),
	(25875, 'TUKADAYA', 1961, 106, 7),
	(25876, 'TUWED', 1961, 106, 7),
	(25877, 'WARNASARI', 1961, 106, 7),
	(25878, 'DELOD BERAWAH', 1962, 106, 7),
	(25879, 'MENDOYO DANGIN TUKAD', 1962, 106, 7),
	(25880, 'MENDOYO DAUH TUKAD', 1962, 106, 7),
	(25881, 'PENYARINGAN', 1962, 106, 7),
	(25882, 'PERGUNG', 1962, 106, 7),
	(25883, 'POHSANTEN', 1962, 106, 7),
	(25884, 'TEGAL CANGKRING', 1962, 106, 7),
	(25885, 'YEH EMBANG', 1962, 106, 7),
	(25886, 'YEH EMBANG KANGIN', 1962, 106, 7),
	(25887, 'YEH EMBANG KAUH', 1962, 106, 7),
	(25888, 'YEH SUMBUL', 1962, 106, 7),
	(25889, 'BALER BALE AGUNG', 1963, 106, 7),
	(25890, 'BALUK', 1963, 106, 7),
	(25891, 'BANJAR TENGAH', 1963, 106, 7),
	(25892, 'BANYUBIRU', 1963, 106, 7),
	(25893, 'BERANGBANG', 1963, 106, 7),
	(25894, 'CUPEL', 1963, 106, 7),
	(25895, 'KALIAKAH', 1963, 106, 7),
	(25896, 'LELATENG', 1963, 106, 7),
	(25897, 'LOLOAN BARAT', 1963, 106, 7),
	(25898, 'PENGAMBENGAN', 1963, 106, 7),
	(25899, 'TEGAL BADENG BARAT', 1963, 106, 7),
	(25900, 'TEGAL BADENG TIMUR', 1963, 106, 7),
	(25901, 'ASAHDUREN', 1964, 106, 7),
	(25902, 'GUMBRIH', 1964, 106, 7),
	(25903, 'MANGGISSARI', 1964, 106, 7),
	(25904, 'MEDEWI', 1964, 106, 7),
	(25905, 'PANGYANGAN', 1964, 106, 7),
	(25906, 'PEKUTATAN', 1964, 106, 7),
	(25907, 'PENGERAGOAN (PENGRAGOAN)', 1964, 106, 7),
	(25908, 'PULUKAN', 1964, 106, 7),
	(25909, 'ABABI', 1965, 107, 7),
	(25910, 'ABANG', 1965, 107, 7),
	(25911, 'BUNUTAN', 1965, 107, 7),
	(25912, 'CULIK', 1965, 107, 7),
	(25913, 'DATAH', 1965, 107, 7),
	(25914, 'KERTA MANDALA', 1965, 107, 7),
	(25915, 'KESIMPAR', 1965, 107, 7),
	(25916, 'LABA SARI', 1965, 107, 7),
	(25917, 'NAWAKERTI', 1965, 107, 7),
	(25918, 'PIDPID', 1965, 107, 7),
	(25919, 'PURWAKERTI', 1965, 107, 7),
	(25920, 'TISTA', 1965, 107, 7),
	(25921, 'TIYINGTALI', 1965, 107, 7),
	(25922, 'TRI BUANA', 1965, 107, 7),
	(25923, 'BEBANDEM', 1966, 107, 7),
	(25924, 'BUANA GIRI', 1966, 107, 7),
	(25925, 'BUDAKELING (BUDE KELING)', 1966, 107, 7),
	(25926, 'BUNGAYA (BUNGAYA KAUH)', 1966, 107, 7),
	(25927, 'BUNGAYA KANGIN', 1966, 107, 7),
	(25928, 'JUNGUTAN', 1966, 107, 7),
	(25929, 'MACANG', 1966, 107, 7),
	(25930, 'SIBETAN', 1966, 107, 7),
	(25931, 'BUGBUG', 1967, 107, 7),
	(25932, 'BUKIT', 1967, 107, 7),
	(25933, 'KARANGASEM', 1967, 107, 7),
	(25934, 'PADANG KERTA', 1967, 107, 7),
	(25935, 'PERTIMA', 1967, 107, 7),
	(25936, 'SERAYA BARAT', 1967, 107, 7),
	(25937, 'SERAYA TENGAH', 1967, 107, 7),
	(25938, 'SERAYA TIMUR', 1967, 107, 7),
	(25939, 'SUBAGAN', 1967, 107, 7),
	(25940, 'TEGALLINGGAH', 1967, 107, 7),
	(25941, 'TUMBU', 1967, 107, 7),
	(25942, 'ANTIGA', 1968, 107, 7),
	(25943, 'ANTIGA KELOD', 1968, 107, 7),
	(25944, 'GEGELANG', 1968, 107, 7),
	(25945, 'MANGGIS', 1968, 107, 7),
	(25946, 'NGIS', 1968, 107, 7),
	(25947, 'NYUH TEBEL', 1968, 107, 7),
	(25948, 'PADANGBAI', 1968, 107, 7),
	(25949, 'PESEDAHAN', 1968, 107, 7),
	(25950, 'SELUMBUNG', 1968, 107, 7),
	(25951, 'SENGKIDU', 1968, 107, 7),
	(25952, 'TENGANAN', 1968, 107, 7),
	(25953, 'ULAKAN', 1968, 107, 7),
	(25954, 'BESAKIH', 1969, 107, 7),
	(25955, 'MENANGA', 1969, 107, 7),
	(25956, 'NONGAN', 1969, 107, 7),
	(25957, 'PEMPATAN', 1969, 107, 7),
	(25958, 'PESABAN', 1969, 107, 7),
	(25959, 'RENDANG', 1969, 107, 7),
	(25960, 'AMERTA BHUANA', 1970, 107, 7),
	(25961, 'DUDA', 1970, 107, 7),
	(25962, 'DUDA TIMUR', 1970, 107, 7),
	(25963, 'DUDA UTARA', 1970, 107, 7),
	(25964, 'MUNCAN', 1970, 107, 7),
	(25965, 'MURUNG KERAMAT', 1970, 107, 7),
	(25966, 'PANAMAS', 1970, 107, 7),
	(25967, 'PERING SARI', 1970, 107, 7),
	(25968, 'PULAU TELO', 1970, 107, 7),
	(25969, 'PULAU TELO BARU', 1970, 107, 7),
	(25970, 'SEBUDI', 1970, 107, 7),
	(25971, 'SELAT', 1970, 107, 7),
	(25972, 'SELAT BARAT', 1970, 107, 7),
	(25973, 'SELAT DALAM', 1970, 107, 7),
	(25974, 'SELAT HILIR', 1970, 107, 7),
	(25975, 'SELAT HULU', 1970, 107, 7),
	(25976, 'SELAT TENGAH', 1970, 107, 7),
	(25977, 'SELAT UTARA', 1970, 107, 7),
	(25978, 'KERTHA BUANA', 1971, 107, 7),
	(25979, 'LOKASARI', 1971, 107, 7),
	(25980, 'SANGKAN GUNUNG', 1971, 107, 7),
	(25981, 'SIDEMEN', 1971, 107, 7),
	(25982, 'SINDU WATI', 1971, 107, 7),
	(25983, 'TALIBENG', 1971, 107, 7),
	(25984, 'TANGKUP', 1971, 107, 7),
	(25985, 'TELAGA TAWANG', 1971, 107, 7),
	(25986, 'TRI EKA BUANA', 1971, 107, 7),
	(25987, 'WISMA KERTA', 1971, 107, 7),
	(25988, 'AAN', 1972, 108, 7),
	(25989, 'BAKAS', 1972, 108, 7),
	(25990, 'BANJARANGKAN', 1972, 108, 7),
	(25991, 'BUNGBUNGAN', 1972, 108, 7),
	(25992, 'GETAKAN', 1972, 108, 7),
	(25993, 'NEGARI', 1972, 108, 7),
	(25994, 'NYALIAN', 1972, 108, 7),
	(25995, 'NYANGLAN', 1972, 108, 7),
	(25996, 'TAKMUNG', 1972, 108, 7),
	(25997, 'TIHINGAN', 1972, 108, 7),
	(25998, 'TIMUHUN', 1972, 108, 7),
	(25999, 'TOHPATI', 1972, 108, 7),
	(26000, 'TUSAN', 1972, 108, 7)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 25")
}
