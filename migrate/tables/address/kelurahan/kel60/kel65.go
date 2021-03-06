package kel60

import "github.com/danangkonang/ceodeaja-go/config"

func Kel65() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
	(65001, 'PEWUNU', 5081, 364, 26),
	(65002, 'RARAMPADENDE', 5081, 364, 26),
	(65003, 'SIBONU', 5081, 364, 26),
	(65004, 'BALONGGA', 5082, 364, 26),
	(65005, 'BALUASE', 5082, 364, 26),
	(65006, 'BANGGA', 5082, 364, 26),
	(65007, 'BULUBETE', 5082, 364, 26),
	(65008, 'JONO', 5082, 364, 26),
	(65009, 'POI', 5082, 364, 26),
	(65010, 'PULU', 5082, 364, 26),
	(65011, 'ROGO', 5082, 364, 26),
	(65012, 'SAMBO', 5082, 364, 26),
	(65013, 'WALATANA', 5082, 364, 26),
	(65014, 'WISOLO', 5082, 364, 26),
	(65015, 'KALAWARA', 5083, 364, 26),
	(65016, 'OMU', 5083, 364, 26),
	(65017, 'PAKULI', 5083, 364, 26),
	(65018, 'PANDERE', 5083, 364, 26),
	(65019, 'SIMORO', 5083, 364, 26),
	(65020, 'TUVA (TUWA)', 5083, 364, 26),
	(65021, 'BALANE', 5084, 364, 26),
	(65022, 'BOLOBIA', 5084, 364, 26),
	(65023, 'DAENGGUNE', 5084, 364, 26),
	(65024, 'DODA', 5084, 364, 26),
	(65025, 'KALORA', 5084, 364, 26),
	(65026, 'KANUNA', 5084, 364, 26),
	(65027, 'PORAME', 5084, 364, 26),
	(65028, 'RONDINGO', 5084, 364, 26),
	(65029, 'UWEMANJE', 5084, 364, 26),
	(65030, 'BANGGAIBA', 5085, 364, 26),
	(65031, 'BOLADANGKO', 5085, 364, 26),
	(65032, 'BOLAPAPU', 5085, 364, 26),
	(65033, 'LONCA', 5085, 364, 26),
	(65034, 'MATAUE', 5085, 364, 26),
	(65035, 'NAMO', 5085, 364, 26),
	(65036, 'RANTEWULU', 5085, 364, 26),
	(65037, 'SALUA', 5085, 364, 26),
	(65038, 'SIWONGI', 5085, 364, 26),
	(65039, 'SUNGKU', 5085, 364, 26),
	(65040, 'TANGKULOWI', 5085, 364, 26),
	(65041, 'TORO', 5085, 364, 26),
	(65042, 'TOWULU', 5085, 364, 26),
	(65043, 'WINATU', 5085, 364, 26),
	(65044, 'GIMPU', 5086, 364, 26),
	(65045, 'LAWUA', 5086, 364, 26),
	(65046, 'LEMPELERO', 5086, 364, 26),
	(65047, 'MOA', 5086, 364, 26),
	(65048, 'O\\O', 5086, 364, 26),
	(65049, 'PALAMAKI', 5086, 364, 26),
	(65050, 'PALIMAKIJAWA (PILIMAKUJAWA/MAKUJAWA)', 5086, 364, 26),
	(65051, 'SALUTOME', 5086, 364, 26),
	(65052, 'TOMPI BUGIS', 5086, 364, 26),
	(65053, 'TOMUA', 5086, 364, 26),
	(65054, 'WANGKA', 5086, 364, 26),
	(65055, 'WATUKILO', 5086, 364, 26),
	(65056, 'ANCA', 5087, 364, 26),
	(65057, 'LANGKO', 5087, 364, 26),
	(65058, 'PURO/PUROO', 5087, 364, 26),
	(65059, 'TOMADO', 5087, 364, 26),
	(65060, 'BALIASE', 5088, 364, 26),
	(65061, 'BEKA', 5088, 364, 26),
	(65062, 'BINANGGA', 5088, 364, 26),
	(65063, 'BOMBA', 5088, 364, 26),
	(65064, 'BOYABALIASE', 5088, 364, 26),
	(65065, 'LEBANU', 5088, 364, 26),
	(65066, 'PADENDE', 5088, 364, 26),
	(65067, 'SIBEDI', 5088, 364, 26),
	(65068, 'SUNJU', 5088, 364, 26),
	(65069, 'TINGGEDE', 5088, 364, 26),
	(65070, 'TINGGEDE SELATAN', 5088, 364, 26),
	(65071, 'DOMBU', 5089, 364, 26),
	(65072, 'LEWARA', 5089, 364, 26),
	(65073, 'MATANTIMALI', 5089, 364, 26),
	(65074, 'ONGULERO', 5089, 364, 26),
	(65075, 'PANASIBAJA', 5089, 364, 26),
	(65076, 'SOI', 5089, 364, 26),
	(65077, 'TAIPANGGABE', 5089, 364, 26),
	(65078, 'WAYU', 5089, 364, 26),
	(65079, 'WIAPORE', 5089, 364, 26),
	(65080, 'WUGAGA', 5089, 364, 26),
	(65081, 'BULILI', 5090, 364, 26),
	(65082, 'KADIDIA', 5090, 364, 26),
	(65083, 'KAMARORA A', 5090, 364, 26),
	(65084, 'KAMARORA B', 5090, 364, 26),
	(65085, 'SOPU', 5090, 364, 26),
	(65086, 'AMPERA', 5091, 364, 26),
	(65087, 'BAHAGIA', 5091, 364, 26),
	(65088, 'BAKU-BAKULU', 5091, 364, 26),
	(65089, 'BERDIKARI', 5091, 364, 26),
	(65090, 'BOBO', 5091, 364, 26),
	(65091, 'BUNGA', 5091, 364, 26),
	(65092, 'KAPIROE', 5091, 364, 26),
	(65093, 'LEMBAN TONGOA', 5091, 364, 26),
	(65094, 'MAKMUR', 5091, 364, 26),
	(65095, 'PETIMBE', 5091, 364, 26),
	(65096, 'RAHMAT', 5091, 364, 26),
	(65097, 'RANTELEDE (RANTELEDA)', 5091, 364, 26),
	(65098, 'REJEKI', 5091, 364, 26),
	(65099, 'SEJAHTERA', 5091, 364, 26),
	(65100, 'SIGIMPU', 5091, 364, 26),
	(65101, 'SINTUWU', 5091, 364, 26),
	(65102, 'TANAH HARAPAN', 5091, 364, 26),
	(65103, 'TONGOA', 5091, 364, 26),
	(65104, 'UENUNI', 5091, 364, 26),
	(65105, 'BANASU', 5092, 364, 26),
	(65106, 'KALAMANTA', 5092, 364, 26),
	(65107, 'KANTEWU', 5092, 364, 26),
	(65108, 'KANTEWU II', 5092, 364, 26),
	(65109, 'KOJA', 5092, 364, 26),
	(65110, 'LAWE', 5092, 364, 26),
	(65111, 'LONE BASA', 5092, 364, 26),
	(65112, 'MAMU', 5092, 364, 26),
	(65113, 'MAPAHI', 5092, 364, 26),
	(65114, 'MORUI (MURUI)', 5092, 364, 26),
	(65115, 'ONU', 5092, 364, 26),
	(65116, 'PARELEA', 5092, 364, 26),
	(65117, 'PEANA', 5092, 364, 26),
	(65118, 'BORA', 5093, 364, 26),
	(65119, 'JONO OGE', 5093, 364, 26),
	(65120, 'KALUKUBULA', 5093, 364, 26),
	(65121, 'LOLU', 5093, 364, 26),
	(65122, 'LORU', 5093, 364, 26),
	(65123, 'MARANATHA', 5093, 364, 26),
	(65124, 'MPANAU', 5093, 364, 26),
	(65125, 'NGATABARU', 5093, 364, 26),
	(65126, 'OLOBUJU', 5093, 364, 26),
	(65127, 'POMBEWE', 5093, 364, 26),
	(65128, 'SIDERA', 5093, 364, 26),
	(65129, 'SIDONDO I', 5093, 364, 26),
	(65130, 'SIDONDO II', 5093, 364, 26),
	(65131, 'SIDONDO III', 5093, 364, 26),
	(65132, 'SIDONDO IV', 5093, 364, 26),
	(65133, 'SOULOWE', 5093, 364, 26),
	(65134, 'WATUNONJU', 5093, 364, 26),
	(65135, 'LAMBARA', 5094, 364, 26),
	(65136, 'SIBALAYA SELATAN', 5094, 364, 26),
	(65137, 'SIBALAYA UTARA', 5094, 364, 26),
	(65138, 'SIBOWI', 5094, 364, 26),
	(65139, 'AMPANA', 5095, 365, 26),
	(65140, 'BAILO', 5095, 365, 26),
	(65141, 'BAILO BARU', 5095, 365, 26),
	(65142, 'BONE RATO', 5095, 365, 26),
	(65143, 'BUNTONGI', 5095, 365, 26),
	(65144, 'DONDO', 5095, 365, 26),
	(65145, 'DONDO BARAT', 5095, 365, 26),
	(65146, 'LABIABAE', 5095, 365, 26),
	(65147, 'LABUAN', 5095, 365, 26),
	(65148, 'MALOTONG', 5095, 365, 26),
	(65149, 'MUARA TOBA', 5095, 365, 26),
	(65150, 'PADANG TUMBUO', 5095, 365, 26),
	(65151, 'PATINGKO', 5095, 365, 26),
	(65152, 'SABULIRA TOBA', 5095, 365, 26),
	(65153, 'SALUABA', 5095, 365, 26),
	(65154, 'SANSARINO', 5095, 365, 26),
	(65155, 'SUMOLI', 5095, 365, 26),
	(65156, 'UAMALINGKU', 5095, 365, 26),
	(65157, 'UENTANAGA ATAS', 5095, 365, 26),
	(65158, 'UENTANAGA BAWAH', 5095, 365, 26),
	(65159, 'BALANGGALA', 5096, 365, 26),
	(65160, 'BALINGARA', 5096, 365, 26),
	(65161, 'BANTUGA', 5096, 365, 26),
	(65162, 'BORONE', 5096, 365, 26),
	(65163, 'BULAN JAYA', 5096, 365, 26),
	(65164, 'GIRI MULYO', 5096, 365, 26),
	(65165, 'KAJULANGKO', 5096, 365, 26),
	(65166, 'LONGGE', 5096, 365, 26),
	(65167, 'MANTANGISI', 5096, 365, 26),
	(65168, 'MPOA', 5096, 365, 26),
	(65169, 'PUSUNGI', 5096, 365, 26),
	(65170, 'SABO', 5096, 365, 26),
	(65171, 'SUKAMAJU', 5096, 365, 26),
	(65172, 'TAMPABATU', 5096, 365, 26),
	(65173, 'TETE ATAS', 5096, 365, 26),
	(65174, 'TETE BAWAH', 5096, 365, 26),
	(65175, 'UEBONE', 5096, 365, 26),
	(65176, 'UEMAKUNI', 5096, 365, 26),
	(65177, 'URUNDAKA', 5096, 365, 26),
	(65178, 'WANASARI', 5096, 365, 26),
	(65179, 'AWO', 5097, 365, 26),
	(65180, 'BANGKAGI', 5097, 365, 26),
	(65181, 'BAULU', 5097, 365, 26),
	(65182, 'BENTENG', 5097, 365, 26),
	(65183, 'BUNGAYO', 5097, 365, 26),
	(65184, 'KATUPAT', 5097, 365, 26),
	(65185, 'KOLOLIO', 5097, 365, 26),
	(65186, 'LEBITI', 5097, 365, 26),
	(65187, 'LEMBANATO', 5097, 365, 26),
	(65188, 'MATOBIAI (MOTOBISI)', 5097, 365, 26),
	(65189, 'PULAU ENAM', 5097, 365, 26),
	(65190, 'SAMPOBAE', 5097, 365, 26),
	(65191, 'TOBIL', 5097, 365, 26),
	(65192, 'TONGKABO', 5097, 365, 26),
	(65193, 'URULEPE', 5097, 365, 26),
	(65194, 'BAHARI', 5098, 365, 26),
	(65195, 'BANANO', 5098, 365, 26),
	(65196, 'BETAUA', 5098, 365, 26),
	(65197, 'DATARAN BUGI', 5098, 365, 26),
	(65198, 'KALEMBA I', 5098, 365, 26),
	(65199, 'KALEMBA II', 5098, 365, 26),
	(65200, 'KORONDODA', 5098, 365, 26),
	(65201, 'LEMORO', 5098, 365, 26),
	(65202, 'PANCUMA', 5098, 365, 26),
	(65203, 'PODI', 5098, 365, 26),
	(65204, 'SANDADA', 5098, 365, 26),
	(65205, 'TAYAWA', 5098, 365, 26),
	(65206, 'TOJO', 5098, 365, 26),
	(65207, 'TONGKU', 5098, 365, 26),
	(65208, 'UEDELE', 5098, 365, 26),
	(65209, 'UEKULI', 5098, 365, 26),
	(65210, 'BAMBALO', 5099, 365, 26),
	(65211, 'GALUGA', 5099, 365, 26),
	(65212, 'KABALO', 5099, 365, 26),
	(65213, 'MALEI TOJO', 5099, 365, 26),
	(65214, 'MALEWA', 5099, 365, 26),
	(65215, 'MATAKO', 5099, 365, 26),
	(65216, 'MAWOMBA', 5099, 365, 26),
	(65217, 'NGGAWIA', 5099, 365, 26),
	(65218, 'TANAMAWAU', 5099, 365, 26),
	(65219, 'TATARI', 5099, 365, 26),
	(65220, 'TOLIBA', 5099, 365, 26),
	(65221, 'TOMBIANO', 5099, 365, 26),
	(65222, 'UJUNG TIBU', 5099, 365, 26),
	(65223, 'BONEBAE I', 5100, 365, 26),
	(65224, 'BONEBAE II', 5100, 365, 26),
	(65225, 'BONEVETO (BONEVOTO)', 5100, 365, 26),
	(65226, 'BONGKA KOY (KOI)', 5100, 365, 26),
	(65227, 'BONGKA MAKMUR', 5100, 365, 26),
	(65228, 'BORNEANG', 5100, 365, 26),
	(65229, 'CEMPA', 5100, 365, 26),
	(65230, 'KASIALA', 5100, 365, 26),
	(65231, 'MAROWO', 5100, 365, 26),
	(65232, 'MIRE', 5100, 365, 26),
	(65233, 'PARANONGE', 5100, 365, 26),
	(65234, 'ROMPI', 5100, 365, 26),
	(65235, 'TAKIBANGKE', 5100, 365, 26),
	(65236, 'TAMPANOMBO', 5100, 365, 26),
	(65237, 'TOBAMAWU', 5100, 365, 26),
	(65238, 'UEKAMBUNO', 5100, 365, 26),
	(65239, 'UEMATOPA', 5100, 365, 26),
	(65240, 'WATUSONGU', 5100, 365, 26),
	(65241, 'BAMBU', 5101, 365, 26),
	(65242, 'BINANGUNA', 5101, 365, 26),
	(65243, 'BOMBA', 5101, 365, 26),
	(65244, 'CENDANA', 5101, 365, 26),
	(65245, 'KAMBUTU', 5101, 365, 26),
	(65246, 'KAVETAN', 5101, 365, 26),
	(65247, 'KULINGKINARI', 5101, 365, 26),
	(65248, 'LEMBANYA', 5101, 365, 26),
	(65249, 'LINDO', 5101, 365, 26),
	(65250, 'LUANGON', 5101, 365, 26),
	(65251, 'MALINO', 5101, 365, 26),
	(65252, 'MOLOWAGU', 5101, 365, 26),
	(65253, 'SIATU', 5101, 365, 26),
	(65254, 'TANIMPO', 5101, 365, 26),
	(65255, 'TANINGKOLA', 5101, 365, 26),
	(65256, 'TANJUNG PUDE', 5101, 365, 26),
	(65257, 'TUMBULAWA', 5101, 365, 26),
	(65258, 'UNA UNA', 5101, 365, 26),
	(65259, 'WAKAI', 5101, 365, 26),
	(65260, 'BIGA', 5102, 365, 26),
	(65261, 'KATOGOP (KOTOGOP)', 5102, 365, 26),
	(65262, 'KONDONGAN', 5102, 365, 26),
	(65263, 'MALAPO', 5102, 365, 26),
	(65264, 'PASOKAN', 5102, 365, 26),
	(65265, 'SALINGGOHA', 5102, 365, 26),
	(65266, 'TINGKI', 5102, 365, 26),
	(65267, 'TONGIDON', 5102, 365, 26),
	(65268, 'DOLONG A (DOLONGA)', 5103, 365, 26),
	(65269, 'DOLONG B', 5103, 365, 26),
	(65270, 'KABALUTAN', 5103, 365, 26),
	(65271, 'KADODA', 5103, 365, 26),
	(65272, 'KALIA', 5103, 365, 26),
	(65273, 'KOLAMI (KOLAMPI)', 5103, 365, 26),
	(65274, 'LOE', 5103, 365, 26),
	(65275, 'LUOK', 5103, 365, 26),
	(65276, 'MALENGE', 5103, 365, 26),
	(65277, 'OLILAN', 5103, 365, 26),
	(65278, 'PAUTU', 5103, 365, 26),
	(65279, 'POPOLII (POPOLII)', 5103, 365, 26),
	(65280, 'TIGA PULAU', 5103, 365, 26),
	(65281, 'TUMOTOK', 5103, 365, 26),
	(65282, 'TUTUNG', 5103, 365, 26),
	(65283, 'BARU', 5104, 366, 26),
	(65284, 'BUNTUNA', 5104, 366, 26),
	(65285, 'DADAKITAN', 5104, 366, 26),
	(65286, 'LELEAN NONO', 5104, 366, 26),
	(65287, 'NALU', 5104, 366, 26),
	(65288, 'PANASAKAN', 5104, 366, 26),
	(65289, 'SIDOARJO', 5104, 366, 26),
	(65290, 'TAMBUN', 5104, 366, 26),
	(65291, 'TUWELEY', 5104, 366, 26),
	(65292, 'GALANDAU', 5105, 366, 26),
	(65293, 'KAYU LOMPA', 5105, 366, 26),
	(65294, 'KINAPASAN', 5105, 366, 26),
	(65295, 'KONGKOMOS', 5105, 366, 26),
	(65296, 'LABONU', 5105, 366, 26),
	(65297, 'OGOSIPAT', 5105, 366, 26),
	(65298, 'SIBALUTON', 5105, 366, 26),
	(65299, 'SILONDOU', 5105, 366, 26),
	(65300, 'DUNGINGIS', 5106, 366, 26),
	(65301, 'GALUMPANG', 5106, 366, 26),
	(65302, 'KAPAS', 5106, 366, 26),
	(65303, 'LINGADAN', 5106, 366, 26),
	(65304, 'ABBAJARENG', 5107, 366, 26),
	(65305, 'BANGKIR', 5107, 366, 26),
	(65306, 'DONGKO', 5107, 366, 26),
	(65307, 'KOMBO', 5107, 366, 26),
	(65308, 'LEMPE', 5107, 366, 26),
	(65309, 'MIMBALA', 5107, 366, 26),
	(65310, 'PADDUMPU', 5107, 366, 26),
	(65311, 'SONI', 5107, 366, 26),
	(65312, 'TAMPIALA', 5107, 366, 26),
	(65313, 'BAMBAPULA', 5108, 366, 26),
	(65314, 'BANAGAN', 5108, 366, 26),
	(65315, 'KABINUANG', 5108, 366, 26),
	(65316, 'MALAMBIGU', 5108, 366, 26),
	(65317, 'OGOTUA', 5108, 366, 26),
	(65318, 'SESE', 5108, 366, 26),
	(65319, 'SIMATANG TANJUNG', 5108, 366, 26),
	(65320, 'SIMATANG UTARA', 5108, 366, 26),
	(65321, 'TOMPOH', 5108, 366, 26),
	(65322, 'ANGGASAN', 5109, 366, 26),
	(65323, 'BAMBAPUN', 5109, 366, 26),
	(65324, 'LAIS', 5109, 366, 26),
	(65325, 'LUOK MANIPI', 5109, 366, 26),
	(65326, 'MALALA', 5109, 366, 26),
	(65327, 'MALOMBA', 5109, 366, 26),
	(65328, 'MALULU', 5109, 366, 26),
	(65329, 'OGOGASANG', 5109, 366, 26),
	(65330, 'OGOGILI', 5109, 366, 26),
	(65331, 'OGOWELE', 5109, 366, 26),
	(65332, 'SALUMBIA', 5109, 366, 26),
	(65333, 'TINABOGAN', 5109, 366, 26),
	(65334, 'AIR RAJA', 5110, 366, 26),
	(65335, 'BAJUGAN', 5110, 366, 26),
	(65336, 'BANDAR KUALA', 5110, 366, 26),
	(65337, 'BARU TITI BESI', 5110, 366, 26),
	(65338, 'BATU LOKONG', 5110, 366, 26),
	(65339, 'GALANG BARAT', 5110, 366, 26),
	(65340, 'GALANG BARU (PULAU)', 5110, 366, 26),
	(65341, 'GALANG KOTA', 5110, 366, 26),
	(65342, 'GALANG SUKA', 5110, 366, 26),
	(65343, 'GINUNGGUNG', 5110, 366, 26),
	(65344, 'JAHARUM A', 5110, 366, 26),
	(65345, 'JAHARUM B', 5110, 366, 26),
	(65346, 'JUHAR BARU', 5110, 366, 26),
	(65347, 'KALANGKANGAN', 5110, 366, 26),
	(65348, 'KARAS (PULAU KARAS)', 5110, 366, 26),
	(65349, 'KELAPA SATU', 5110, 366, 26),
	(65350, 'KERAMAT GAJAH', 5110, 366, 26),
	(65351, 'KOTANGAN', 5110, 366, 26),
	(65352, 'KOTASAN', 5110, 366, 26),
	(65353, 'LAKATAN', 5110, 366, 26),
	(65354, 'LALOS', 5110, 366, 26),
	(65355, 'LANTAPAN', 5110, 366, 26),
	(65356, 'NOGOREJO', 5110, 366, 26),
	(65357, 'OGOMULI (OGOMOLI)', 5110, 366, 26),
	(65358, 'PAKU', 5110, 366, 26),
	(65359, 'PAYA ITIK', 5110, 366, 26),
	(65360, 'PAYA KUDA', 5110, 366, 26),
	(65361, 'PAYA SAMPIR', 5110, 366, 26),
	(65362, 'PETANGGUHAN', 5110, 366, 26),
	(65363, 'PETUMBUKAN', 5110, 366, 26),
	(65364, 'PISANG PALA', 5110, 366, 26),
	(65365, 'PULAU ABANG', 5110, 366, 26),
	(65366, 'PULAU TAGOR BATU', 5110, 366, 26),
	(65367, 'REMPANG CATE', 5110, 366, 26),
	(65368, 'SABANG', 5110, 366, 26),
	(65369, 'SANDANA', 5110, 366, 26),
	(65370, 'SEI KARANG', 5110, 366, 26),
	(65371, 'SEI PUTIH', 5110, 366, 26),
	(65372, 'SEMBULANG (PULAU REMPANG)', 5110, 366, 26),
	(65373, 'SIJANTUNG', 5110, 366, 26),
	(65374, 'SUBANG MAS', 5110, 366, 26),
	(65375, 'TANAH ABANG', 5110, 366, 26),
	(65376, 'TANAH MERAH', 5110, 366, 26),
	(65377, 'TANJUNG GUSTI', 5110, 366, 26),
	(65378, 'TANJUNG SIPORKIS', 5110, 366, 26),
	(65379, 'TENDE', 5110, 366, 26),
	(65380, 'TIMBANG DELI', 5110, 366, 26),
	(65381, 'TINIGI', 5110, 366, 26),
	(65382, 'JANJA', 5111, 366, 26),
	(65383, 'LAMPASIO', 5111, 366, 26),
	(65384, 'MAIBUA', 5111, 366, 26),
	(65385, 'MULYASARI (MULIA SARI)', 5111, 366, 26),
	(65386, 'OGOMATANANG', 5111, 366, 26),
	(65387, 'OYOM', 5111, 366, 26),
	(65388, 'SALUGAN', 5111, 366, 26),
	(65389, 'SIBEA', 5111, 366, 26),
	(65390, 'TINADING', 5111, 366, 26),
	(65391, 'BAMBALAGA', 5112, 366, 26),
	(65392, 'BATUILO', 5112, 366, 26),
	(65393, 'BILO', 5112, 366, 26),
	(65394, 'BUGA', 5112, 366, 26),
	(65395, 'KABETAN', 5112, 366, 26),
	(65396, 'KAMALU', 5112, 366, 26),
	(65397, 'LABUAN LOBO', 5112, 366, 26),
	(65398, 'MUARA BESAR', 5112, 366, 26),
	(65399, 'PAGAITAN', 5112, 366, 26),
	(65400, 'PULIAS', 5112, 366, 26),
	(65401, 'SAMBUJAN', 5112, 366, 26),
	(65402, 'BINONTOAN', 5113, 366, 26),
	(65403, 'DIULE', 5113, 366, 26),
	(65404, 'LAKUAN TOLI TOLI', 5113, 366, 26),
	(65405, 'LAULALANG', 5113, 366, 26),
	(65406, 'PINJAN', 5113, 366, 26),
	(65407, 'SALUMPAGA', 5113, 366, 26),
	(65408, 'SANTIGI', 5113, 366, 26),
	(65409, 'TELUK JAYA', 5113, 366, 26),
	(65410, 'TIMBOLO', 5113, 366, 26),




	
	(65411, 'BOLIHUTUO', 5114, 367, 27),
	(65412, 'BOTUMOITO', 5114, 367, 27),
	(65413, 'DULANGEA (DULANGEYA)', 5114, 367, 27),
	(65414, 'HUTAMONU', 5114, 367, 27),
	(65415, 'PATOAMEME', 5114, 367, 27),
	(65416, 'POTANGA', 5114, 367, 27),
	(65417, 'RUMBIA', 5114, 367, 27),
	(65418, 'TAPADAA', 5114, 367, 27),
	(65419, 'TUTULO', 5114, 367, 27),
	(65420, 'DULUPI', 5115, 367, 27),
	(65421, 'KOTARAJA', 5115, 367, 27),
	(65422, 'PANGI', 5115, 367, 27),
	(65423, 'POLOHUNGO', 5115, 367, 27),
	(65424, 'TABONGO', 5115, 367, 27),
	(65425, 'TANAH PUTIH', 5115, 367, 27),
	(65426, 'TANGGA BARITO', 5115, 367, 27),
	(65427, 'TANGGA JAYA', 5115, 367, 27),
	(65428, 'BENDUNGAN', 5116, 367, 27),
	(65429, 'BUTI', 5116, 367, 27),
	(65430, 'KAARUYAN', 5116, 367, 27),
	(65431, 'KERAMAT', 5116, 367, 27),
	(65432, 'MANANGGU', 5116, 367, 27),
	(65433, 'PONTOLO', 5116, 367, 27),
	(65434, 'SALILAMA', 5116, 367, 27),
	(65435, 'TABULO', 5116, 367, 27),
	(65436, 'TABULO SELATAN', 5116, 367, 27),
	(65437, 'BALATE JAYA', 5117, 367, 27),
	(65438, 'BATU KRAMAT/KERAMAT', 5117, 367, 27),
	(65439, 'BONGO IV', 5117, 367, 27),
	(65440, 'BONGO NOL', 5117, 367, 27),
	(65441, 'BONGO TUA', 5117, 367, 27),
	(65442, 'BUALO', 5117, 367, 27),
	(65443, 'DILOATO', 5117, 367, 27),
	(65444, 'GIRISA', 5117, 367, 27),
	(65445, 'HULAWE (HULAWA)', 5117, 367, 27),
	(65446, 'HUWONGO', 5117, 367, 27),
	(65447, 'KARYA MURNI', 5117, 367, 27),
	(65448, 'KUALALUMPUR', 5117, 367, 27),
	(65449, 'MOLOMBULAHE', 5117, 367, 27),
	(65450, 'MUSTIKA', 5117, 367, 27),
	(65451, 'MUTIARA', 5117, 367, 27),
	(65452, 'PERMATA', 5117, 367, 27),
	(65453, 'REJONEGORO', 5117, 367, 27),
	(65454, 'SARIPI', 5117, 367, 27),
	(65455, 'SOSIAL', 5117, 367, 27),
	(65456, 'SUMBER JAYA **', 5117, 367, 27),
	(65457, 'TANGKOBU', 5117, 367, 27),
	(65458, 'TENILO', 5117, 367, 27),
	(65459, 'WONGGAHU', 5117, 367, 27),
	(65460, 'APITALAWU (APITALAWO)', 5118, 367, 27),
	(65461, 'BANGGA', 5118, 367, 27),
	(65462, 'BUBAA', 5118, 367, 27),
	(65463, 'BUKIT KARYA', 5118, 367, 27),
	(65464, 'LIMBATIHU', 5118, 367, 27),
	(65465, 'LITO', 5118, 367, 27),
	(65466, 'OLIBU', 5118, 367, 27),
	(65467, 'TOWAYU', 5118, 367, 27),
	(65468, 'AYUHULALO', 5119, 367, 27),
	(65469, 'BAJO', 5119, 367, 27),
	(65470, 'HUNGAYONAA', 5119, 367, 27),
	(65471, 'LAHUMBO', 5119, 367, 27),
	(65472, 'LAMU', 5119, 367, 27),
	(65473, 'LIMBATO', 5119, 367, 27),
	(65474, 'MODELOMO', 5119, 367, 27),
	(65475, 'MOHUNGO', 5119, 367, 27),
	(65476, 'PENTADU BARAT', 5119, 367, 27),
	(65477, 'PENTADU TIMUR', 5119, 367, 27),
	(65478, 'PILOLIYANGA', 5119, 367, 27),
	(65479, 'TENILO', 5119, 367, 27),
	(65480, 'BALEHARJO', 5120, 367, 27),
	(65481, 'BANGELAN', 5120, 367, 27),
	(65482, 'BENDOARUM', 5120, 367, 27),
	(65483, 'BENER', 5120, 367, 27),
	(65484, 'BENTANGAN', 5120, 367, 27),
	(65485, 'BOLALI', 5120, 367, 27),
	(65486, 'BONGO II', 5120, 367, 27),
	(65487, 'BONGO III', 5120, 367, 27),
	(65488, 'BOTO', 5120, 367, 27),
	(65489, 'BULAN', 5120, 367, 27),
	(65490, 'DIMITO', 5120, 367, 27),
	(65491, 'DULOHUPA', 5120, 367, 27),
	(65492, 'DUWET', 5120, 367, 27),
	(65493, 'DUWET', 5120, 367, 27),
	(65494, 'GARI', 5120, 367, 27),
	(65495, 'GUNTING', 5120, 367, 27),
	(65496, 'HARAPAN', 5120, 367, 27),
	(65497, 'JATI MULYA', 5120, 367, 27),
	(65498, 'JELOBO', 5120, 367, 27),
	(65499, 'JUMPONG', 5120, 367, 27),
	(65500, 'KAPURAN', 5120, 367, 27),
	(65501, 'KARANG REJEK', 5120, 367, 27),
	(65502, 'KARANG TENGAH', 5120, 367, 27),
	(65503, 'KEBOBANG', 5120, 367, 27),
	(65504, 'KEPEK', 5120, 367, 27),
	(65505, 'KINGKANG', 5120, 367, 27),
	(65506, 'KLUWUT', 5120, 367, 27),
	(65507, 'LOMBOK KULON', 5120, 367, 27),
	(65508, 'LOMBOK WETAN', 5120, 367, 27),
	(65509, 'LUMBUNG KEREP', 5120, 367, 27),
	(65510, 'MAKMUR **', 5120, 367, 27),
	(65511, 'MEKAR JAYA', 5120, 367, 27),
	(65512, 'MULO', 5120, 367, 27),
	(65513, 'NGREDEN', 5120, 367, 27),
	(65514, 'PANDANAN', 5120, 367, 27),
	(65515, 'PANGEYA (PANGEA)', 5120, 367, 27),
	(65516, 'PASAREJO', 5120, 367, 27),
	(65517, 'PELALANGAN', 5120, 367, 27),
	(65518, 'PIYAMAN', 5120, 367, 27),
	(65519, 'PLANDI', 5120, 367, 27),
	(65520, 'PLAOSAN', 5120, 367, 27),
	(65521, 'PULUTAN', 5120, 367, 27),
	(65522, 'RAHARJA', 5120, 367, 27),
	(65523, 'SARI TANI', 5120, 367, 27),
	(65524, 'SEJAHTERA **', 5120, 367, 27),
	(65525, 'SEKARAN', 5120, 367, 27),
	(65526, 'SELANG', 5120, 367, 27),
	(65527, 'SIDOWARNO', 5120, 367, 27),
	(65528, 'SIRAMAN', 5120, 367, 27),
	(65529, 'SUKA MAJU', 5120, 367, 27),
	(65530, 'SUKA MULIA (SUKAMULYA)', 5120, 367, 27),
	(65531, 'SUKOREJO', 5120, 367, 27),
	(65532, 'SUMBER TEMPUR', 5120, 367, 27),
	(65533, 'SUMBERDEM', 5120, 367, 27),
	(65534, 'SUMBERKALONG', 5120, 367, 27),
	(65535, 'TANGSIL WETAN', 5120, 367, 27),
	(65536, 'TANJUNG HARAPAN', 5120, 367, 27),
	(65537, 'TEGALGONDO', 5120, 367, 27),
	(65538, 'TELOYO', 5120, 367, 27),
	(65539, 'TRAKTAKAN', 5120, 367, 27),
	(65540, 'TRI RUKUN', 5120, 367, 27),
	(65541, 'TUMPENG', 5120, 367, 27),
	(65542, 'WADUNG GETAS', 5120, 367, 27),
	(65543, 'WARENG', 5120, 367, 27),
	(65544, 'WONOSARI', 5120, 367, 27),
	(65545, 'WONOSARI', 5120, 367, 27),
	(65546, 'WONOSARI', 5120, 367, 27),
	(65547, 'WUNUNG', 5120, 367, 27),
	(65548, 'BILOLANTUNA', 5121, 368, 27),
	(65549, 'CENDANA PUTIH', 5121, 368, 27),
	(65550, 'ILOHUUWA', 5121, 368, 27),
	(65551, 'INOGALUMA', 5121, 368, 27),
	(65552, 'MASIAGA', 5121, 368, 27),
	(65553, 'MOLAMAHU', 5121, 368, 27),
	(65554, 'MONANO', 5121, 368, 27),
	(65555, 'MOODULIO', 5121, 368, 27),
	(65556, 'MUARA BONE', 5121, 368, 27),
	(65557, 'PERMATA', 5121, 368, 27),
	(65558, 'SOGITIA', 5121, 368, 27),
	(65559, 'TALUDAA', 5121, 368, 27),
	(65560, 'TUMBUH MEKAR', 5121, 368, 27),
	(65561, 'WALUHU', 5121, 368, 27),
	(65562, 'ALO', 5122, 368, 27),
	(65563, 'BUNGA', 5122, 368, 27),
	(65564, 'INOMATA', 5122, 368, 27),
	(65565, 'LAUT BIRU', 5122, 368, 27),
	(65566, 'MOOPIYA (MOPUYA)', 5122, 368, 27),
	(65567, 'MOOTAWA', 5122, 368, 27),
	(65568, 'MOOTAYU', 5122, 368, 27),
	(65569, 'MOOTINELO', 5122, 368, 27),
	(65570, 'PELITA JAYA', 5122, 368, 27),
	(65571, 'TOMBULILATO', 5122, 368, 27),
	(65572, 'BATU HIJAU', 5123, 368, 27),
	(65573, 'BILUNGALA', 5123, 368, 27),
	(65574, 'BILUNGALA UTARA', 5123, 368, 27),
	(65575, 'KEMIRI', 5123, 368, 27),
	(65576, 'LEMBAH HIJAU', 5123, 368, 27),
	(65577, 'OMBULO HIJAU', 5123, 368, 27),
	(65578, 'PELITA HIJAU', 5123, 368, 27),
	(65579, 'TAMBOO', 5123, 368, 27),
	(65580, 'TIHU', 5123, 368, 27),
	(65581, 'TOLOTIO', 5123, 368, 27),
	(65582, 'TONGO', 5123, 368, 27),
	(65583, 'TUNAS JAYA', 5123, 368, 27),
	(65584, 'UABANGA', 5123, 368, 27),
	(65585, 'BUATA', 5124, 368, 27),
	(65586, 'LUWOHU', 5124, 368, 27),
	(65587, 'PANGGULO', 5124, 368, 27),
	(65588, 'PANGGULO BARAT', 5124, 368, 27),
	(65589, 'SUKMA', 5124, 368, 27),
	(65590, 'TANAH PUTIH', 5124, 368, 27),
	(65591, 'TIMBUOLO', 5124, 368, 27),
	(65592, 'TIMBUOLO TENGAH', 5124, 368, 27),
	(65593, 'TIMBUOLO TIMUR', 5124, 368, 27),
	(65594, 'AYULA SELATAN', 5125, 368, 27),
	(65595, 'AYULA TILANGO', 5125, 368, 27),
	(65596, 'AYULA TIMUR', 5125, 368, 27),
	(65597, 'AYULA UTARA', 5125, 368, 27),
	(65598, 'HUNTU SELATAN', 5125, 368, 27),
	(65599, 'HUNTU UTARA', 5125, 368, 27),
	(65600, 'LAMAHU', 5125, 368, 27),
	(65601, 'MEKAR JAYA', 5125, 368, 27),
	(65602, 'SEJAHTERA', 5125, 368, 27),
	(65603, 'TINELO AYULA', 5125, 368, 27),
	(65604, 'BULOTALANGI', 5126, 368, 27),
	(65605, 'BULOTALANGI BARAT', 5126, 368, 27),
	(65606, 'BULOTALANGI TIMUR', 5126, 368, 27),
	(65607, 'POPODU', 5126, 368, 27),
	(65608, 'TOLUWAYA (TULOA)', 5126, 368, 27),
	(65609, 'ILOMATA', 5127, 368, 27),
	(65610, 'MONGIILO', 5127, 368, 27),
	(65611, 'MONGIILO UTARA', 5127, 368, 27),
	(65612, 'OWATA', 5127, 368, 27),
	(65613, 'PILOLAHEYA', 5127, 368, 27),
	(65614, 'SUKA MAKMUR', 5127, 368, 27),
	(65615, 'BANDUNGAN', 5128, 368, 27),
	(65616, 'BOIDU', 5128, 368, 27),
	(65617, 'BUNUO', 5128, 368, 27),
	(65618, 'KOPI', 5128, 368, 27),
	(65619, 'LOMAYA', 5128, 368, 27),
	(65620, 'LONGALO', 5128, 368, 27),
	(65621, 'SUKA DAMAI', 5128, 368, 27),
	(65622, 'TULOA', 5128, 368, 27),
	(65623, 'TUPA', 5128, 368, 27),
	(65624, 'BUKIT HIJAU', 5129, 368, 27),
	(65625, 'BUNGA HIJAU', 5129, 368, 27),
	(65626, 'KAIDUNDU', 5129, 368, 27),
	(65627, 'KAIDUNDU BARAT', 5129, 368, 27),
	(65628, 'MAMUNGAA', 5129, 368, 27),
	(65629, 'MAMUNGAA TIMUR', 5129, 368, 27),
	(65630, 'MOPUYA', 5129, 368, 27),
	(65631, 'NYIUR HIJAU', 5129, 368, 27),
	(65632, 'PATOA', 5129, 368, 27),
	(65633, 'DUTOHE', 5130, 368, 27),
	(65634, 'DUTOHE BARAT', 5130, 368, 27),
	(65635, 'OLUHUTA', 5130, 368, 27),
	(65636, 'OLUHUTA UTARA', 5130, 368, 27),
	(65637, 'PADENGO', 5130, 368, 27),
	(65638, 'PAUWO', 5130, 368, 27),
	(65639, 'POOWO', 5130, 368, 27),
	(65640, 'POOWO BARAT', 5130, 368, 27),
	(65641, 'TALANGO', 5130, 368, 27),
	(65642, 'TANGGILINGO', 5130, 368, 27),
	(65643, 'TOTO SELATAN', 5130, 368, 27),
	(65644, 'TUMBIHE', 5130, 368, 27),
	(65645, 'BILUANGO', 5131, 368, 27),
	(65646, 'BINTALAHE', 5131, 368, 27),
	(65647, 'BOTUBARANI', 5131, 368, 27),
	(65648, 'BOTUTONUO (BOTUTONUA)', 5131, 368, 27),
	(65649, 'HUANGOBOTU', 5131, 368, 27),
	(65650, 'MODELOMO', 5131, 368, 27),
	(65651, 'MOLOTABU (MOLUTABU)', 5131, 368, 27),
	(65652, 'OLELE', 5131, 368, 27),
	(65653, 'OLUHUTA', 5131, 368, 27),
	(65654, 'BANGIO', 5132, 368, 27),
	(65655, 'DATARAN HIJAU', 5132, 368, 27),
	(65656, 'PINOGU', 5132, 368, 27),
	(65657, 'PINOGU PERMAI', 5132, 368, 27),
	(65658, 'TILONGGIBILA', 5132, 368, 27),
	(65659, 'BOLUDAWA', 5133, 368, 27),
	(65660, 'BUBE', 5133, 368, 27),
	(65661, 'BUBE BARU', 5133, 368, 27),
	(65662, 'BUBEYA', 5133, 368, 27),
	(65663, 'HELUMO', 5133, 368, 27),
	(65664, 'HULUDUOTAMO', 5133, 368, 27),
	(65665, 'TINELO', 5133, 368, 27),
	(65666, 'TINGKOHUBU', 5133, 368, 27),
	(65667, 'TINGKOHUBU TIMUR', 5133, 368, 27),
	(65668, 'ULANTA', 5133, 368, 27),
	(65669, 'BONDA RAYA', 5134, 368, 27),
	(65670, 'BONDAWUNA', 5134, 368, 27),
	(65671, 'BONEDAA', 5134, 368, 27),
	(65672, 'BULONTALA', 5134, 368, 27),
	(65673, 'BULONTALA TIMUR', 5134, 368, 27),
	(65674, 'LIBUNGO', 5134, 368, 27),
	(65675, 'MOLINTOGUPO', 5134, 368, 27),
	(65676, 'PANCURAN', 5134, 368, 27),
	(65677, 'ALELE', 5135, 368, 27),
	(65678, 'DUANO', 5135, 368, 27),
	(65679, 'LOMBONGO', 5135, 368, 27),
	(65680, 'LOMPOTOO', 5135, 368, 27),
	(65681, 'TAPADAA', 5135, 368, 27),
	(65682, 'TOLOMATO', 5135, 368, 27),
	(65683, 'DATARAN HIJAU', 5136, 368, 27),
	(65684, 'DUMBAYABULAN', 5136, 368, 27),
	(65685, 'DUMBAYABULAN TIMUR', 5136, 368, 27),
	(65686, 'PANGGULO', 5136, 368, 27),
	(65687, 'PODUOMA', 5136, 368, 27),
	(65688, 'TILANGOBULA', 5136, 368, 27),
	(65689, 'TULABOLO', 5136, 368, 27),
	(65690, 'TULABOLO BARAT', 5136, 368, 27),
	(65691, 'TULABOLO TIMUR', 5136, 368, 27),
	(65692, 'DUNGGALA', 5137, 368, 27),
	(65693, 'KERAMAT', 5137, 368, 27),
	(65694, 'LANGGE', 5137, 368, 27),
	(65695, 'MERANTI', 5137, 368, 27),
	(65696, 'TALULOBUTU', 5137, 368, 27),
	(65697, 'TALULOBUTU SELATAN', 5137, 368, 27),
	(65698, 'TALUMOPATU', 5137, 368, 27),
	(65699, 'BERLIAN', 5138, 368, 27),
	(65700, 'BONGOHULOWA', 5138, 368, 27),
	(65701, 'BONGOIME', 5138, 368, 27),
	(65702, 'BONGOPINI', 5138, 368, 27),
	(65703, 'BUTU', 5138, 368, 27),
	(65704, 'ILOHELUMA', 5138, 368, 27),
	(65705, 'LONUO', 5138, 368, 27),
	(65706, 'MOTILANGO (MOOTILANGO)', 5138, 368, 27),
	(65707, 'MOUTONG', 5138, 368, 27),
	(65708, 'PERMATA', 5138, 368, 27),
	(65709, 'TAMBOO', 5138, 368, 27),
	(65710, 'TOTO UTARA', 5138, 368, 27),
	(65711, 'TUNGGULO', 5138, 368, 27),
	(65712, 'TUNGGULO SELATAN', 5138, 368, 27),
	(65713, 'BIHE', 5139, 369, 27),
	(65714, 'BONTULO', 5139, 369, 27),
	(65715, 'BULULI', 5139, 369, 27),
	(65716, 'KARYA BARU', 5139, 369, 27),
	(65717, 'KARYA INDAH', 5139, 369, 27),
	(65718, 'MOHIYOLO', 5139, 369, 27),
	(65719, 'OLIMOHULO', 5139, 369, 27),
	(65720, 'PANGAHU', 5139, 369, 27),
	(65721, 'PRIMA', 5139, 369, 27),
	(65722, 'TIOHU', 5139, 369, 27),
	(65723, 'BARAKATI', 5140, 369, 27),
	(65724, 'BUA', 5140, 369, 27),
	(65725, 'DUNGGALA', 5140, 369, 27),
	(65726, 'HUNTU', 5140, 369, 27),
	(65727, 'ILOHUNGAYO', 5140, 369, 27),
	(65728, 'ILUTA', 5140, 369, 27),
	(65729, 'PAYUNGA', 5140, 369, 27),
	(65730, 'PILOBUHUTA', 5140, 369, 27),
	(65731, 'BILUHU TIMUR', 5141, 369, 27),
	(65732, 'BONGO', 5141, 369, 27),
	(65733, 'BUHUDAA', 5141, 369, 27),
	(65734, 'KAYUBULAN', 5141, 369, 27),
	(65735, 'LAMU', 5141, 369, 27),
	(65736, 'LANGGULA', 5141, 369, 27),
	(65737, 'LOPO', 5141, 369, 27),
	(65738, 'OLIMOO', 5141, 369, 27),
	(65739, 'TONTAYUO', 5141, 369, 27),
	(65740, 'BILATO', 5142, 369, 27),
	(65741, 'BUMELA', 5142, 369, 27),
	(65742, 'ILOMATA', 5142, 369, 27),
	(65743, 'JURIYA', 5142, 369, 27),
	(65744, 'LAMAHU', 5142, 369, 27),
	(65745, 'MUSYAWARAH', 5142, 369, 27),
	(65746, 'PELEHU', 5142, 369, 27),
	(65747, 'SUKA DAMAI', 5142, 369, 27),
	(65748, 'TAULAA', 5142, 369, 27),
	(65749, 'TOTOPO', 5142, 369, 27),
	(65750, 'BILUHU BARAT', 5143, 369, 27),
	(65751, 'BILUHU TENGAH', 5143, 369, 27),
	(65752, 'BOTUBOLU\\O (BOTUBOLUO)', 5143, 369, 27),
	(65753, 'HUWONGO', 5143, 369, 27),
	(65754, 'LOBUTO', 5143, 369, 27),
	(65755, 'LOBUTO TIMUR', 5143, 369, 27),
	(65756, 'LULUO', 5143, 369, 27),
	(65757, 'OLIMEYALA', 5143, 369, 27),
	(65758, 'BANDUNG REJO', 5144, 369, 27),
	(65759, 'BONGONGAYU', 5144, 369, 27),
	(65760, 'DILONIYOHU', 5144, 369, 27),
	(65761, 'DULOHUPA', 5144, 369, 27),
	(65762, 'ILOHELUMA', 5144, 369, 27),
	(65763, 'MONGGOLITO', 5144, 369, 27),
	(65764, 'MOTODUTO', 5144, 369, 27),
	(65765, 'PARUNGI', 5144, 369, 27),
	(65766, 'POTANGA', 5144, 369, 27),
	(65767, 'SIDODADI', 5144, 369, 27),
	(65768, 'SIDOMULYO', 5144, 369, 27),
	(65769, 'SIDOMULYO SELATAN', 5144, 369, 27),
	(65770, 'TOLITE', 5144, 369, 27),
	(65771, 'BATU LORENG', 5145, 369, 27),
	(65772, 'BATULAYAR', 5145, 369, 27),
	(65773, 'BONGOHULAWA', 5145, 369, 27),
	(65774, 'DULAMAYO', 5145, 369, 27),
	(65775, 'HUNTU LO HULAWA', 5145, 369, 27),
	(65776, 'KAYUMERAH', 5145, 369, 27),
	(65777, 'LIYODU', 5145, 369, 27),
	(65778, 'LIYOTO', 5145, 369, 27),
	(65779, 'MOLANIHU', 5145, 369, 27),
	(65780, 'MOLAS', 5145, 369, 27),
	(65781, 'MOLOPATODU', 5145, 369, 27),
	(65782, 'OTOPADE', 5145, 369, 27),
	(65783, 'OWALANGA', 5145, 369, 27),
	(65784, 'TOHUPO', 5145, 369, 27),
	(65785, 'UPOMELA', 5145, 369, 27),
	(65786, 'BOTU', 5146, 369, 27),
	(65787, 'BUGIS', 5146, 369, 27),
	(65788, 'LEATO SELATAN', 5146, 369, 27),
	(65789, 'LEATO UTARA', 5146, 369, 27),
	(65790, 'TALUMOLO', 5146, 369, 27),
	(65791, 'AMBARA', 5147, 369, 27),
	(65792, 'AYUHULA', 5147, 369, 27),
	(65793, 'BONGOMEME', 5147, 369, 27),
	(65794, 'BOTUBU LOWE', 5147, 369, 27),
	(65795, 'DUNGALIYO', 5147, 369, 27),
	(65796, 'DUWANGA', 5147, 369, 27),
	(65797, 'KALIYOSO', 5147, 369, 27),
	(65798, 'MOMALA', 5147, 369, 27),
	(65799, 'PANGADAA', 5147, 369, 27),
	(65800, 'PILOLALENGA', 5147, 369, 27),
	(65801, 'HUANGOBOTU', 5148, 369, 27),
	(65802, 'LIBUO', 5148, 369, 27),
	(65803, 'TOMULABUTAO', 5148, 369, 27),
	(65804, 'TOMULABUTAO SELATAN', 5148, 369, 27),
	(65805, 'TULADENGGI', 5148, 369, 27),
	(65806, 'DONGGALA', 5149, 369, 27),
	(65807, 'POHE', 5149, 369, 27),
	(65808, 'SIENDENG', 5149, 369, 27),
	(65809, 'TANJUNG KRAMAT', 5149, 369, 27),
	(65810, 'TENDA', 5149, 369, 27),
	(65811, 'BULADU', 5150, 369, 27),
	(65812, 'BULIIDE', 5150, 369, 27),
	(65813, 'DEMBE I', 5150, 369, 27),
	(65814, 'LEKOBALO', 5150, 369, 27),
	(65815, 'MOLOSIFAT W (MOLOSIPAT W)', 5150, 369, 27),
	(65816, 'PILOLODAA', 5150, 369, 27),
	(65817, 'TENILO', 5150, 369, 27),
	(65818, 'BIAWAO', 5151, 369, 27),
	(65819, 'BIAWU', 5151, 369, 27),
	(65820, 'LIMBA B', 5151, 369, 27),
	(65821, 'LIMBA U DUA (LIMBA U II)', 5151, 369, 27),
	(65822, 'LIMBA U SATU (LIMBA U I)', 5151, 369, 27),
	(65823, 'DULALOWO', 5152, 369, 27),
	(65824, 'DULALOWO TIMUR', 5152, 369, 27),
	(65825, 'LILUWO', 5152, 369, 27),
	(65826, 'PAGUYAMAN', 5152, 369, 27),
	(65827, 'PULUBALA', 5152, 369, 27),
	(65828, 'WUMIALO', 5152, 369, 27),
	(65829, 'HELEDULAA SELATAN', 5153, 369, 27),
	(65830, 'HELEDULAA UTARA', 5153, 369, 27),
	(65831, 'IPILO', 5153, 369, 27),
	(65832, 'MOODU', 5153, 369, 27),
	(65833, 'PADEBUOLO', 5153, 369, 27),
	(65834, 'TAMALATE', 5153, 369, 27),
	(65835, 'DEMBE II', 5154, 369, 27),
	(65836, 'DEMBE JAYA', 5154, 369, 27),
	(65837, 'DULOMO SELATAN', 5154, 369, 27),
	(65838, 'DULOMO UTARA', 5154, 369, 27),
	(65839, 'WONGKADITI BARAT', 5154, 369, 27),
	(65840, 'WONGKADITI TIMUR', 5154, 369, 27),
	(65841, 'BIYONGA', 5155, 369, 27),
	(65842, 'BOLIHUANGGA', 5155, 369, 27),
	(65843, 'BONGOHULAWA', 5155, 369, 27),
	(65844, 'BULOTA', 5155, 369, 27),
	(65845, 'DUTULANAA', 5155, 369, 27),
	(65846, 'HEPUHULAWA', 5155, 369, 27),
	(65847, 'HUNGGALUWA', 5155, 369, 27),
	(65848, 'HUTUO', 5155, 369, 27),
	(65849, 'KAYUBULAN', 5155, 369, 27),
	(65850, 'KAYUMERAH', 5155, 369, 27),
	(65851, 'MALAHU', 5155, 369, 27),
	(65852, 'POLOHUNGO', 5155, 369, 27),
	(65853, 'TENILO', 5155, 369, 27),
	(65854, 'TILIHUA', 5155, 369, 27),
	(65855, 'DAENAA', 5156, 369, 27),
	(65856, 'HAYAHAYA', 5156, 369, 27),
	(65857, 'HUIDU', 5156, 369, 27),
	(65858, 'HUIDU UTARA', 5156, 369, 27),
	(65859, 'HUTABOHU', 5156, 369, 27),
	(65860, 'OMBULO', 5156, 369, 27),
	(65861, 'PADENGO', 5156, 369, 27),
	(65862, 'PONE', 5156, 369, 27),
	(65863, 'TUNGGULO', 5156, 369, 27),
	(65864, 'YOSONEGORO (YOOSONEGORO)', 5156, 369, 27),
	(65865, 'HELUMO', 5157, 369, 27),
	(65866, 'HUYULA', 5157, 369, 27),
	(65867, 'KARYA MUKTI', 5157, 369, 27),
	(65868, 'PARIS', 5157, 369, 27),
	(65869, 'PAYU', 5157, 369, 27),
	(65870, 'PILOMONU', 5157, 369, 27),
	(65871, 'SATRIA', 5157, 369, 27),
	(65872, 'SIDO MUKTI', 5157, 369, 27),
	(65873, 'SUKA MAJU', 5157, 369, 27),
	(65874, 'TALUMOPATU', 5157, 369, 27),
	(65875, 'AYUMOLINGO', 5158, 369, 27),
	(65876, 'BAKTI', 5158, 369, 27),
	(65877, 'BUKIT AREN', 5158, 369, 27),
	(65878, 'MOLALAHU', 5158, 369, 27),
	(65879, 'MOLAMAHU', 5158, 369, 27),
	(65880, 'MULYONEGORO', 5158, 369, 27),
	(65881, 'PONGONGAILA', 5158, 369, 27),
	(65882, 'PULUBALA', 5158, 369, 27),
	(65883, 'PUNCAK', 5158, 369, 27),
	(65884, 'TOYIDITO (TOYDITO)', 5158, 369, 27),
	(65885, 'TRIDARMA', 5158, 369, 27),
	(65886, 'BULOTADAA BARAT', 5159, 369, 27),
	(65887, 'BULOTADAA TIMUR', 5159, 369, 27),
	(65888, 'MOLOSIFAT U (MOLOSIPAT U)', 5159, 369, 27),
	(65889, 'TANGGIKIKI', 5159, 369, 27),
	(65890, 'TAPA', 5159, 369, 27),
	(65891, 'ILOMANGGA', 5160, 369, 27),
	(65892, 'LIMEHE BARAT', 5160, 369, 27),
	(65893, 'LIMEHE TIMUR', 5160, 369, 27),
	(65894, 'LIMEHU', 5160, 369, 27),
	(65895, 'MOAHUDU', 5160, 369, 27),
	(65896, 'MOTINELO', 5160, 369, 27),
	(65897, 'TABONGO BARAT', 5160, 369, 27),
	(65898, 'TABONGO TIMUR', 5160, 369, 27),
	(65899, 'TERATAI', 5160, 369, 27),
	(65900, 'BULILA', 5161, 369, 27),
	(65901, 'DULAMAYO BARAT', 5161, 369, 27),
	(65902, 'DULAMAYO SELATAN', 5161, 369, 27),
	(65903, 'DULOHUPA', 5161, 369, 27),
	(65904, 'HULAWA', 5161, 369, 27),
	(65905, 'LUHU', 5161, 369, 27),
	(65906, 'MONGOLATO', 5161, 369, 27),
	(65907, 'PILOHAYANGA', 5161, 369, 27),
	(65908, 'PILOHAYANGA BARAT', 5161, 369, 27),
	(65909, 'DULAMAYO UTARA', 5162, 369, 27),
	(65910, 'DUMATI', 5162, 369, 27),
	(65911, 'LUPOYO', 5162, 369, 27),
	(65912, 'MODELIDU (MODELLIDU)', 5162, 369, 27),
	(65913, 'PANTUNGO', 5162, 369, 27),
	(65914, 'PENTADIO BARAT', 5162, 369, 27),
	(65915, 'PENTADIO TIMUR', 5162, 369, 27),
	(65916, 'TALUMELITO', 5162, 369, 27),
	(65917, 'TAPALULUO', 5162, 369, 27),
	(65918, 'TIMUATO', 5162, 369, 27),
	(65919, 'TINELO', 5162, 369, 27),
	(65920, 'TONALA', 5162, 369, 27),
	(65921, 'TULADENGGI', 5162, 369, 27),
	(65922, 'ULAPATO A', 5162, 369, 27),
	(65923, 'ULAPATO B', 5162, 369, 27),
	(65924, 'BUHU', 5163, 369, 27),
	(65925, 'BULOTA', 5163, 369, 27),
	(65926, 'BUNGGALO', 5163, 369, 27),
	(65927, 'HUTADAA', 5163, 369, 27),
	(65928, 'LUWOO', 5163, 369, 27),
	(65929, 'BALAHU', 5164, 369, 27),
	(65930, 'BOTUMOPUTI', 5164, 369, 27),
	(65931, 'BUHU', 5164, 369, 27),
	(65932, 'DATAHU', 5164, 369, 27),
	(65933, 'DUNGGALA', 5164, 369, 27),
	(65934, 'ILOMATA', 5164, 369, 27),
	(65935, 'ILOPONU', 5164, 369, 27),
	(65936, 'ISIMU RAYA', 5164, 369, 27),
	(65937, 'ISIMU SELATAN', 5164, 369, 27),
	(65938, 'ISIMU UTARA', 5164, 369, 27),
	(65939, 'LABANU', 5164, 369, 27),
	(65940, 'MOLOWAHU', 5164, 369, 27),
	(65941, 'MOTILANGO', 5164, 369, 27),
	(65942, 'REKSONEGORO', 5164, 369, 27),
	(65943, 'TOLOTIO', 5164, 369, 27),
	(65944, 'ULOBUA', 5164, 369, 27),
	(65945, 'DULOMO', 5165, 369, 27),
	(65946, 'ILOTIDEAA', 5165, 369, 27),
	(65947, 'LAUWONU', 5165, 369, 27),
	(65948, 'TABUMELA', 5165, 369, 27),
	(65949, 'TENGGELA', 5165, 369, 27),
	(65950, 'TILOTE', 5165, 369, 27),
	(65951, 'TINELO', 5165, 369, 27),
	(65952, 'TUALANGO', 5165, 369, 27),
	(65953, 'BINA JAYA', 5166, 369, 27),
	(65954, 'GANDARIA', 5166, 369, 27),
	(65955, 'GANDASARI', 5166, 369, 27),
	(65956, 'HIMALAYA', 5166, 369, 27),
	(65957, 'LAKEYA', 5166, 369, 27),
	(65958, 'MAKMUR ABADI', 5166, 369, 27),
	(65959, 'MARGOMULYO', 5166, 369, 27),
	(65960, 'MOLOHU', 5166, 369, 27),
	(65961, 'OMBULO TANGO', 5166, 369, 27),
	(65962, 'POLOHUNGO', 5166, 369, 27),
	(65963, 'SIDOHARJO (SIDOARJO)', 5166, 369, 27),
	(65964, 'SUKAMAKMUR', 5166, 369, 27),
	(65965, 'SUKAMAKMUR UTARA', 5166, 369, 27),
	(65966, 'TAMAILA', 5166, 369, 27),
	(65967, 'TAMALIA UTARA', 5166, 369, 27),
	(65968, 'DATAHU', 5167, 370, 27),
	(65969, 'DUDEPO', 5167, 370, 27),
	(65970, 'HELUMO', 5167, 370, 27),
	(65971, 'HIYALOOILE', 5167, 370, 27),
	(65972, 'IBARAT', 5167, 370, 27),
	(65973, 'ILANGATA', 5167, 370, 27),
	(65974, 'ILODULUNGA', 5167, 370, 27),
	(65975, 'ILOHELUMA', 5167, 370, 27),
	(65976, 'LANGGE', 5167, 370, 27),
	(65977, 'MOOTILANGO', 5167, 370, 27),
	(65978, 'POPALO', 5167, 370, 27),
	(65979, 'PUTIANA', 5167, 370, 27),
	(65980, 'TOLANGO', 5167, 370, 27),
	(65981, 'TOLONGIO', 5167, 370, 27),
	(65982, 'TUTUWOTO', 5167, 370, 27),
	(65983, 'BINTANA', 5168, 370, 27),
	(65984, 'BUATA', 5168, 370, 27),
	(65985, 'ILOHELUMA', 5168, 370, 27),
	(65986, 'IMANA', 5168, 370, 27),
	(65987, 'INOMATA (ILOMATA)', 5168, 370, 27),
	(65988, 'KOTAJIN (KOTA JIN)', 5168, 370, 27),
	(65989, 'KOTAJIN UTARA', 5168, 370, 27),
	(65990, 'MONGGUPO', 5168, 370, 27),
	(65991, 'OLUHUTA', 5168, 370, 27),
	(65992, 'PINONTOYONGA', 5168, 370, 27),
	(65993, 'POSONO', 5168, 370, 27),
	(65994, 'SIGASO', 5168, 370, 27),
	(65995, 'TOMBULILATO', 5168, 370, 27),
	(65996, 'WAPALO', 5168, 370, 27),
	(65997, 'BIAWU (BIAU)', 5169, 370, 27),
	(65998, 'BOHULO', 5169, 370, 27),
	(65999, 'BUALO', 5169, 370, 27),
	(66000, 'BUGIS', 5169, 370, 27)
	`)
}
