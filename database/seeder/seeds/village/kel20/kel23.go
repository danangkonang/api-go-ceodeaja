package kel20

import (
	"fmt"

	"github.com/danangkonang/rest-api/config"
)

func Kel23() {
	db := config.Connect()
	_, err := db.Exec(`INSERT INTO villages (id,village,sub_district_id,regency_id,province_id) VALUES
	(23001, 'BROTO', 1751, 91, 6),
	(23002, 'CALUK', 1751, 91, 6),
	(23003, 'CRABAK', 1751, 91, 6),
	(23004, 'DURI', 1751, 91, 6),
	(23005, 'GALAK', 1751, 91, 6),
	(23006, 'GOMBANG', 1751, 91, 6),
	(23007, 'GUNDIK', 1751, 91, 6),
	(23008, 'JANTI', 1751, 91, 6),
	(23009, 'JEBENG (JABENG)', 1751, 91, 6),
	(23010, 'KAMBENG', 1751, 91, 6),
	(23011, 'MENGGARE', 1751, 91, 6),
	(23012, 'MOJOPITU', 1751, 91, 6),
	(23013, 'NAILAN', 1751, 91, 6),
	(23014, 'NGILO-ILO', 1751, 91, 6),
	(23015, 'NGLONING', 1751, 91, 6),
	(23016, 'PLANCUNGAN', 1751, 91, 6),
	(23017, 'SENEPO', 1751, 91, 6),
	(23018, 'SIMO', 1751, 91, 6),
	(23019, 'SLAHUNG', 1751, 91, 6),
	(23020, 'TRUNENG', 1751, 91, 6),
	(23021, 'TUGUREJO', 1751, 91, 6),
	(23022, 'WATES', 1751, 91, 6),
	(23023, 'BEDOHO', 1752, 91, 6),
	(23024, 'BLIMBINGSARI', 1752, 91, 6),
	(23025, 'BRANGKAL', 1752, 91, 6),
	(23026, 'GEMEKAN', 1752, 91, 6),
	(23027, 'JAMPIROGO', 1752, 91, 6),
	(23028, 'JAPAN', 1752, 91, 6),
	(23029, 'JURUG', 1752, 91, 6),
	(23030, 'KARANGKEDAWANG', 1752, 91, 6),
	(23031, 'KEDUNGMALING', 1752, 91, 6),
	(23032, 'KLEPU', 1752, 91, 6),
	(23033, 'KLINTEREJO', 1752, 91, 6),
	(23034, 'MODONGAN', 1752, 91, 6),
	(23035, 'MOJORANU', 1752, 91, 6),
	(23036, 'NGADIROJO (NGADIREJO)', 1752, 91, 6),
	(23037, 'NGINGASREMBYONG', 1752, 91, 6),
	(23038, 'SAMBIROTO', 1752, 91, 6),
	(23039, 'SOOKO', 1752, 91, 6),
	(23040, 'SOOKO', 1752, 91, 6),
	(23041, 'SURU (SERU)', 1752, 91, 6),
	(23042, 'TEMPURAN', 1752, 91, 6),
	(23043, 'WRINGINREJO', 1752, 91, 6),
	(23044, 'BANTARAN', 1753, 92, 6),
	(23045, 'BESUK', 1753, 92, 6),
	(23046, 'GUNUNG TUGEL', 1753, 92, 6),
	(23047, 'KARANG ANYAR', 1753, 92, 6),
	(23048, 'KEDUNG REJO', 1753, 92, 6),
	(23049, 'KRAMAT AGUNG (KRAMATUNG)', 1753, 92, 6),
	(23050, 'KROPAK', 1753, 92, 6),
	(23051, 'LEGUNDI', 1753, 92, 6),
	(23052, 'PATOKAN', 1753, 92, 6),
	(23053, 'TEMPURAN', 1753, 92, 6),
	(23054, 'ALASSAPI', 1754, 92, 6),
	(23055, 'BANYUANYAR KIDUL', 1754, 92, 6),
	(23056, 'BANYUANYAR TENGAH', 1754, 92, 6),
	(23057, 'BLADO WETAN', 1754, 92, 6),
	(23058, 'GADING KULON', 1754, 92, 6),
	(23059, 'GUNUNG GENI', 1754, 92, 6),
	(23060, 'KLENANG KIDUL', 1754, 92, 6),
	(23061, 'KLENANG LOR', 1754, 92, 6),
	(23062, 'LIPRAK KIDUL', 1754, 92, 6),
	(23063, 'LIPRAK KULON', 1754, 92, 6),
	(23064, 'LIPRAK WETAN', 1754, 92, 6),
	(23065, 'PENDIL', 1754, 92, 6),
	(23066, 'SENTULAN', 1754, 92, 6),
	(23067, 'TAROKAN', 1754, 92, 6),
	(23068, 'ALAS KANDANG', 1755, 92, 6),
	(23069, 'ALAS NYIUR', 1755, 92, 6),
	(23070, 'ALAS SUMUR LOR', 1755, 92, 6),
	(23071, 'ALAS TENGAH', 1755, 92, 6),
	(23072, 'BAGO', 1755, 92, 6),
	(23073, 'BESUK AGUNG', 1755, 92, 6),
	(23074, 'BESUK KIDUL', 1755, 92, 6),
	(23075, 'JAMBANGAN', 1755, 92, 6),
	(23076, 'KECIK', 1755, 92, 6),
	(23077, 'KLAMPOKAN', 1755, 92, 6),
	(23078, 'KRAMPILAN', 1755, 92, 6),
	(23079, 'MATEKAN', 1755, 92, 6),
	(23080, 'RANDU JALAK', 1755, 92, 6),
	(23081, 'SINDET ANYAR', 1755, 92, 6),
	(23082, 'SINDET LAMI', 1755, 92, 6),
	(23083, 'SUMBERAN', 1755, 92, 6),
	(23084, 'SUMUR DALAM', 1755, 92, 6),
	(23085, 'DRINGU', 1756, 92, 6),
	(23086, 'KALIREJO', 1756, 92, 6),
	(23087, 'KALISALAM', 1756, 92, 6),
	(23088, 'KEDUNGDALEM', 1756, 92, 6),
	(23089, 'MRANGGON LAWANG', 1756, 92, 6),
	(23090, 'NGEPOH', 1756, 92, 6),
	(23091, 'PABEAN', 1756, 92, 6),
	(23092, 'RANDUPUTIH', 1756, 92, 6),
	(23093, 'SEKARKARE', 1756, 92, 6),
	(23094, 'SUMBER AGUNG', 1756, 92, 6),
	(23095, 'SUMBER SUKO', 1756, 92, 6),
	(23096, 'TAMAN SARI', 1756, 92, 6),
	(23097, 'TEGALREJO', 1756, 92, 6),
	(23098, 'WATUWUNGKUK', 1756, 92, 6),
	(23099, 'BATEKTAMAN (BETEKTAMAN)', 1757, 92, 6),
	(23100, 'BATUR', 1757, 92, 6),
	(23101, 'BULUPANDAK', 1757, 92, 6),
	(23102, 'CONDONG', 1757, 92, 6),
	(23103, 'DANDANG', 1757, 92, 6),
	(23104, 'DUREN', 1757, 92, 6),
	(23105, 'GADING WETAN', 1757, 92, 6),
	(23106, 'JURANGJERO', 1757, 92, 6),
	(23107, 'KALIACAR', 1757, 92, 6),
	(23108, 'KEBEN', 1757, 92, 6),
	(23109, 'KERTOSONO', 1757, 92, 6),
	(23110, 'MOJOLEGI', 1757, 92, 6),
	(23111, 'NOGOSAREN', 1757, 92, 6),
	(23112, 'PRASI', 1757, 92, 6),
	(23113, 'RANUWURUNG', 1757, 92, 6),
	(23114, 'RENTENG', 1757, 92, 6),
	(23115, 'SENTUL', 1757, 92, 6),
	(23116, 'SUMBERSECANG', 1757, 92, 6),
	(23117, 'WANGKAL', 1757, 92, 6),
	(23118, 'BANYUANYAR LOR', 1758, 92, 6),
	(23119, 'BRUMBUNGAN LOR', 1758, 92, 6),
	(23120, 'BULANG', 1758, 92, 6),
	(23121, 'CURAHSAWO', 1758, 92, 6),
	(23122, 'GENDING', 1758, 92, 6),
	(23123, 'JATIADI', 1758, 92, 6),
	(23124, 'KLASEMAN', 1758, 92, 6),
	(23125, 'PAJURANGAN', 1758, 92, 6),
	(23126, 'PESISIR', 1758, 92, 6),
	(23127, 'PIKATAN', 1758, 92, 6),
	(23128, 'RANDUPITU', 1758, 92, 6),
	(23129, 'SEBAUNG', 1758, 92, 6),
	(23130, 'SUMBERKERANG', 1758, 92, 6),
	(23131, 'BENDOSARI', 1759, 92, 6),
	(23132, 'DARUNGAN', 1759, 92, 6),
	(23133, 'DAWUHAN', 1759, 92, 6),
	(23134, 'JIMBE', 1759, 92, 6),
	(23135, 'KADEMANGAN', 1759, 92, 6),
	(23136, 'KADEMANGAN', 1759, 92, 6),
	(23137, 'KEBONSARI', 1759, 92, 6),
	(23138, 'KETAPANG', 1759, 92, 6),
	(23139, 'MARON', 1759, 92, 6),
	(23140, 'PAKISAJI', 1759, 92, 6),
	(23141, 'PANGGUNGDUWET', 1759, 92, 6),
	(23142, 'PILANG', 1759, 92, 6),
	(23143, 'PLOSOREJO', 1759, 92, 6),
	(23144, 'PLUMPUNGREJO', 1759, 92, 6),
	(23145, 'POHSANGIT KIDUL', 1759, 92, 6),
	(23146, 'REJOWINANGUN', 1759, 92, 6),
	(23147, 'SUMBERJATI', 1759, 92, 6),
	(23148, 'SUMBERJO', 1759, 92, 6),
	(23149, 'SURUHWADANG', 1759, 92, 6),
	(23150, 'TRIWUNG KIDUL', 1759, 92, 6),
	(23151, 'TRIWUNG LOR', 1759, 92, 6),
	(23152, 'CURAHGRINTING', 1760, 92, 6),
	(23153, 'KANIGARAN', 1760, 92, 6),
	(23154, 'KEBONSARI KULON', 1760, 92, 6),
	(23155, 'KEBONSARI WETAN', 1760, 92, 6),
	(23156, 'SUKOHARJO', 1760, 92, 6),
	(23157, 'TISNONEGARAN', 1760, 92, 6),
	(23158, 'JREBENG KULON', 1761, 92, 6),
	(23159, 'JREBENG LOR', 1761, 92, 6),
	(23160, 'JREBENG WETAN', 1761, 92, 6),
	(23161, 'KARENG LOR', 1761, 92, 6),
	(23162, 'KEDOPOK', 1761, 92, 6),
	(23163, 'SUMBER WETAN', 1761, 92, 6),
	(23164, 'CURAHTEMU', 1762, 92, 6),
	(23165, 'KEDUNGREJOSO', 1762, 92, 6),
	(23166, 'KOTA ANYAR', 1762, 92, 6),
	(23167, 'PASEMBON', 1762, 92, 6),
	(23168, 'SAMBIRAMPAK KIDUL', 1762, 92, 6),
	(23169, 'SAMBIRAMPAK LOR', 1762, 92, 6),
	(23170, 'SIDOMULYO', 1762, 92, 6),
	(23171, 'SIDOREJO', 1762, 92, 6),
	(23172, 'SUKOREJO', 1762, 92, 6),
	(23173, 'SUMBERCENTANG (SUMBERCENTENG)', 1762, 92, 6),
	(23174, 'TALKADANG', 1762, 92, 6),
	(23175, 'TAMBAKUKIR', 1762, 92, 6),
	(23176, 'TRIWUNGAN', 1762, 92, 6),
	(23177, 'ALASSUMUR KULON', 1763, 92, 6),
	(23178, 'ASEMBAGUS', 1763, 92, 6),
	(23179, 'ASEMBAKOR', 1763, 92, 6),
	(23180, 'BULU', 1763, 92, 6),
	(23181, 'KALIBUNTU', 1763, 92, 6),
	(23182, 'KANDANGJATI KULON', 1763, 92, 6),
	(23183, 'KANDANGJATI WETAN', 1763, 92, 6),
	(23184, 'KEBONAGUNG', 1763, 92, 6),
	(23185, 'KRAKSAAN WETAN', 1763, 92, 6),
	(23186, 'KREGENAN (KEREGENAN)', 1763, 92, 6),
	(23187, 'PATOKAN', 1763, 92, 6),
	(23188, 'RANGANG (RANGKANG)', 1763, 92, 6),
	(23189, 'RONDOKUNING', 1763, 92, 6),
	(23190, 'SEMAMPIR', 1763, 92, 6),
	(23191, 'SIDOMUKTI', 1763, 92, 6),
	(23192, 'SIDOPEKSO', 1763, 92, 6),
	(23193, 'SUMBERLELE', 1763, 92, 6),
	(23194, 'TAMAN SARI', 1763, 92, 6),
	(23195, 'DAWUAN', 1764, 92, 6),
	(23196, 'GEBANGAN', 1764, 92, 6),
	(23197, 'JATIURIP', 1764, 92, 6),
	(23198, 'KAMALKUNING', 1764, 92, 6),
	(23199, 'KARANGREN', 1764, 92, 6),
	(23200, 'KEDUNGCALUK', 1764, 92, 6),
	(23201, 'KREJENGAN', 1764, 92, 6),
	(23202, 'OPO OPO', 1764, 92, 6),
	(23203, 'PATEMON', 1764, 92, 6),
	(23204, 'RAWAN', 1764, 92, 6),
	(23205, 'SEBORO', 1764, 92, 6),
	(23206, 'SENTONG', 1764, 92, 6),
	(23207, 'SOKAAN', 1764, 92, 6),
	(23208, 'SUMBERKATIMOHO', 1764, 92, 6),
	(23209, 'TANJUNGSARI', 1764, 92, 6),
	(23210, 'TEMENGGUNGAN', 1764, 92, 6),
	(23211, 'WIDORO', 1764, 92, 6),
	(23212, 'BERMI (BREMI)', 1765, 92, 6),
	(23213, 'BETEK', 1765, 92, 6),
	(23214, 'GUYANGAN', 1765, 92, 6),
	(23215, 'KALIANAN', 1765, 92, 6),
	(23216, 'KERTOSUKO', 1765, 92, 6),
	(23217, 'KROBONGAN (KROBUNGAN)', 1765, 92, 6),
	(23218, 'KRUCIL', 1765, 92, 6),
	(23219, 'PANDANLARAS', 1765, 92, 6),
	(23220, 'PLAOSAN', 1765, 92, 6),
	(23221, 'ROTO', 1765, 92, 6),
	(23222, 'SENENG', 1765, 92, 6),
	(23223, 'SUMBER DUREN', 1765, 92, 6),
	(23224, 'TAMBELANG', 1765, 92, 6),
	(23225, 'WATUPANJANG', 1765, 92, 6),
	(23226, 'CLARAK', 1766, 92, 6),
	(23227, 'JORONGAN', 1766, 92, 6),
	(23228, 'KEPRANGAN (KERPANGAN)', 1766, 92, 6),
	(23229, 'LECES', 1766, 92, 6),
	(23230, 'MALASAN KULON', 1766, 92, 6),
	(23231, 'PONDOK WULUH', 1766, 92, 6),
	(23232, 'SUMBER KEDAWUNG', 1766, 92, 6),
	(23233, 'TIGASAN KULON', 1766, 92, 6),
	(23234, 'TIGASAN WETAN', 1766, 92, 6),
	(23235, 'WARUJINGGO', 1766, 92, 6),
	(23236, 'BANJARIMBO', 1767, 92, 6),
	(23237, 'BOTO', 1767, 92, 6),
	(23238, 'BRANGGAH', 1767, 92, 6),
	(23239, 'BULUKANDANG', 1767, 92, 6),
	(23240, 'CUKURGULING', 1767, 92, 6),
	(23241, 'KARANG ASEM', 1767, 92, 6),
	(23242, 'KARANG JATI', 1767, 92, 6),
	(23243, 'KRONTO', 1767, 92, 6),
	(23244, 'LAMBANGKUNING', 1767, 92, 6),
	(23245, 'LUMBANG', 1767, 92, 6),
	(23246, 'LUMBANG', 1767, 92, 6),
	(23247, 'NEGOROREJO', 1767, 92, 6),
	(23248, 'PALANGBESI', 1767, 92, 6),
	(23249, 'PANCUR', 1767, 92, 6),
	(23250, 'PANDITAN', 1767, 92, 6),
	(23251, 'PURUT', 1767, 92, 6),
	(23252, 'SAPIH', 1767, 92, 6),
	(23253, 'TANDONGSENTUL (TANDON SENTUL)', 1767, 92, 6),
	(23254, 'WATULUMBUNG', 1767, 92, 6),
	(23255, 'WELULANG', 1767, 92, 6),
	(23256, 'WONOGORO', 1767, 92, 6),
	(23257, 'WONOREJO', 1767, 92, 6),
	(23258, 'BRABE', 1768, 92, 6),
	(23259, 'BRANI KULON', 1768, 92, 6),
	(23260, 'BRANI WETAN', 1768, 92, 6),
	(23261, 'BRUMBUNGAN KIDUL', 1768, 92, 6),
	(23262, 'GANTING KULON', 1768, 92, 6),
	(23263, 'GANTING WETAN', 1768, 92, 6),
	(23264, 'GERONGAN', 1768, 92, 6),
	(23265, 'KEDUNGSARI', 1768, 92, 6),
	(23266, 'MARON KIDUL', 1768, 92, 6),
	(23267, 'MARON KULON', 1768, 92, 6),
	(23268, 'MARON WETAN', 1768, 92, 6),
	(23269, 'PEGALANGAN KIDUL', 1768, 92, 6),
	(23270, 'PUSPAN', 1768, 92, 6),
	(23271, 'SATREYAN', 1768, 92, 6),
	(23272, 'SUKO', 1768, 92, 6),
	(23273, 'SUMBERDAWE', 1768, 92, 6),
	(23274, 'SUMBERPOH', 1768, 92, 6),
	(23275, 'WONOREJO', 1768, 92, 6),
	(23276, 'JATI', 1769, 92, 6),
	(23277, 'MANGUNHARJO', 1769, 92, 6),
	(23278, 'MAYANGAN', 1769, 92, 6),
	(23279, 'SUKABUMI', 1769, 92, 6),
	(23280, 'WIROBORANG', 1769, 92, 6),
	(23281, 'ALAS TENGAH', 1770, 92, 6),
	(23282, 'BHINAR', 1770, 92, 6),
	(23283, 'JABUNG CANDI', 1770, 92, 6),
	(23284, 'JABUNG SISIR', 1770, 92, 6),
	(23285, 'JABUNG WETAN', 1770, 92, 6),
	(23286, 'KALIKAJAR KULON', 1770, 92, 6),
	(23287, 'KALIKAJAR WETAN', 1770, 92, 6),
	(23288, 'KARANGANYAR', 1770, 92, 6),
	(23289, 'PAITON', 1770, 92, 6),
	(23290, 'PANDEAN', 1770, 92, 6),
	(23291, 'PATUNJUNGAN', 1770, 92, 6),
	(23292, 'PLAMPANG', 1770, 92, 6),
	(23293, 'PONDOK KELOR', 1770, 92, 6),
	(23294, 'RANDU MERAK', 1770, 92, 6),
	(23295, 'RANDU TATAH', 1770, 92, 6),
	(23296, 'SIDODADI', 1770, 92, 6),
	(23297, 'SUKODADI', 1770, 92, 6),
	(23298, 'SUMBERANYAR', 1770, 92, 6),
	(23299, 'SUMBEREJO', 1770, 92, 6),
	(23300, 'TAMAN', 1770, 92, 6),
	(23301, 'GEJUNGAN', 1771, 92, 6),
	(23302, 'KARANGBONG', 1771, 92, 6),
	(23303, 'KARANGGEGER', 1771, 92, 6),
	(23304, 'KARANGPRANTI', 1771, 92, 6),
	(23305, 'KETOMPEN', 1771, 92, 6),
	(23306, 'PAJARAKAN KULON', 1771, 92, 6),
	(23307, 'PENAMBANGAN', 1771, 92, 6),
	(23308, 'SELOGUDIG KULON', 1771, 92, 6),
	(23309, 'SELOGUDIG WETAN', 1771, 92, 6),
	(23310, 'SUKOKERTO', 1771, 92, 6),
	(23311, 'SUKOMULYO', 1771, 92, 6),
	(23312, 'TANJUNG', 1771, 92, 6),
	(23313, 'ALAS PANDAN', 1772, 92, 6),
	(23314, 'BIMO', 1772, 92, 6),
	(23315, 'BLIMBING', 1772, 92, 6),
	(23316, 'BUCOR KULON', 1772, 92, 6),
	(23317, 'BUCOR WETAN', 1772, 92, 6),
	(23318, 'GLAGAH', 1772, 92, 6),
	(23319, 'GONDOSULI', 1772, 92, 6),
	(23320, 'GUNGGUNGAN KIDUL', 1772, 92, 6),
	(23321, 'GUNGGUNGAN LOR', 1772, 92, 6),
	(23322, 'KALIDANDAN', 1772, 92, 6),
	(23323, 'KEDUNGSUMUR', 1772, 92, 6),
	(23324, 'KERTONEGORO', 1772, 92, 6),
	(23325, 'PAKUNIRAN', 1772, 92, 6),
	(23326, 'PATEMON KULON', 1772, 92, 6),
	(23327, 'RANON', 1772, 92, 6),
	(23328, 'SOGAAN', 1772, 92, 6),
	(23329, 'SUMBER KEMBAR', 1772, 92, 6),
	(23330, 'JETAK', 1773, 92, 6),
	(23331, 'KEDASIH', 1773, 92, 6),
	(23332, 'NGADAS', 1773, 92, 6),
	(23333, 'NGADIREJO', 1773, 92, 6),
	(23334, 'NGADISARI', 1773, 92, 6),
	(23335, 'NGEPUNG', 1773, 92, 6),
	(23336, 'PAKEL', 1773, 92, 6),
	(23337, 'SAPIKEREP', 1773, 92, 6),
	(23338, 'SARIWANI', 1773, 92, 6),
	(23339, 'SUKAPURA', 1773, 92, 6),
	(23340, 'WONOKERTO', 1773, 92, 6),
	(23341, 'WONOTORO', 1773, 92, 6),
	(23342, 'BABAKAN', 1774, 92, 6),
	(23343, 'BOGOREJO', 1774, 92, 6),
	(23344, 'CEPOKO', 1774, 92, 6),
	(23345, 'GEGUNUNG', 1774, 92, 6),
	(23346, 'GEMITO', 1774, 92, 6),
	(23347, 'GRAWAN', 1774, 92, 6),
	(23348, 'JADI', 1774, 92, 6),
	(23349, 'JATIHADI', 1774, 92, 6),
	(23350, 'KALIWADAS', 1774, 92, 6),
	(23351, 'KEDUNGASEM', 1774, 92, 6),
	(23352, 'KEDUNGTULUP', 1774, 92, 6),
	(23353, 'KEMANTREN', 1774, 92, 6),
	(23354, 'KENANGA', 1774, 92, 6),
	(23355, 'KRIKILAN', 1774, 92, 6),
	(23356, 'LEDOKOMBO', 1774, 92, 6),
	(23357, 'LOGEDE', 1774, 92, 6),
	(23358, 'LOGUNG', 1774, 92, 6),
	(23359, 'MATANGAJI', 1774, 92, 6),
	(23360, 'MEGULUNG', 1774, 92, 6),
	(23361, 'PANDANSARI', 1774, 92, 6),
	(23362, 'PASALAKAN', 1774, 92, 6),
	(23363, 'PEJAMBON', 1774, 92, 6),
	(23364, 'PELEMSARI', 1774, 92, 6),
	(23365, 'PERBUTULAN', 1774, 92, 6),
	(23366, 'POLBAYEM', 1774, 92, 6),
	(23367, 'RAMBAAN', 1774, 92, 6),
	(23368, 'RANDUAGUNG', 1774, 92, 6),
	(23369, 'RONGGOMULYO', 1774, 92, 6),
	(23370, 'SEKARSARI', 1774, 92, 6),
	(23371, 'SENDANG', 1774, 92, 6),
	(23372, 'SIDAWANGI', 1774, 92, 6),
	(23373, 'SUKOREJO', 1774, 92, 6),
	(23374, 'SUMBER', 1774, 92, 6),
	(23375, 'SUMBER', 1774, 92, 6),
	(23376, 'SUMBER', 1774, 92, 6),
	(23377, 'SUMBERANOM', 1774, 92, 6),
	(23378, 'TLOGOTUNGGAL', 1774, 92, 6),
	(23379, 'TUKMUDAL', 1774, 92, 6),
	(23380, 'TUKUL', 1774, 92, 6),
	(23381, 'WATUBELAH', 1774, 92, 6),
	(23382, 'WONOKERSO', 1774, 92, 6),
	(23383, 'AMBULU', 1775, 92, 6),
	(23384, 'BANJARSARI', 1775, 92, 6),
	(23385, 'GILI KETAPANG', 1775, 92, 6),
	(23386, 'JANGUR', 1775, 92, 6),
	(23387, 'LAWEYAN', 1775, 92, 6),
	(23388, 'LEMAHKEMBAR', 1775, 92, 6),
	(23389, 'MENTOR', 1775, 92, 6),
	(23390, 'MUNENG', 1775, 92, 6),
	(23391, 'MUNENG KIDUL', 1775, 92, 6),
	(23392, 'PESISIR', 1775, 92, 6),
	(23393, 'POHSANGIT LERES', 1775, 92, 6),
	(23394, 'SUMBERBENDO', 1775, 92, 6),
	(23395, 'SUMURMATI', 1775, 92, 6),
	(23396, 'BANJARSAWAH', 1776, 92, 6),
	(23397, 'BLADO KULON', 1776, 92, 6),
	(23398, 'BULUJARAN KIDUL', 1776, 92, 6),
	(23399, 'BULUJARAN LOR', 1776, 92, 6),
	(23400, 'GUNUNG BEKEL', 1776, 92, 6),
	(23401, 'MALASAN WETAN', 1776, 92, 6),
	(23402, 'PARAS', 1776, 92, 6),
	(23403, 'SUMBERBULU', 1776, 92, 6),
	(23404, 'SUMBERKLEDUNG', 1776, 92, 6),
	(23405, 'TEGAL MOJO', 1776, 92, 6),
	(23406, 'TEGAL SIWALAN', 1776, 92, 6),
	(23407, 'TEGAL SONO', 1776, 92, 6),
	(23408, 'ANDUNG BIRU', 1777, 92, 6),
	(23409, 'ANDUNG SARI', 1777, 92, 6),
	(23410, 'JANGKANG', 1777, 92, 6),
	(23411, 'PEGADANGAN', 1777, 92, 6),
	(23412, 'PESAWAHAN', 1777, 92, 6),
	(23413, 'RACEK', 1777, 92, 6),
	(23414, 'RANUAGUNG', 1777, 92, 6),
	(23415, 'RANUGEDANG', 1777, 92, 6),
	(23416, 'RENJING (REJING)', 1777, 92, 6),
	(23417, 'SEGARAN', 1777, 92, 6),
	(23418, 'TEGALWATU', 1777, 92, 6),
	(23419, 'TIRIS', 1777, 92, 6),
	(23420, 'TLOGOARGO', 1777, 92, 6),
	(23421, 'TLOGOSARI', 1777, 92, 6),
	(23422, 'TULUPARI', 1777, 92, 6),
	(23423, 'WEDUSAN', 1777, 92, 6),
	(23424, 'BAYEMAN', 1778, 92, 6),
	(23425, 'CURAH DRINGU', 1778, 92, 6),
	(23426, 'CURAH TULIS', 1778, 92, 6),
	(23427, 'DUNGUN', 1778, 92, 6),
	(23428, 'KLAMPOK', 1778, 92, 6),
	(23429, 'PAMATAN', 1778, 92, 6),
	(23430, 'SUMBEREJO', 1778, 92, 6),
	(23431, 'SUMBERKRAMAT', 1778, 92, 6),
	(23432, 'SUMENDI', 1778, 92, 6),
	(23433, 'TAMBAKREJO', 1778, 92, 6),
	(23434, 'TANJUNGREJO', 1778, 92, 6),
	(23435, 'TONGAS KULON', 1778, 92, 6),
	(23436, 'TONGAS WETAN', 1778, 92, 6),
	(23437, 'WRINGINANOM', 1778, 92, 6),
	(23438, 'JREBENG KIDUL', 1779, 92, 6),
	(23439, 'KEDUNGASEM', 1779, 92, 6),
	(23440, 'KEDUNGGALENG', 1779, 92, 6),
	(23441, 'PAKISTAJI', 1779, 92, 6),
	(23442, 'SUMBER TAMAN', 1779, 92, 6),
	(23443, 'WONOASIH', 1779, 92, 6),
	(23444, 'JREBENG', 1780, 92, 6),
	(23445, 'KARENG KIDUL', 1780, 92, 6),
	(23446, 'KEDUNG SUPIT', 1780, 92, 6),
	(23447, 'PATALAN', 1780, 92, 6),
	(23448, 'POHSANGIT LOR', 1780, 92, 6),
	(23449, 'POHSANGIT NGISOR', 1780, 92, 6),
	(23450, 'POHSANGIT TENGAH', 1780, 92, 6),
	(23451, 'SEPUH GEMPOL', 1780, 92, 6),
	(23452, 'SUMBERKARE', 1780, 92, 6),
	(23453, 'TUNGGAK CERME', 1780, 92, 6),
	(23454, 'WONOREJO', 1780, 92, 6),
	(23455, 'ASEM JARAN', 1781, 93, 6),
	(23456, 'BANYUATES', 1781, 93, 6),
	(23457, 'BATIOH', 1781, 93, 6),
	(23458, 'JATRA TIMUR', 1781, 93, 6),
	(23459, 'KEMBANG JERUK', 1781, 93, 6),
	(23460, 'LAR LAR', 1781, 93, 6),
	(23461, 'MASARAN', 1781, 93, 6),
	(23462, 'MONTOR', 1781, 93, 6),
	(23463, 'MORBATOH', 1781, 93, 6),
	(23464, 'NAGASAREH', 1781, 93, 6),
	(23465, 'NEPA', 1781, 93, 6),
	(23466, 'OLOR', 1781, 93, 6),
	(23467, 'PLANGGARAN BARAT', 1781, 93, 6),
	(23468, 'PLANGGARAN TIMUR', 1781, 93, 6),
	(23469, 'TAPAAN', 1781, 93, 6),
	(23470, 'TEBANAH', 1781, 93, 6),
	(23471, 'TEROSAN', 1781, 93, 6),
	(23472, 'TLAGAH', 1781, 93, 6),
	(23473, 'TOLANG', 1781, 93, 6),
	(23474, 'TRAPANG', 1781, 93, 6),
	(23475, 'ANGGERSEK', 1782, 93, 6),
	(23476, 'BANJAR TABULU', 1782, 93, 6),
	(23477, 'BANJAR TALELA', 1782, 93, 6),
	(23478, 'BATU KARANG', 1782, 93, 6),
	(23479, 'DHARMA CAMPLONG', 1782, 93, 6),
	(23480, 'DHARMA TANJUNG', 1782, 93, 6),
	(23481, 'MADUPAT', 1782, 93, 6),
	(23482, 'PAMOLAAN', 1782, 93, 6),
	(23483, 'PLAMPAAN', 1782, 93, 6),
	(23484, 'PRAJJAN', 1782, 93, 6),
	(23485, 'RABASAN', 1782, 93, 6),
	(23486, 'SEJATI', 1782, 93, 6),
	(23487, 'TADDAN', 1782, 93, 6),
	(23488, 'TAMBAAN', 1782, 93, 6),
	(23489, 'ASEM NONGGAL', 1783, 93, 6),
	(23490, 'ASEM RAJA', 1783, 93, 6),
	(23491, 'BANCELOK', 1783, 93, 6),
	(23492, 'BUKER', 1783, 93, 6),
	(23493, 'JRENGIK', 1783, 93, 6),
	(23494, 'JUNGKARANG', 1783, 93, 6),
	(23495, 'KALANGAN PRAO', 1783, 93, 6),
	(23496, 'KOTAH', 1783, 93, 6),
	(23497, 'MAJANGAN', 1783, 93, 6),
	(23498, 'MARGANTOKO', 1783, 93, 6),
	(23499, 'MLAKA', 1783, 93, 6),
	(23500, 'PANYEPEN', 1783, 93, 6),
	(23501, 'PLAKARAN', 1783, 93, 6),
	(23502, 'TAMAN', 1783, 93, 6),
	(23503, 'BLUURAN (BLU URAN)', 1784, 93, 6),
	(23504, 'BULMATET', 1784, 93, 6),
	(23505, 'GUNUNG KESAN', 1784, 93, 6),
	(23506, 'KARANG PENANG OLOH', 1784, 93, 6),
	(23507, 'KARANG PENANG ONJUR', 1784, 93, 6),
	(23508, 'POREH', 1784, 93, 6),
	(23509, 'TLAMBAH', 1784, 93, 6),
	(23510, 'BAJRASOKAH', 1785, 93, 6),
	(23511, 'BANJAR', 1785, 93, 6),
	(23512, 'BANYUKAPAH', 1785, 93, 6),
	(23513, 'BATOPORO BARAT', 1785, 93, 6),
	(23514, 'BATOPORO TIMUR', 1785, 93, 6),
	(23515, 'DALEMAN', 1785, 93, 6),
	(23516, 'GUNUNG ELEH', 1785, 93, 6),
	(23517, 'KEDUNGDUNG', 1785, 93, 6),
	(23518, 'KOMIS', 1785, 93, 6),
	(23519, 'KRAMAT', 1785, 93, 6),
	(23520, 'MUKTESAREH (MOKTESAREH)', 1785, 93, 6),
	(23521, 'NYELOH', 1785, 93, 6),
	(23522, 'OMBUL', 1785, 93, 6),
	(23523, 'PAJERUAN', 1785, 93, 6),
	(23524, 'PALENGGIYAN (PALENGGIAN)', 1785, 93, 6),
	(23525, 'PASARENAN', 1785, 93, 6),
	(23526, 'RABASAN', 1785, 93, 6),
	(23527, 'ROHAYU', 1785, 93, 6),
	(23528, 'BANGUN REJO', 1786, 93, 6),
	(23529, 'BANYUSOKAH', 1786, 93, 6),
	(23530, 'BERUNDUNG', 1786, 93, 6),
	(23531, 'BIRA BARAT', 1786, 93, 6),
	(23532, 'BUNTEN BARAT', 1786, 93, 6),
	(23533, 'BUNTEN TIMUR', 1786, 93, 6),
	(23534, 'KARANG ANYAR', 1786, 93, 6),
	(23535, 'KARANG SARI', 1786, 93, 6),
	(23536, 'KEMUKUS', 1786, 93, 6),
	(23537, 'KETAPANG', 1786, 93, 6),
	(23538, 'KETAPANG BARAT', 1786, 93, 6),
	(23539, 'KETAPANG DAYA', 1786, 93, 6),
	(23540, 'KETAPANG LAOK', 1786, 93, 6),
	(23541, 'KETAPANG TIMUR', 1786, 93, 6),
	(23542, 'LEGUNDI', 1786, 93, 6),
	(23543, 'LEMBUNG NALA', 1786, 93, 6),
	(23544, 'PANCOR', 1786, 93, 6),
	(23545, 'PANGEREMAN', 1786, 93, 6),
	(23546, 'PAOPALE DAYA', 1786, 93, 6),
	(23547, 'PAOPALELAOK', 1786, 93, 6),
	(23548, 'PEMATANG PASIR', 1786, 93, 6),
	(23549, 'RABIYAN', 1786, 93, 6),
	(23550, 'RUGUK', 1786, 93, 6),
	(23551, 'SIDO LUHUR', 1786, 93, 6),
	(23552, 'SIDOASIH', 1786, 93, 6),
	(23553, 'SRI PENDOWO', 1786, 93, 6),
	(23554, 'SUMBERNADI', 1786, 93, 6),
	(23555, 'SUMUR', 1786, 93, 6),
	(23556, 'TAMAN SARI', 1786, 93, 6),
	(23557, 'TRIDARMAYOGA', 1786, 93, 6),
	(23558, 'WAI SIDOMUKTI', 1786, 93, 6),
	(23559, 'ANGSOKAH', 1787, 93, 6),
	(23560, 'ASTAPAH', 1787, 93, 6),
	(23561, 'GERSEMPAL', 1787, 93, 6),
	(23562, 'JRANGOAN (JRANGUAN)', 1787, 93, 6),
	(23563, 'KAMONDUNG', 1787, 93, 6),
	(23564, 'KARANG GAYAM', 1787, 93, 6),
	(23565, 'KARANG NANGGER', 1787, 93, 6),
	(23566, 'KEBUN SAREH', 1787, 93, 6),
	(23567, 'MADULANG', 1787, 93, 6),
	(23568, 'METENG', 1787, 93, 6),
	(23569, 'NAPO DAYA (NAPA DAYA)', 1787, 93, 6),
	(23570, 'NAPOLAOK (NAPA LAOK)', 1787, 93, 6),
	(23571, 'OMBEN', 1787, 93, 6),
	(23572, 'PANDAN', 1787, 93, 6),
	(23573, 'RAPA DAYA', 1787, 93, 6),
	(23574, 'RAPA LAOK', 1787, 93, 6),
	(23575, 'RONGDALEM (RONGDALAM)', 1787, 93, 6),
	(23576, 'SOGIYAN (SOGIAN)', 1787, 93, 6),
	(23577, 'TAMBAK', 1787, 93, 6),
	(23578, 'TEMORAN', 1787, 93, 6),
	(23579, 'APAAN', 1788, 93, 6),
	(23580, 'GULBUNG', 1788, 93, 6),
	(23581, 'PACANGGAAN', 1788, 93, 6),
	(23582, 'PANGARENGAN', 1788, 93, 6),
	(23583, 'PANYERANGAN (PANYIRANGAN)', 1788, 93, 6),
	(23584, 'RAGUNG', 1788, 93, 6),
	(23585, 'BAPELLE', 1789, 93, 6),
	(23586, 'GUNUNG RANCAK', 1789, 93, 6),
	(23587, 'JELGUNG', 1789, 93, 6),
	(23588, 'LEPELLE', 1789, 93, 6),
	(23589, 'PANDIYANGAN', 1789, 93, 6),
	(23590, 'ROBATAL', 1789, 93, 6),
	(23591, 'SAWAH TENGAH', 1789, 93, 6),
	(23592, 'TORJUNAN', 1789, 93, 6),
	(23593, 'TRAGIH', 1789, 93, 6),
	(23594, 'AENG SAREH', 1790, 93, 6),
	(23595, 'BANYUANYAR', 1790, 93, 6),
	(23596, 'BANYUMAS', 1790, 93, 6),
	(23597, 'BARUH', 1790, 93, 6),
	(23598, 'BRANI', 1790, 93, 6),
	(23599, 'DALPENANG', 1790, 93, 6),
	(23600, 'GUNUNG MADDAH', 1790, 93, 6),
	(23601, 'GUNUNG SEKAR', 1790, 93, 6),
	(23602, 'KAMONING', 1790, 93, 6),
	(23603, 'KARANG DALEM', 1790, 93, 6),
	(23604, 'KARANGASEM', 1790, 93, 6),
	(23605, 'KARANGJATI', 1790, 93, 6),
	(23606, 'KARANGTENGAH', 1790, 93, 6),
	(23607, 'KETANGGUNG', 1790, 93, 6),
	(23608, 'NUSAJATI', 1790, 93, 6),
	(23609, 'PABERASAN', 1790, 93, 6),
	(23610, 'PAKETINGAN', 1790, 93, 6),
	(23611, 'PANGELEN', 1790, 93, 6),
	(23612, 'PANGGUNG', 1790, 93, 6),
	(23613, 'PASEYAN', 1790, 93, 6),
	(23614, 'PEKALONGAN (PAKALONGAN)', 1790, 93, 6),
	(23615, 'POLAGAN', 1790, 93, 6),
	(23616, 'PULAU MANDANGIN (MANDINGAN)', 1790, 93, 6),
	(23617, 'RONG TENGAH', 1790, 93, 6),
	(23618, 'SAMPANG', 1790, 93, 6),
	(23619, 'SIDASARI', 1790, 93, 6),
	(23620, 'TAMAN SAREH', 1790, 93, 6),
	(23621, 'TANGGUMONG', 1790, 93, 6),
	(23622, 'BIRA TENGAH', 1791, 93, 6),
	(23623, 'BIRA TIMUR', 1791, 93, 6),
	(23624, 'SOKOBANAH DAYA', 1791, 93, 6),
	(23625, 'SOKOBANAH LAOK', 1791, 93, 6),
	(23626, 'SOKOBANAH TENGAH', 1791, 93, 6),
	(23627, 'TAMBERU BARAT', 1791, 93, 6),
	(23628, 'TAMBERU DAYA', 1791, 93, 6),
	(23629, 'TAMBERU LAOK', 1791, 93, 6),
	(23630, 'TAMBERU TIMUR', 1791, 93, 6),
	(23631, 'TOBAI BARAT', 1791, 93, 6),
	(23632, 'TOBAI TENGAH', 1791, 93, 6),
	(23633, 'TOBAI TIMUR', 1791, 93, 6),
	(23634, 'BANGSAH', 1792, 93, 6),
	(23635, 'BUNDAH', 1792, 93, 6),
	(23636, 'DISANAH', 1792, 93, 6),
	(23637, 'JUNUK (JUNOK)', 1792, 93, 6),
	(23638, 'KLOBUR', 1792, 93, 6),
	(23639, 'LABANG', 1792, 93, 6),
	(23640, 'LABUHAN', 1792, 93, 6),
	(23641, 'MARPARAN', 1792, 93, 6),
	(23642, 'NOREH', 1792, 93, 6),
	(23643, 'PLASAH', 1792, 93, 6),
	(23644, 'SRESEH', 1792, 93, 6),
	(23645, 'TAMAN', 1792, 93, 6),
	(23646, 'BANJAR BILLAH', 1793, 93, 6),
	(23647, 'BARUNG GAGAH', 1793, 93, 6),
	(23648, 'BATORASANG (BATURASANG)', 1793, 93, 6),
	(23649, 'BERINGIN (BRINGIN)', 1793, 93, 6),
	(23650, 'BIREM', 1793, 93, 6),
	(23651, 'KARANG ANYAR', 1793, 93, 6),
	(23652, 'MAMBULU BARAT', 1793, 93, 6),
	(23653, 'SAMARAN', 1793, 93, 6),
	(23654, 'SOMBER', 1793, 93, 6),
	(23655, 'TAMBELANGAN', 1793, 93, 6),
	(23656, 'BRINGIN NONGGAL', 1794, 93, 6),
	(23657, 'DULANG', 1794, 93, 6),
	(23658, 'JERUK POROT', 1794, 93, 6),
	(23659, 'KANJAR', 1794, 93, 6),
	(23660, 'KARA', 1794, 93, 6),
	(23661, 'KODAK', 1794, 93, 6),
	(23662, 'KRAMPON', 1794, 93, 6),
	(23663, 'PANGONGSEAN', 1794, 93, 6),
	(23664, 'PATAPAN', 1794, 93, 6),
	(23665, 'PATARONGAN', 1794, 93, 6),
	(23666, 'TANAH MERAH', 1794, 93, 6),
	(23667, 'TORJUN', 1794, 93, 6),
	(23668, 'BAKALAN WRINGINPITU', 1795, 94, 6),
	(23669, 'BAKUNG PRINGGODANI', 1795, 94, 6),
	(23670, 'BAKUNG TEMENGGUNGAN', 1795, 94, 6),
	(23671, 'BALONGBENDO', 1795, 94, 6),
	(23672, 'BOGEM PINGGIR', 1795, 94, 6),
	(23673, 'GAGANG KEPUHSARI', 1795, 94, 6),
	(23674, 'JABARAN', 1795, 94, 6),
	(23675, 'JERUK LEGI', 1795, 94, 6),
	(23676, 'KEDUNG SUKODANI', 1795, 94, 6),
	(23677, 'KEMANGSEN', 1795, 94, 6),
	(23678, 'PENAMBANGAN', 1795, 94, 6),
	(23679, 'SEDURI', 1795, 94, 6),
	(23680, 'SEKETI', 1795, 94, 6),
	(23681, 'SINGKALAN', 1795, 94, 6),
	(23682, 'SUMOKEMBANGSRI', 1795, 94, 6),
	(23683, 'SUWALUH', 1795, 94, 6),
	(23684, 'WARUBERON', 1795, 94, 6),
	(23685, 'WATESARI', 1795, 94, 6),
	(23686, 'WONOKARANG', 1795, 94, 6),
	(23687, 'WONOKUPANG', 1795, 94, 6),
	(23688, 'BANJARKEMANTREN', 1796, 94, 6),
	(23689, 'BANJARSARI', 1796, 94, 6),
	(23690, 'BUDURAN', 1796, 94, 6),
	(23691, 'DAMARSI', 1796, 94, 6),
	(23692, 'DUKUHTENGAH', 1796, 94, 6),
	(23693, 'ENTALSEWU', 1796, 94, 6),
	(23694, 'PAGERWOJO', 1796, 94, 6),
	(23695, 'PRASUNG', 1796, 94, 6),
	(23696, 'SAWOHAN', 1796, 94, 6),
	(23697, 'SIDOKEPUNG', 1796, 94, 6),
	(23698, 'SIDOKERTO', 1796, 94, 6),
	(23699, 'SIDOMULYO', 1796, 94, 6),
	(23700, 'SIWALAN PANJI', 1796, 94, 6),
	(23701, 'SUKOREJO', 1796, 94, 6),
	(23702, 'WADUNGASIH', 1796, 94, 6),
	(23703, 'BALONGDOWO', 1797, 94, 6),
	(23704, 'BALONGGABUS', 1797, 94, 6),
	(23705, 'BLIGO', 1797, 94, 6),
	(23706, 'CANDI', 1797, 94, 6),
	(23707, 'DURUNGBANJAR', 1797, 94, 6),
	(23708, 'DURUNGBEDUG', 1797, 94, 6),
	(23709, 'GELAM', 1797, 94, 6),
	(23710, 'JAMBANGAN', 1797, 94, 6),
	(23711, 'KALIPECABEAN', 1797, 94, 6),
	(23712, 'KARANGTANJUNG', 1797, 94, 6),
	(23713, 'KEBUNSARI', 1797, 94, 6),
	(23714, 'KEDUNG PELUK', 1797, 94, 6),
	(23715, 'KEDUNGKENDO', 1797, 94, 6),
	(23716, 'KENDALPECABEAN', 1797, 94, 6),
	(23717, 'KLURAK', 1797, 94, 6),
	(23718, 'LARANGAN', 1797, 94, 6),
	(23719, 'NGAMPELSARI', 1797, 94, 6),
	(23720, 'SEPANDE', 1797, 94, 6),
	(23721, 'SIDODADI', 1797, 94, 6),
	(23722, 'SUGIH WARAS', 1797, 94, 6),
	(23723, 'SUMOKALI', 1797, 94, 6),
	(23724, 'SUMORAME', 1797, 94, 6),
	(23725, 'TENGGULUNAN', 1797, 94, 6),
	(23726, 'WEDORO KLURAK', 1797, 94, 6),
	(23727, 'BALONGTANI', 1798, 94, 6),
	(23728, 'BESUKI', 1798, 94, 6),
	(23729, 'DUKUHSARI', 1798, 94, 6),
	(23730, 'JEMIRAHAN', 1798, 94, 6),
	(23731, 'KEBOGUYANG', 1798, 94, 6),
	(23732, 'KEDUNGCANGKRING', 1798, 94, 6),
	(23733, 'KEDUNGPANDAN', 1798, 94, 6),
	(23734, 'KEDUNGREJO', 1798, 94, 6),
	(23735, 'KUPANG', 1798, 94, 6),
	(23736, 'PANGGREH', 1798, 94, 6),
	(23737, 'PEJARAKAN', 1798, 94, 6),
	(23738, 'PERMISAN', 1798, 94, 6),
	(23739, 'SEMAMBUNG', 1798, 94, 6),
	(23740, 'TAMBAK KALISOGO', 1798, 94, 6),
	(23741, 'TROMPOASRI', 1798, 94, 6),
	(23742, 'BALONG GARUT', 1799, 94, 6),
	(23743, 'CANGKRING', 1799, 94, 6),
	(23744, 'GADING', 1799, 94, 6),
	(23745, 'JENGGOT', 1799, 94, 6),
	(23746, 'KANDANGAN', 1799, 94, 6),
	(23747, 'KEDUNGRAWAN', 1799, 94, 6),
	(23748, 'KEDUNGSUMUR', 1799, 94, 6),
	(23749, 'KEPER', 1799, 94, 6),
	(23750, 'KERET', 1799, 94, 6),
	(23751, 'KREMBUNG', 1799, 94, 6),
	(23752, 'LEMUJUT', 1799, 94, 6),
	(23753, 'MOJORUNTUT', 1799, 94, 6),
	(23754, 'PLOSO', 1799, 94, 6),
	(23755, 'REJENI', 1799, 94, 6),
	(23756, 'TAMBAKREJO', 1799, 94, 6),
	(23757, 'TANJEG WAGIR', 1799, 94, 6),
	(23758, 'WANGKAL', 1799, 94, 6),
	(23759, 'WAUNG', 1799, 94, 6),
	(23760, 'WONOMLATI', 1799, 94, 6),
	(23761, 'BARENGKRAJAN', 1800, 94, 6),
	(23762, 'GAMPING', 1800, 94, 6),
	(23763, 'JATIKALANG', 1800, 94, 6),
	(23764, 'JERUK GAMPING', 1800, 94, 6),
	(23765, 'JUNWANGI', 1800, 94, 6),
	(23766, 'KATRUNGAN (KATERUNGAN)', 1800, 94, 6),
	(23767, 'KEBOHARAN', 1800, 94, 6),
	(23768, 'KEMASAN', 1800, 94, 6),
	(23769, 'KRATON', 1800, 94, 6),
	(23770, 'KRIAN', 1800, 94, 6),
	(23771, 'PONOKAWAN', 1800, 94, 6),
	(23772, 'SEDENGAN MIJEN', 1800, 94, 6),
	(23773, 'SIDOMOJO', 1800, 94, 6),
	(23774, 'SIDOMULYO', 1800, 94, 6),
	(23775, 'SIDOREJO', 1800, 94, 6),
	(23776, 'TAMBAK KEMERAKAN', 1800, 94, 6),
	(23777, 'TEMPEL', 1800, 94, 6),
	(23778, 'TERIK', 1800, 94, 6),
	(23779, 'TERUNG KULON', 1800, 94, 6),
	(23780, 'TERUNG WETAN', 1800, 94, 6),
	(23781, 'TROPODO', 1800, 94, 6),
	(23782, 'WATUGOLONG', 1800, 94, 6),
	(23783, 'CANDIPARI', 1801, 94, 6),
	(23784, 'GEDANG', 1801, 94, 6),
	(23785, 'GLAGAH ARUM', 1801, 94, 6),
	(23786, 'JATIREJO', 1801, 94, 6),
	(23787, 'JUWET KENONGO', 1801, 94, 6),
	(23788, 'KEBAKALAN', 1801, 94, 6),
	(23789, 'KEBONAGUNG', 1801, 94, 6),
	(23790, 'KEDUNGBOTO', 1801, 94, 6),
	(23791, 'KEDUNGSOLO', 1801, 94, 6),
	(23792, 'KESAMBI', 1801, 94, 6),
	(23793, 'LAJUK', 1801, 94, 6),
	(23794, 'MINDI', 1801, 94, 6),
	(23795, 'PAMOTAN', 1801, 94, 6),
	(23796, 'PESAWAHAN', 1801, 94, 6),
	(23797, 'PLUMBON', 1801, 94, 6),
	(23798, 'PORONG', 1801, 94, 6),
	(23799, 'RENOKENONGO', 1801, 94, 6),
	(23800, 'SIRING', 1801, 94, 6),
	(23801, 'WUNUT', 1801, 94, 6),
	(23802, 'BANJAR KEMUNING', 1802, 94, 6),
	(23803, 'BETRO', 1802, 94, 6),
	(23804, 'BUNCITAN', 1802, 94, 6),
	(23805, 'CEMANDI', 1802, 94, 6),
	(23806, 'GISIK CEMANDI', 1802, 94, 6),
	(23807, 'KALANGANYAR', 1802, 94, 6),
	(23808, 'KWANGSAN', 1802, 94, 6),
	(23809, 'PABEAN', 1802, 94, 6),
	(23810, 'PEPE', 1802, 94, 6),
	(23811, 'PRANTI', 1802, 94, 6),
	(23812, 'PULUNGAN', 1802, 94, 6),
	(23813, 'SEDATI AGUNG', 1802, 94, 6),
	(23814, 'SEDATI GEDE', 1802, 94, 6),
	(23815, 'SEGORO TAMBAK', 1802, 94, 6),
	(23816, 'SEMAMPIR', 1802, 94, 6),
	(23817, 'TAMBAK CEMANDI', 1802, 94, 6),
	(23818, 'BANJARBENDO', 1803, 94, 6),
	(23819, 'BLURU KIDUL', 1803, 94, 6),
	(23820, 'BULUSIDOKARE', 1803, 94, 6),
	(23821, 'CELEP', 1803, 94, 6),
	(23822, 'CEMENG BAKALAN', 1803, 94, 6),
	(23823, 'CEMENG KALANG', 1803, 94, 6),
	(23824, 'GEBANG', 1803, 94, 6),
	(23825, 'JATI', 1803, 94, 6),
	(23826, 'KEMIRI', 1803, 94, 6),
	(23827, 'LEBO', 1803, 94, 6),
	(23828, 'LEMAHPUTRO', 1803, 94, 6),
	(23829, 'MAGERSARI', 1803, 94, 6),
	(23830, 'PEKAUMAN', 1803, 94, 6),
	(23831, 'PUCANG', 1803, 94, 6),
	(23832, 'PUCANGANOM', 1803, 94, 6),
	(23833, 'RANGKAHKIDUL', 1803, 94, 6),
	(23834, 'SARI ROGO', 1803, 94, 6),
	(23835, 'SEKARDANGAN', 1803, 94, 6),
	(23836, 'SIDOKARE', 1803, 94, 6),
	(23837, 'SIDOKLUMPUK', 1803, 94, 6),
	(23838, 'SIDOKUMPUL', 1803, 94, 6),
	(23839, 'SUKO', 1803, 94, 6),
	(23840, 'SUMPUT', 1803, 94, 6),
	(23841, 'URANGAGUNG (JEDONG)', 1803, 94, 6),
	(23842, 'BANJAR ASRI', 1804, 94, 6),
	(23843, 'BANJAR PANJI', 1804, 94, 6),
	(23844, 'BORO', 1804, 94, 6),
	(23845, 'GANGGANG PANJANG', 1804, 94, 6),
	(23846, 'GEMPOL SARI', 1804, 94, 6),
	(23847, 'KALIDAWIR', 1804, 94, 6),
	(23848, 'KALISAMPURNO', 1804, 94, 6),
	(23849, 'KALITENGAH', 1804, 94, 6),
	(23850, 'KEDENSARI', 1804, 94, 6),
	(23851, 'KEDUNG BANTENG', 1804, 94, 6),
	(23852, 'KEDUNG BENDO', 1804, 94, 6),
	(23853, 'KETAPANG', 1804, 94, 6),
	(23854, 'KETEGAN', 1804, 94, 6),
	(23855, 'KLUDAN', 1804, 94, 6),
	(23856, 'NGABAN', 1804, 94, 6),
	(23857, 'PENATARSEWU', 1804, 94, 6),
	(23858, 'PUTAT', 1804, 94, 6),
	(23859, 'RANDEGAN', 1804, 94, 6),
	(23860, 'SENTUL', 1804, 94, 6),
	(23861, 'BALONGMACEKAN', 1805, 94, 6),
	(23862, 'BANJARWUNGU', 1805, 94, 6),
	(23863, 'GAMPINGROWO', 1805, 94, 6),
	(23864, 'GEMPOLKLUTUK', 1805, 94, 6),
	(23865, 'JANTI', 1805, 94, 6),
	(23866, 'KALIMATI', 1805, 94, 6),
	(23867, 'KEDINDING', 1805, 94, 6),
	(23868, 'KEDUNGBOCOK', 1805, 94, 6),
	(23869, 'KEMUNING', 1805, 94, 6),
	(23870, 'KENDALSEWU', 1805, 94, 6),
	(23871, 'KLANTINGSARI', 1805, 94, 6),
	(23872, 'KRAMAT TEMENGGUNG', 1805, 94, 6),
	(23873, 'MERGOBENER', 1805, 94, 6),
	(23874, 'MERGOSARI', 1805, 94, 6),
	(23875, 'MINDUGADING', 1805, 94, 6),
	(23876, 'MLIRIPROWO', 1805, 94, 6),
	(23877, 'SEBANI', 1805, 94, 6),
	(23878, 'SEGODOBANCANG', 1805, 94, 6),
	(23879, 'SINGOGALIH', 1805, 94, 6),
	(23880, 'TARIK', 1805, 94, 6),
	(23881, 'GELANG', 1806, 94, 6),
	(23882, 'GRABAGAN', 1806, 94, 6),
	(23883, 'GRINTING', 1806, 94, 6),
	(23884, 'GROGOL', 1806, 94, 6),
	(23885, 'JANTI', 1806, 94, 6),
	(23886, 'JIKEN', 1806, 94, 6),
	(23887, 'KAJEKSAN', 1806, 94, 6),
	(23888, 'KEBARON', 1806, 94, 6),
	(23889, 'KEDONDONG', 1806, 94, 6),
	(23890, 'KEMANTREN', 1806, 94, 6),
	(23891, 'KENONGO', 1806, 94, 6),
	(23892, 'KEPADANGAN', 1806, 94, 6),
	(23893, 'KEPATIHAN', 1806, 94, 6),
	(23894, 'KEPUH KEMIRI', 1806, 94, 6),
	(23895, 'KEPUNTEN', 1806, 94, 6),
	(23896, 'MEDALEM', 1806, 94, 6),
	(23897, 'MODONG', 1806, 94, 6),
	(23898, 'PANGKEMIRI', 1806, 94, 6),
	(23899, 'SINGOPADU', 1806, 94, 6),
	(23900, 'SUDIMORO', 1806, 94, 6),
	(23901, 'TLASIH', 1806, 94, 6),
	(23902, 'TULANGAN', 1806, 94, 6),
	(23903, 'BECIRONGENGOR', 1807, 94, 6),
	(23904, 'CANDINEGORO', 1807, 94, 6),
	(23905, 'JIMBARAN KULON', 1807, 94, 6),
	(23906, 'JIMBARAN WETAN', 1807, 94, 6),
	(23907, 'KARANGPURI', 1807, 94, 6),
	(23908, 'KETIMANG', 1807, 94, 6),
	(23909, 'LAMBANGAN', 1807, 94, 6),
	(23910, 'MOJORANGAGUNG', 1807, 94, 6),
	(23911, 'MULYODADI', 1807, 94, 6),
	(23912, 'PAGERNGUMBUK', 1807, 94, 6),
	(23913, 'PILANG', 1807, 94, 6),
	(23914, 'PLAOSAN', 1807, 94, 6),
	(23915, 'PLOSO', 1807, 94, 6),
	(23916, 'POPOH', 1807, 94, 6),
	(23917, 'SAWOCANGKRING', 1807, 94, 6),
	(23918, 'SEMAMBUNG', 1807, 94, 6),
	(23919, 'SIMO ANGIN ANGIN', 1807, 94, 6),
	(23920, 'SIMOKETAWANG', 1807, 94, 6),
	(23921, 'SUMBEREJO', 1807, 94, 6),
	(23922, 'TANGGUL', 1807, 94, 6),
	(23923, 'WONOAYU', 1807, 94, 6),
	(23924, 'WONOKALANG', 1807, 94, 6),
	(23925, 'WONOKASIAN', 1807, 94, 6),
	(23926, 'ASEMBAGUS', 1808, 95, 6),
	(23927, 'AWAR-AWAR', 1808, 95, 6),
	(23928, 'BANTAL', 1808, 95, 6),
	(23929, 'GUDANG', 1808, 95, 6),
	(23930, 'KEDONGLO', 1808, 95, 6),
	(23931, 'KERTOSARI', 1808, 95, 6),
	(23932, 'MOJOSARI', 1808, 95, 6),
	(23933, 'PERANTE', 1808, 95, 6),
	(23934, 'TRIGONCO', 1808, 95, 6),
	(23935, 'WRINGIN ANOM', 1808, 95, 6),
	(23936, 'BANYUGLUGUR', 1809, 95, 6),
	(23937, 'KALIANGET', 1809, 95, 6),
	(23938, 'KALISARI', 1809, 95, 6),
	(23939, 'LUBAWANG', 1809, 95, 6),
	(23940, 'SELOBANTENG', 1809, 95, 6),
	(23941, 'TELEMPONG', 1809, 95, 6),
	(23942, 'TEPOS', 1809, 95, 6),
	(23943, 'BANARAN', 1810, 95, 6),
	(23944, 'BANYU PUTIH', 1810, 95, 6),
	(23945, 'BANYUPUTIH', 1810, 95, 6),
	(23946, 'BULU', 1810, 95, 6),
	(23947, 'DLIMAS', 1810, 95, 6),
	(23948, 'KALIBALIK', 1810, 95, 6),
	(23949, 'KEDAWUNG', 1810, 95, 6),
	(23950, 'LUWUNG', 1810, 95, 6),
	(23951, 'PENUNDAN', 1810, 95, 6),
	(23952, 'SEMBUNG', 1810, 95, 6),
	(23953, 'SUMBERANYAR', 1810, 95, 6),
	(23954, 'SUMBEREJO', 1810, 95, 6),
	(23955, 'SUMBERWARU', 1810, 95, 6),
	(23956, 'TIMBANG', 1810, 95, 6),
	(23957, 'WONOREJO', 1810, 95, 6),
	(23958, 'BESOLE', 1811, 95, 6),
	(23959, 'BESUKI', 1811, 95, 6),
	(23960, 'BESUKI', 1811, 95, 6),
	(23961, 'BLIMBING', 1811, 95, 6),
	(23962, 'BLORO', 1811, 95, 6),
	(23963, 'DEMUNG', 1811, 95, 6),
	(23964, 'JETIS', 1811, 95, 6),
	(23965, 'KALIMAS', 1811, 95, 6),
	(23966, 'KEBOIRENG', 1811, 95, 6),
	(23967, 'LANGKAP', 1811, 95, 6),
	(23968, 'PESISIR', 1811, 95, 6),
	(23969, 'SEDAYUGUNUNG', 1811, 95, 6),
	(23970, 'SIYOTOBAGUS', 1811, 95, 6),
	(23971, 'SUMBEREJO', 1811, 95, 6),
	(23972, 'TANGGULKUNDUNG', 1811, 95, 6),
	(23973, 'TANGGULTURUS', 1811, 95, 6),
	(23974, 'TANGGULWELAHAN', 1811, 95, 6),
	(23975, 'TULUNGREJO', 1811, 95, 6),
	(23976, 'WATESKROYO', 1811, 95, 6),
	(23977, 'WIDORO PAYUNG', 1811, 95, 6),
	(23978, 'BLETOK', 1812, 95, 6),
	(23979, 'BUNGATAN', 1812, 95, 6),
	(23980, 'MLANDINGAN WETAN', 1812, 95, 6),
	(23981, 'PASIR PUTIH', 1812, 95, 6),
	(23982, 'PATEMON', 1812, 95, 6),
	(23983, 'SELOWOGO', 1812, 95, 6),
	(23984, 'SUMBERTENGAH', 1812, 95, 6),
	(23985, 'AGEL', 1813, 95, 6),
	(23986, 'CURAH KALAK', 1813, 95, 6),
	(23987, 'GADINGAN', 1813, 95, 6),
	(23988, 'JANGKAR', 1813, 95, 6),
	(23989, 'KUMBANGSARI', 1813, 95, 6),
	(23990, 'PALANGAN', 1813, 95, 6),
	(23991, 'PESANGGRAHAN', 1813, 95, 6),
	(23992, 'SOPET', 1813, 95, 6),
	(23993, 'CURAH SURI', 1814, 95, 6),
	(23994, 'JATIBANTENG', 1814, 95, 6),
	(23995, 'KEMBANGSARI', 1814, 95, 6),
	(23996, 'PATEGALAN', 1814, 95, 6),
	(23997, 'PATEMON', 1814, 95, 6),
	(23998, 'SEMAMBUNG', 1814, 95, 6),
	(23999, 'SUMBERANYAR', 1814, 95, 6),
	(24000, 'WRINGINANOM', 1814, 95, 6)
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert villige 23")
}