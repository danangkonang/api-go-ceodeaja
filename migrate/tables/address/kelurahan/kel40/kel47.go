package kel40

import "github.com/danangkonang/rest-api/config"

func Kel47() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
	(47001, 'PANGGAL PANGGAL', 3480, 233, 15),
	(47002, 'PENGARINGAN', 3480, 233, 15),
	(47003, 'RAKSA JIWA', 3480, 233, 15),
	(47004, 'SELEMAN', 3480, 233, 15),
	(47005, 'SINGAPURA', 3480, 233, 15),
	(47006, 'SUKA MERINDU', 3480, 233, 15),
	(47007, 'SUKARAMI', 3480, 233, 15),
	(47008, 'TANJUNG KURUNG', 3480, 233, 15),
	(47009, 'TEBING KAMPUNG', 3480, 233, 15),
	(47010, 'TUBOHAN', 3480, 233, 15),
	(47011, 'ULAK PANDAN', 3480, 233, 15),
	(47012, 'KARYA JAYA', 3481, 233, 15),
	(47013, 'KARYA MUKTI', 3481, 233, 15),
	(47014, 'MARGA BAKTI', 3481, 233, 15),
	(47015, 'MARGA MULYA', 3481, 233, 15),
	(47016, 'SRI MULYA', 3481, 233, 15),
	(47017, 'TANJUNG MAKMUR', 3481, 233, 15),
	(47018, 'BANDAR', 3482, 233, 15),
	(47019, 'KUNGKILAN', 3482, 233, 15),
	(47020, 'LUBUK BARU', 3482, 233, 15),
	(47021, 'LUBUK LEBAN', 3482, 233, 15),
	(47022, 'MEKAR JAYA', 3482, 233, 15),
	(47023, 'MEKAR SARI', 3482, 233, 15),
	(47024, 'NEGERI SINDANG', 3482, 233, 15),
	(47025, 'PENANTIAN', 3482, 233, 15),
	(47026, 'PENYANDINGAN', 3482, 233, 15),
	(47027, 'RANTAU KUMPAI', 3482, 233, 15),
	(47028, 'TUNGKU JAYA', 3482, 233, 15),
	(47029, 'BELANDANG', 3483, 233, 15),
	(47030, 'GUNUNG TIGA', 3483, 233, 15),
	(47031, 'KELUMPANG', 3483, 233, 15),
	(47032, 'MENDINGIN', 3483, 233, 15),
	(47033, 'PEDATARAN', 3483, 233, 15),
	(47034, 'SUKAJADI', 3483, 233, 15),
	(47035, 'ULAK LEBAR', 3483, 233, 15),
	(47036, 'AIR RUPIK', 3484, 234, 15),
	(47037, 'BANDAR AGUNG RANAU', 3484, 234, 15),
	(47038, 'BANDING AGUNG', 3484, 234, 15),
	(47039, 'KARANG INDAH', 3484, 234, 15),
	(47040, 'PENANTIAN', 3484, 234, 15),
	(47041, 'RANTAU NIPIS', 3484, 234, 15),
	(47042, 'SIDODADI', 3484, 234, 15),
	(47043, 'SIPATUHU', 3484, 234, 15),
	(47044, 'SIPATUHU II', 3484, 234, 15),
	(47045, 'SUGIH WARAS', 3484, 234, 15),
	(47046, 'SUKA NEGERI', 3484, 234, 15),
	(47047, 'SUKAMAJU', 3484, 234, 15),
	(47048, 'SUMBER MAKMUR', 3484, 234, 15),
	(47049, 'SURABAYA', 3484, 234, 15),
	(47050, 'SURABAYA TIMUR', 3484, 234, 15),
	(47051, 'TALANG MERBAU', 3484, 234, 15),
	(47052, 'TANGSI AGUNG', 3484, 234, 15),
	(47053, 'TANJUNG AGUNG', 3484, 234, 15),
	(47054, 'TANJUNG HARAPAN', 3484, 234, 15),
	(47055, 'TELANAI', 3484, 234, 15),
	(47056, 'TERAP MULIA', 3484, 234, 15),
	(47057, 'WAY TIMAH', 3484, 234, 15),
	(47058, 'BANDAR', 3485, 234, 15),
	(47059, 'DAMARPURA', 3485, 234, 15),
	(47060, 'GEMIUNG', 3485, 234, 15),
	(47061, 'JAGARAGA', 3485, 234, 15),
	(47062, 'SINAR DANAU', 3485, 234, 15),
	(47063, 'TANJUNG BERINGIN', 3485, 234, 15),
	(47064, 'TEKANA', 3485, 234, 15),
	(47065, 'TUNAS JAYA', 3485, 234, 15),
	(47066, 'AIR KELIAN', 3486, 234, 15),
	(47067, 'DANAU JAYA', 3486, 234, 15),
	(47068, 'DURIAN SEMBILAN', 3486, 234, 15),
	(47069, 'KARET JAYA', 3486, 234, 15),
	(47070, 'KEMBANG TINGGI', 3486, 234, 15),
	(47071, 'KOTA WAY', 3486, 234, 15),
	(47072, 'MEKAR JAYA', 3486, 234, 15),
	(47073, 'SERAKAT JAYA', 3486, 234, 15),
	(47074, 'SIDO RAHAYU', 3486, 234, 15),
	(47075, 'SIDODADI', 3486, 234, 15),
	(47076, 'SINAR BARU', 3486, 234, 15),
	(47077, 'SINAR NAPALAN', 3486, 234, 15),
	(47078, 'SIPIN', 3486, 234, 15),
	(47079, 'SRI MENANTI', 3486, 234, 15),
	(47080, 'SUMBER RAYA', 3486, 234, 15),
	(47081, 'SUMBER RINGIN', 3486, 234, 15),
	(47082, 'TALANG PADANG', 3486, 234, 15),
	(47083, 'TANJUNG BARU', 3486, 234, 15),
	(47084, 'TANJUNG DURIAN', 3486, 234, 15),
	(47085, 'TANJUNG JAYA', 3486, 234, 15),
	(47086, 'TANJUNG MENANG', 3486, 234, 15),
	(47087, 'TANJUNG SARI', 3486, 234, 15),
	(47088, 'GEDONG/GEDUNG BARU', 3487, 234, 15),
	(47089, 'HANGKUSA/LINGKUSA', 3487, 234, 15),
	(47090, 'JEPARA', 3487, 234, 15),
	(47091, 'PADANG RATU', 3487, 234, 15),
	(47092, 'PAKHDA SUKA', 3487, 234, 15),
	(47093, 'SERUMPUN JAYA', 3487, 234, 15),
	(47094, 'SIMPANG SENDER', 3487, 234, 15),
	(47095, 'SIMPANG SENDER SELATAN', 3487, 234, 15),
	(47096, 'SIMPANG SENDER TENGAH', 3487, 234, 15),
	(47097, 'SIMPANG SENDER TIMUR', 3487, 234, 15),
	(47098, 'SIMPANG SENDER UTARA', 3487, 234, 15),
	(47099, 'SUBIK', 3487, 234, 15),
	(47100, 'SUKABUMI', 3487, 234, 15),
	(47101, 'SUKAMARGA', 3487, 234, 15),
	(47102, 'SUKARAMI', 3487, 234, 15),
	(47103, 'SUMBER JAYA', 3487, 234, 15),
	(47104, 'SUMBER MULIA', 3487, 234, 15),
	(47105, 'TANJUNG BARU RANAU', 3487, 234, 15),
	(47106, 'TANJUNG KEMALA', 3487, 234, 15),
	(47107, 'TANJUNG SARI', 3487, 234, 15),
	(47108, 'TANJUNG SETIA', 3487, 234, 15),
	(47109, 'WAY RELAI', 3487, 234, 15),
	(47110, 'BANJAR AGUNG', 3488, 234, 15),
	(47111, 'BENDI', 3488, 234, 15),
	(47112, 'BUMI AGUNG JAYA', 3488, 234, 15),
	(47113, 'BUMIJAYA', 3488, 234, 15),
	(47114, 'GUNUNG CAHYA', 3488, 234, 15),
	(47115, 'MAJAR', 3488, 234, 15),
	(47116, 'PEKUOLAN', 3488, 234, 15),
	(47117, 'PELAWI', 3488, 234, 15),
	(47118, 'RANTAU PANJANG', 3488, 234, 15),
	(47119, 'RUOS', 3488, 234, 15),
	(47120, 'SUKAJAYA', 3488, 234, 15),
	(47121, 'BEDENG BLAMBANGAN', 3489, 234, 15),
	(47122, 'BELAMBANGAN', 3489, 234, 15),
	(47123, 'KAGELANG BLAMBANGAN', 3489, 234, 15),
	(47124, 'KOTA AMAN', 3489, 234, 15),
	(47125, 'NAGAR AGUNG', 3489, 234, 15),
	(47126, 'NEGERI BATIN BARU', 3489, 234, 15),
	(47127, 'PADANG BINDU', 3489, 234, 15),
	(47128, 'PADANG SARI', 3489, 234, 15),
	(47129, 'PENINJAUAN', 3489, 234, 15),
	(47130, 'PERUPUS BLAMBANGAN', 3489, 234, 15),
	(47131, 'SAUNG NAGA', 3489, 234, 15),
	(47132, 'SIMPANG SAGA', 3489, 234, 15),
	(47133, 'SUGIH WARAS', 3489, 234, 15),
	(47134, 'SUKAJADI BLAMBANGAN', 3489, 234, 15),
	(47135, 'BUNGA MAS', 3490, 234, 15),
	(47136, 'GUNUNG TERANG', 3490, 234, 15),
	(47137, 'KENALI', 3490, 234, 15),
	(47138, 'KOTA KARANG', 3490, 234, 15),
	(47139, 'LUBUK LIKU', 3490, 234, 15),
	(47140, 'MADURA', 3490, 234, 15),
	(47141, 'NEGERI AGUNG', 3490, 234, 15),
	(47142, 'NEGERI BATIN', 3490, 234, 15),
	(47143, 'NEGERI CAHYA', 3490, 234, 15),
	(47144, 'SUKARAJA I', 3490, 234, 15),
	(47145, 'SUKARAMI', 3490, 234, 15),
	(47146, 'TALANG BARU', 3490, 234, 15),
	(47147, 'TANJUNG IMAN', 3490, 234, 15),
	(47148, 'TANJUNG MENANG ILIR', 3490, 234, 15),
	(47149, 'TANJUNG MENANG ULU', 3490, 234, 15),
	(47150, 'TANJUNG RAYA', 3490, 234, 15),
	(47151, 'CAMPANG JAYA', 3491, 234, 15),
	(47152, 'KEBAN AGUNG', 3491, 234, 15),
	(47153, 'MUARA SINDANG', 3491, 234, 15),
	(47154, 'PANGANDONAN', 3491, 234, 15),
	(47155, 'PIUS', 3491, 234, 15),
	(47156, 'PULAU KEMILING', 3491, 234, 15),
	(47157, 'SIMPANG CAMPANG', 3491, 234, 15),
	(47158, 'SIRING/SRING ALAM', 3491, 234, 15),
	(47159, 'TANJUNG JATI', 3491, 234, 15),
	(47160, 'AIR ALUN', 3492, 234, 15),
	(47161, 'BALAYAN', 3492, 234, 15),
	(47162, 'BANDAR ALAM LAMA', 3492, 234, 15),
	(47163, 'BERASANG', 3492, 234, 15),
	(47164, 'DANAU RATA', 3492, 234, 15),
	(47165, 'GUNUNG MEGANG', 3492, 234, 15),
	(47166, 'KOTA PADANG', 3492, 234, 15),
	(47167, 'MUARA PAYANG', 3492, 234, 15),
	(47168, 'PADANG BINDU', 3492, 234, 15),
	(47169, 'PADANG LAY', 3492, 234, 15),
	(47170, 'PAJAR BULAN', 3492, 234, 15),
	(47171, 'PULAU PANGGUNG', 3492, 234, 15),
	(47172, 'SIMPANG EMPAT', 3492, 234, 15),
	(47173, 'SIMPANG TIGA', 3492, 234, 15),
	(47174, 'SINGA LAGA', 3492, 234, 15),
	(47175, 'SIRING AGUNG', 3492, 234, 15),
	(47176, 'TEBAT GABUS', 3492, 234, 15),
	(47177, 'TENANG', 3492, 234, 15),
	(47178, 'ULAK PANDAN', 3492, 234, 15),
	(47179, 'AIR BARU', 3493, 234, 15),
	(47180, 'BUNUT', 3493, 234, 15),
	(47181, 'GALANG TINGGI', 3493, 234, 15),
	(47182, 'KEMANG BANDUNG', 3493, 234, 15),
	(47183, 'KEPAYANG', 3493, 234, 15),
	(47184, 'KOTA BARU', 3493, 234, 15),
	(47185, 'KOTA DALAM', 3493, 234, 15),
	(47186, 'PERE\\AN', 3493, 234, 15),
	(47187, 'PULAU DUKU', 3493, 234, 15),
	(47188, 'SELABUNG BELIMBING JAYA', 3493, 234, 15),
	(47189, 'SINAR MARGA', 3493, 234, 15),
	(47190, 'SRI MENANTI', 3493, 234, 15),
	(47191, 'SUKARAJA', 3493, 234, 15),
	(47192, 'TANJUNG BESAR', 3493, 234, 15),
	(47193, 'TELUK AGUNG', 3493, 234, 15),
	(47194, 'BATU BELANG DUA', 3494, 234, 15),
	(47195, 'BATU BELANG JAYA', 3494, 234, 15),
	(47196, 'BUMI AGUNG', 3494, 234, 15),
	(47197, 'DATAR', 3494, 234, 15),
	(47198, 'GEDUNG LEPIHAN', 3494, 234, 15),
	(47199, 'GUNUNG TIGA', 3494, 234, 15),
	(47200, 'KISAU', 3494, 234, 15),
	(47201, 'MEHANGGIN', 3494, 234, 15),
	(47202, 'PANCUR PUNGAH', 3494, 234, 15),
	(47203, 'PASAR MUARADUA', 3494, 234, 15),
	(47204, 'PELANGKI', 3494, 234, 15),
	(47205, 'PENDAGAN', 3494, 234, 15),
	(47206, 'SUKA BANJAR', 3494, 234, 15),
	(47207, 'SUKARAJA DUA', 3494, 234, 15),
	(47208, 'ALUN DUA', 3495, 234, 15),
	(47209, 'BANDAR ALAM BARU', 3495, 234, 15),
	(47210, 'BAYUR TENGAH', 3495, 234, 15),
	(47211, 'DUSUN TENGAH', 3495, 234, 15),
	(47212, 'GUNUNG GARE', 3495, 234, 15),
	(47213, 'LAWANG AGUNG', 3495, 234, 15),
	(47214, 'MUARA DUA KISAM', 3495, 234, 15),
	(47215, 'PAGAR DEWA', 3495, 234, 15),
	(47216, 'PENANTIAN', 3495, 234, 15),
	(47217, 'PENYANDINGAN', 3495, 234, 15),
	(47218, 'SIMPANG LUBUK DALAM', 3495, 234, 15),
	(47219, 'SUGIHAN', 3495, 234, 15),
	(47220, 'SUKANANTI', 3495, 234, 15),
	(47221, 'SUKARAJA', 3495, 234, 15),
	(47222, 'TANJUNG BERINGIN', 3495, 234, 15),
	(47223, 'TANJUNG TEBAT', 3495, 234, 15),
	(47224, 'ULAK AGUNG ILIR', 3495, 234, 15),
	(47225, 'ULAK AGUNG ULU', 3495, 234, 15),
	(47226, 'ANUGERAH KEMU', 3496, 234, 15),
	(47227, 'AROMANTAI', 3496, 234, 15),
	(47228, 'GUNUNG BATU', 3496, 234, 15),
	(47229, 'KEMU', 3496, 234, 15),
	(47230, 'KEMU ULU', 3496, 234, 15),
	(47231, 'PAGAR AGUNG', 3496, 234, 15),
	(47232, 'PEMATANG OBAR', 3496, 234, 15),
	(47233, 'PULAU BERINGIN', 3496, 234, 15),
	(47234, 'PULAU BERINGIN UTARA', 3496, 234, 15),
	(47235, 'SIMPANG PANCUR', 3496, 234, 15),
	(47236, 'TANJUNG BULAN', 3496, 234, 15),
	(47237, 'TANJUNG BULAN ULU', 3496, 234, 15),
	(47238, 'TANJUNG KARI', 3496, 234, 15),
	(47239, 'AIR BARU', 3497, 234, 15),
	(47240, 'BUMI GENAP', 3497, 234, 15),
	(47241, 'GEDUNG NYAWA', 3497, 234, 15),
	(47242, 'GEDUNG WANI', 3497, 234, 15),
	(47243, 'KARANG ENDAH', 3497, 234, 15),
	(47244, 'MERPANG', 3497, 234, 15),
	(47245, 'PENANGGUNGAN', 3497, 234, 15),
	(47246, 'SURA', 3497, 234, 15),
	(47247, 'TANJUNG KURUNG', 3497, 234, 15),
	(47248, 'BUNGIN CAMPANG', 3498, 234, 15),
	(47249, 'KARANG AGUNG', 3498, 234, 15),
	(47250, 'LUBAR', 3498, 234, 15),
	(47251, 'SIMPANG AGUNG', 3498, 234, 15),
	(47252, 'SIMPANGAN', 3498, 234, 15),
	(47253, 'SINAR MULYO', 3498, 234, 15),
	(47254, 'TANJUNG SARI', 3498, 234, 15),
	(47255, 'MUARA SINDANG ILIR', 3499, 234, 15),
	(47256, 'MUARA SINDANG TENGAH', 3499, 234, 15),
	(47257, 'PEMATANG DANAU', 3499, 234, 15),
	(47258, 'TANJUNG HARAPAN', 3499, 234, 15),
	(47259, 'TEBAT LAYANG', 3499, 234, 15),
	(47260, 'ULU DANAU', 3499, 234, 15),
	(47261, 'WATES', 3499, 234, 15),
	(47262, 'CUKOH NAU (CUKUH/CUKO NAU)', 3500, 234, 15),
	(47263, 'GUNTUNG JAYA', 3500, 234, 15),
	(47264, 'PECAH PINGGAN', 3500, 234, 15),
	(47265, 'PULAU KEMUNING', 3500, 234, 15),
	(47266, 'SADAU JAYA', 3500, 234, 15),
	(47267, 'SEBAJA', 3500, 234, 15),
	(47268, 'SIMPANG LUAS', 3500, 234, 15),
	(47269, 'TANAH PILIH', 3500, 234, 15),
	(47270, 'UJAN MAS', 3500, 234, 15),
	(47271, 'KARANG PENDETA', 3501, 234, 15),
	(47272, 'KOTA AGUNG', 3501, 234, 15),
	(47273, 'KURIPAN', 3501, 234, 15),
	(47274, 'KURIPAN II / DUA', 3501, 234, 15),
	(47275, 'PENINGGIRAN', 3501, 234, 15),
	(47276, 'SUKABUMI', 3501, 234, 15),
	(47277, 'SUKARENA', 3501, 234, 15),
	(47278, 'SURABAYA', 3501, 234, 15),
	(47279, 'BEDENG TIGA', 3502, 234, 15),
	(47280, 'BUMI AGUNG', 3502, 234, 15),
	(47281, 'GEDUNG RANAU', 3502, 234, 15),
	(47282, 'GUNUNG AJI', 3502, 234, 15),
	(47283, 'GUNUNG RAYA', 3502, 234, 15),
	(47284, 'KIWIS RAYA', 3502, 234, 15),
	(47285, 'KOTA BATU', 3502, 234, 15),
	(47286, 'MEKARSARI', 3502, 234, 15),
	(47287, 'PAGAR DEWA', 3502, 234, 15),
	(47288, 'PILLA', 3502, 234, 15),
	(47289, 'REMANAM JAYA', 3502, 234, 15),
	(47290, 'SEGIGOK RAYA', 3502, 234, 15),
	(47291, 'SUKAJAYA', 3502, 234, 15),
	(47292, 'TANJUNG BARU', 3502, 234, 15),
	(47293, 'TANJUNG JATI', 3502, 234, 15),
	(47294, 'WAY WANGI SEMINUNG', 3502, 234, 15),
	(47295, 'BANGUN REJO', 3503, 235, 15),
	(47296, 'BATU MAS', 3503, 235, 15),
	(47297, 'DARMA BUANA', 3503, 235, 15),
	(47298, 'KARANG JAYA', 3503, 235, 15),
	(47299, 'KARANG MANIK', 3503, 235, 15),
	(47300, 'KELI REJO', 3503, 235, 15),
	(47301, 'KEMUNING JAYA', 3503, 235, 15),
	(47302, 'MARGO MULYO', 3503, 235, 15),
	(47303, 'PURWOREJO', 3503, 235, 15),
	(47304, 'PURWOSARI', 3503, 235, 15),
	(47305, 'RAMAN JAYA', 3503, 235, 15),
	(47306, 'REJO MULYO', 3503, 235, 15),
	(47307, 'SRI BANTOLO', 3503, 235, 15),
	(47308, 'SRIJAYA', 3503, 235, 15),
	(47309, 'SUKA JAYA', 3503, 235, 15),
	(47310, 'SUMBER HARAPAN', 3503, 235, 15),
	(47311, 'SUMBER JAYA', 3503, 235, 15),
	(47312, 'SUMBER RAHAYU', 3503, 235, 15),
	(47313, 'SUMBER REJO', 3503, 235, 15),
	(47314, 'SUMBER SARI', 3503, 235, 15),
	(47315, 'TANJUNG KEMUNING', 3503, 235, 15),
	(47316, 'TEGAL BESAR', 3503, 235, 15),
	(47317, 'TEGAL SARI', 3503, 235, 15),
	(47318, 'TOTO REJO', 3503, 235, 15),
	(47319, 'DADI REJO', 3504, 235, 15),
	(47320, 'GANTI WARNO', 3504, 235, 15),
	(47321, 'KARANG JADI', 3504, 235, 15),
	(47322, 'KARANG SARI', 3504, 235, 15),
	(47323, 'KARYA MAJU', 3504, 235, 15),
	(47324, 'KUTOSARI', 3504, 235, 15),
	(47325, 'NUSA AGUNG', 3504, 235, 15),
	(47326, 'NUSA BAKTI', 3504, 235, 15),
	(47327, 'NUSA BALI', 3504, 235, 15),
	(47328, 'NUSA JAYA', 3504, 235, 15),
	(47329, 'NUSA MAJU', 3504, 235, 15),
	(47330, 'NUSA TENGGARA', 3504, 235, 15),
	(47331, 'NUSA TUNGGAL', 3504, 235, 15),
	(47332, 'NUSARAYA', 3504, 235, 15),
	(47333, 'RINGIN SARI', 3504, 235, 15),
	(47334, 'SENU MARGA', 3504, 235, 15),
	(47335, 'SINAR BALI', 3504, 235, 15),
	(47336, 'SUKANEGARA', 3504, 235, 15),
	(47337, 'SUMBER REJO', 3504, 235, 15),
	(47338, 'TRI KARYA', 3504, 235, 15),
	(47339, 'ARGO MULYO', 3505, 235, 15),
	(47340, 'BANJAR REJO', 3505, 235, 15),
	(47341, 'GIRI MULYO', 3505, 235, 15),
	(47342, 'KARSA JAYA', 3505, 235, 15),
	(47343, 'KARYA MAKMUR', 3505, 235, 15),
	(47344, 'MADU GONDO', 3505, 235, 15),
	(47345, 'MARGO KOYO', 3505, 235, 15),
	(47346, 'PANCA TUNGGAL', 3505, 235, 15),
	(47347, 'REJOSARI', 3505, 235, 15),
	(47348, 'REJOSARI JAYA', 3505, 235, 15),
	(47349, 'SIDOREJO', 3505, 235, 15),
	(47350, 'SUMBER AGUNG', 3505, 235, 15),
	(47351, 'WINDUSARI', 3505, 235, 15),
	(47352, 'BANGSA NEGARA', 3506, 235, 15),
	(47353, 'JATI MULYO I', 3506, 235, 15),
	(47354, 'KARANG BINANGUN', 3506, 235, 15),
	(47355, 'LUBUK HARJO', 3506, 235, 15),
	(47356, 'MARGA CINTA', 3506, 235, 15),
	(47357, 'MEKAR JAYA', 3506, 235, 15),
	(47358, 'PANDAN SARI I', 3506, 235, 15),
	(47359, 'PELITA JAYA', 3506, 235, 15),
	(47360, 'RANTAU JAYA', 3506, 235, 15),
	(47361, 'TANAH MERAH', 3506, 235, 15),
	(47362, 'TEBING SARI MULYA', 3506, 235, 15),
	(47363, 'TUGU HARUM', 3506, 235, 15),
	(47364, 'TUGU MULYO', 3506, 235, 15),
	(47365, 'TULUS AYU', 3506, 235, 15),
	(47366, 'YOSOWINANGUN', 3506, 235, 15),
	(47367, 'MULYA SARI', 3507, 235, 15),
	(47368, 'PETANGGAN', 3507, 235, 15),
	(47369, 'PURWODADI', 3507, 235, 15),
	(47370, 'REJOSARI', 3507, 235, 15),
	(47371, 'SARI GUNA', 3507, 235, 15),
	(47372, 'SIDOWALUYO', 3507, 235, 15),
	(47373, 'SRI BUDAYA', 3507, 235, 15),
	(47374, 'SRI MULYO', 3507, 235, 15),
	(47375, 'SUGIH WARAS', 3507, 235, 15),
	(47376, 'SUKOHARJO', 3507, 235, 15),
	(47377, 'TULUNG SARI', 3507, 235, 15),
	(47378, 'ULAK BUNTAR', 3507, 235, 15),
	(47379, 'AMAN JAYA', 3508, 235, 15),
	(47380, 'CIPTA MUDA', 3508, 235, 15),
	(47381, 'GANJAR AGUNG', 3508, 235, 15),
	(47382, 'KURUNGAN NYAWA', 3508, 235, 15),
	(47383, 'KURUNGAN NYAWA I', 3508, 235, 15),
	(47384, 'KURUNGAN NYAWA II', 3508, 235, 15),
	(47385, 'KURUNGAN NYAWA III', 3508, 235, 15),
	(47386, 'MUDA SENTOSA', 3508, 235, 15),
	(47387, 'MULYO AGUNG', 3508, 235, 15),
	(47388, 'PISANG JAYA', 3508, 235, 15),
	(47389, 'SRI DADI (SARIDATI)', 3508, 235, 15),
	(47390, 'SUKARAJA', 3508, 235, 15),
	(47391, 'SUKARAJA TUHA', 3508, 235, 15),
	(47392, 'SUMBER AGUNG', 3508, 235, 15),
	(47393, 'TANJUNG BULAN', 3508, 235, 15),
	(47394, 'TEBAT JAYA', 3508, 235, 15),
	(47395, 'WAY HALOM', 3508, 235, 15),
	(47396, 'BANGUN HARJO', 3509, 235, 15),
	(47397, 'BANYUMAS ASRI', 3509, 235, 15),
	(47398, 'KARANG TENGAH', 3509, 235, 15),
	(47399, 'KEDU', 3509, 235, 15),
	(47400, 'KEDUNG REJO', 3509, 235, 15),
	(47401, 'KUMPUL REJO', 3509, 235, 15),
	(47402, 'LIMANSARI', 3509, 235, 15),
	(47403, 'METRO REJO', 3509, 235, 15),
	(47404, 'PENGANDONAN', 3509, 235, 15),
	(47405, 'REJODADI', 3509, 235, 15),
	(47406, 'ROWODADI', 3509, 235, 15),
	(47407, 'SRI KATON', 3509, 235, 15),
	(47408, 'SUKA MAJU', 3509, 235, 15),
	(47409, 'SUKODADI', 3509, 235, 15),
	(47410, 'SUKOHARJO', 3509, 235, 15),
	(47411, 'SUMBER ASRI', 3509, 235, 15),
	(47412, 'SUMBER HARJO', 3509, 235, 15),
	(47413, 'SUMBER MULYO', 3509, 235, 15),
	(47414, 'SUMEDANG SARI', 3509, 235, 15),
	(47415, 'TAMBAK BOYO', 3509, 235, 15),
	(47416, 'TANJUNGMAS', 3509, 235, 15),
	(47417, 'TANJUNGSARI', 3509, 235, 15),
	(47418, 'TEKOREJO (TOKOREJO)', 3509, 235, 15),
	(47419, 'ANYAR', 3510, 235, 15),
	(47420, 'MUNCAK KABAU', 3510, 235, 15),
	(47421, 'PANDAN SARI III', 3510, 235, 15),
	(47422, 'RAWASARI', 3510, 235, 15),
	(47423, 'SRI BULAN', 3510, 235, 15),
	(47424, 'SRI BUNGA', 3510, 235, 15),
	(47425, 'SURYA MENANG', 3510, 235, 15),
	(47426, 'BANDAR JAYA', 3511, 235, 15),
	(47427, 'BANTAN', 3511, 235, 15),
	(47428, 'BANTAN PELITA', 3511, 235, 15),
	(47429, 'BANUAYU', 3511, 235, 15),
	(47430, 'BANUMAS', 3511, 235, 15),
	(47431, 'NEGERI AGUNG', 3511, 235, 15),
	(47432, 'NEGERI AGUNG JAYA', 3511, 235, 15),
	(47433, 'NEGERI PAKUAN', 3511, 235, 15),
	(47434, 'PAHANG ASRI', 3511, 235, 15),
	(47435, 'PEMETUNG BASUKI', 3511, 235, 15),
	(47436, 'PULAU NEGARA', 3511, 235, 15),
	(47437, 'SAUNG DADI', 3511, 235, 15),
	(47438, 'TERANTANG/TRANTANG SAKTI', 3511, 235, 15),
	(47439, 'BUNGA MAYANG', 3512, 235, 15),
	(47440, 'CONDONG', 3512, 235, 15),
	(47441, 'JAYAPURA', 3512, 235, 15),
	(47442, 'KAMBANG', 3512, 235, 15),
	(47443, 'MENDAH', 3512, 235, 15),
	(47444, 'PERACAK JAYA', 3512, 235, 15),
	(47445, 'TUMIJAYA', 3512, 235, 15),
	(47446, 'WAY SALAK', 3512, 235, 15),
	(47447, 'AGUNG JATI', 3513, 235, 15),
	(47448, 'GUNUNG TERANG', 3513, 235, 15),
	(47449, 'HARJO MULYO', 3513, 235, 15),
	(47450, 'JATISARI', 3513, 235, 15),
	(47451, 'JAYA BAKTI', 3513, 235, 15),
	(47452, 'KARTA MULYA', 3513, 235, 15),
	(47453, 'MENDAYUN', 3513, 235, 15),
	(47454, 'MENGULAK', 3513, 235, 15),
	(47455, 'RASUAN', 3513, 235, 15),
	(47456, 'RASUAN DARAT', 3513, 235, 15),
	(47457, 'SIMPANG KARTA MULYA', 3513, 235, 15),
	(47458, 'TRIDADI', 3513, 235, 15),
	(47459, 'BANBAN REJO', 3514, 235, 15),
	(47460, 'DADI MULYO', 3514, 235, 15),
	(47461, 'JATI MULYO II (MULYA II)', 3514, 235, 15),
	(47462, 'KARANG NEGARA', 3514, 235, 15),
	(47463, 'KERTANEGARA', 3514, 235, 15),
	(47464, 'KOTA NEGARA', 3514, 235, 15),
	(47465, 'MARGO TANI', 3514, 235, 15),
	(47466, 'MARGOTANI II', 3514, 235, 15),
	(47467, 'PANDAN AGUNG', 3514, 235, 15),
	(47468, 'PANDAN JAYA', 3514, 235, 15),
	(47469, 'RASUAN BARU', 3514, 235, 15),
	(47470, 'RIANG BANDUNG', 3514, 235, 15),
	(47471, 'RIANG BANDUNG ILIR', 3514, 235, 15),
	(47472, 'SRI KENCANA', 3514, 235, 15),
	(47473, 'SRI MULYO', 3514, 235, 15),
	(47474, 'SUKANEGARA', 3514, 235, 15),
	(47475, 'BANDING AGUNG', 3515, 235, 15),
	(47476, 'BATUMARTA VI', 3515, 235, 15),
	(47477, 'BATUMARTA X', 3515, 235, 15),
	(47478, 'BINA AMARTA', 3515, 235, 15),
	(47479, 'KARYA MAKMUR', 3515, 235, 15),
	(47480, 'NIKAN', 3515, 235, 15),
	(47481, 'SUKA DAMAI', 3515, 235, 15),
	(47482, 'SURABAYA', 3515, 235, 15),
	(47483, 'WANA BHAKTI', 3515, 235, 15),
	(47484, 'BUKIT SARI', 3516, 235, 15),
	(47485, 'DUSUN MARTAPURA', 3516, 235, 15),
	(47486, 'KEROMONGAN', 3516, 235, 15),
	(47487, 'KOTA BARU', 3516, 235, 15),
	(47488, 'KOTA BARU BARAT', 3516, 235, 15),
	(47489, 'KOTA BARU SELATAN', 3516, 235, 15),
	(47490, 'KURUNGAN NYAWA I', 3516, 235, 15),
	(47491, 'PAKU SENGKUNYIT', 3516, 235, 15),
	(47492, 'PASAR MARTAPURA', 3516, 235, 15),
	(47493, 'PERJAYA', 3516, 235, 15),
	(47494, 'SUKOMULYO', 3516, 235, 15),
	(47495, 'SUNGAI TUHA JAYA', 3516, 235, 15),
	(47496, 'TANJUNG KEMALA', 3516, 235, 15),
	(47497, 'TANJUNG KEMALA BARAT', 3516, 235, 15),
	(47498, 'TERUKIS RAHAYU', 3516, 235, 15),
	(47499, 'VETERAN JAYA', 3516, 235, 15),
	(47500, 'ADU MANIS', 3517, 235, 15),
	(47501, 'BETUNG', 3517, 235, 15),
	(47502, 'BETUNG TIMUR', 3517, 235, 15),
	(47503, 'KANGKUNG', 3517, 235, 15),
	(47504, 'KANGKUNG ILIR', 3517, 235, 15),
	(47505, 'MENANGA BESAR', 3517, 235, 15),
	(47506, 'MENANGA SARI', 3517, 235, 15),
	(47507, 'MENANGA TENGAH', 3517, 235, 15),
	(47508, 'SRI TANJUNG', 3517, 235, 15),
	(47509, 'SUKA NEGERI', 3517, 235, 15),
	(47510, 'TANJUNG KUKUH', 3517, 235, 15),
	(47511, 'TANJUNG MAS', 3517, 235, 15),
	(47512, 'CAHAYA/CAHYA NEGERI', 3518, 235, 15),
	(47513, 'GUNUNG SUGIH', 3518, 235, 15),
	(47514, 'JAYA MULIA/MULYA', 3518, 235, 15),
	(47515, 'KARANG ENDAH', 3518, 235, 15),
	(47516, 'KARANG MARGA', 3518, 235, 15),
	(47517, 'KRUJON', 3518, 235, 15),
	(47518, 'MARGODADI', 3518, 235, 15),
	(47519, 'MARGOREJO', 3518, 235, 15),
	(47520, 'MUJO RAHAYU', 3518, 235, 15),
	(47521, 'SRIWANGI', 3518, 235, 15),
	(47522, 'SRIWANGI ULU', 3518, 235, 15),
	(47523, 'SUKA MULYA', 3518, 235, 15),
	(47524, 'TAMAN AGUNG', 3518, 235, 15),
	(47525, 'TAMAN ASRI', 3518, 235, 15),
	(47526, 'TAMAN HARJO', 3518, 235, 15),
	(47527, 'TAMAN MULYO', 3518, 235, 15),
	(47528, 'TARAMAN', 3518, 235, 15),
	(47529, 'TRIMO REJO', 3518, 235, 15),
	(47530, 'TRIMOHARJO', 3518, 235, 15),
	(47531, 'BUNGIN JAYA', 3519, 235, 15),
	(47532, 'BURNAI JAYA', 3519, 235, 15),
	(47533, 'BURNAI MULYA', 3519, 235, 15),
	(47534, 'HARAPAN JAYA', 3519, 235, 15),
	(47535, 'KARANG ANYAR', 3519, 235, 15),
	(47536, 'KARANG MELATI', 3519, 235, 15),
	(47537, 'KARANG MENJANGAN', 3519, 235, 15),
	(47538, 'KARANG MULYA', 3519, 235, 15),
	(47539, 'KOTA MULYA', 3519, 235, 15),
	(47540, 'KOTA TANAH', 3519, 235, 15),
	(47541, 'MELATI AGUNG', 3519, 235, 15),
	(47542, 'MELATI JAYA', 3519, 235, 15),
	(47543, 'MULIA JAYA', 3519, 235, 15),
	(47544, 'NIRWANA', 3519, 235, 15),
	(47545, 'TULUNG HARAPAN', 3519, 235, 15),
	(47546, 'WANA MAKMUR', 3519, 235, 15),
	(47547, 'WARNA SARI (WANASARI)', 3519, 235, 15),
	(47548, 'ATUNG BUNGSU', 3520, 236, 15),
	(47549, 'KANCE DIWE', 3520, 236, 15),
	(47550, 'LUBUK BUNTAK', 3520, 236, 15),
	(47551, 'PENJALANG', 3520, 236, 15),
	(47552, 'PRAHU DIPO', 3520, 236, 15),
	(47553, 'CANDI JAYA', 3521, 236, 15),
	(47554, 'JOKOH', 3521, 236, 15),
	(47555, 'KARANG DALO', 3521, 236, 15),
	(47556, 'PADANG TEMU', 3521, 236, 15),
	(47557, 'PELANG KENIDAI', 3521, 236, 15),
	(47558, 'AGUNG LAWANGAN', 3522, 236, 15),
	(47559, 'BUMI AGUNG', 3522, 236, 15),
	(47560, 'BURUNG DINANG', 3522, 236, 15),
	(47561, 'JANGKAR MAS', 3522, 236, 15),
	(47562, 'MUARA SIBAN', 3522, 236, 15),
	(47563, 'PAGAR WANGI', 3522, 236, 15),
	(47564, 'REBA TINGGI', 3522, 236, 15),
	(47565, 'BESEMAH SERASAN', 3523, 236, 15),
	(47566, 'GUNUNG DEMPO', 3523, 236, 15),
	(47567, 'NENDAGUNG', 3523, 236, 15),
	(47568, 'SIDOREJO', 3523, 236, 15),
	(47569, 'TANJUNG PAYANG', 3523, 236, 15),
	(47570, 'TEBAT GIRI INDAH', 3523, 236, 15),
	(47571, 'TUMBAK ULAS', 3523, 236, 15),
	(47572, 'ULU LURAH/RURAH', 3523, 236, 15),
	(47573, 'ALUN DUA', 3524, 236, 15),
	(47574, 'BANGUN JAYA', 3524, 236, 15),
	(47575, 'BANGUN REJO', 3524, 236, 15),
	(47576, 'BERINGIN JAYA', 3524, 236, 15),
	(47577, 'CURUP JARE', 3524, 236, 15),
	(47578, 'DEMPO MAKMUR', 3524, 236, 15),
	(47579, 'KURIPAN BABAS', 3524, 236, 15),
	(47580, 'PAGAR ALAM', 3524, 236, 15),
	(47581, 'SELIBAR', 3524, 236, 15),
	(47582, 'SUKOREJO', 3524, 236, 15),
	(47583, 'ALANG-ALANG LEBAR', 3525, 237, 15),
	(47584, 'KARYA BARU', 3525, 237, 15),
	(47585, 'SRIJAYA', 3525, 237, 15),
	(47586, 'TALANG KELAPA', 3525, 237, 15),
	(47587, '19 ILIR', 3526, 237, 15),
	(47588, '22 ILIR', 3526, 237, 15),
	(47589, '23 ILIR', 3526, 237, 15),
	(47590, '24 ILIR', 3526, 237, 15),
	(47591, '26 ILIR', 3526, 237, 15),
	(47592, 'TALANG SEMUT', 3526, 237, 15),
	(47593, '36 ILIR', 3527, 237, 15),
	(47594, 'GANDUS', 3527, 237, 15),
	(47595, 'KARANG ANYAR', 3527, 237, 15),
	(47596, 'KARANG JAYA', 3527, 237, 15),
	(47597, 'PULO KERTO', 3527, 237, 15),
	(47598, '26 ILIR D. I', 3528, 237, 15),
	(47599, 'BUKIT BARU', 3528, 237, 15),
	(47600, 'BUKIT LAMA', 3528, 237, 15),
	(47601, 'DEMANG LEBAR DAUN', 3528, 237, 15),
	(47602, 'LOROK PAKJO', 3528, 237, 15),
	(47603, 'SIRING AGUNG', 3528, 237, 15),
	(47604, '27 ILIR', 3529, 237, 15),
	(47605, '28 ILIR', 3529, 237, 15),
	(47606, '29 ILIR', 3529, 237, 15),
	(47607, '30 ILIR', 3529, 237, 15),
	(47608, '32 ILIR', 3529, 237, 15),
	(47609, '35 ILIR', 3529, 237, 15),
	(47610, 'KEMANG MANIS', 3529, 237, 15),
	(47611, '13 ILIR', 3530, 237, 15),
	(47612, '14 ILIR', 3530, 237, 15),
	(47613, '15 ILIR', 3530, 237, 15),
	(47614, '16 ILIR', 3530, 237, 15),
	(47615, '17 ILIR', 3530, 237, 15),
	(47616, '18 ILIR', 3530, 237, 15),
	(47617, '20 ILIR I', 3530, 237, 15),
	(47618, '20 ILIR III', 3530, 237, 15),
	(47619, '20 ILIR IV', 3530, 237, 15),
	(47620, 'KEPANDEAN BARU', 3530, 237, 15),
	(47621, 'SEI PANGERAN', 3530, 237, 15),
	(47622, '1 ILIR', 3531, 237, 15),
	(47623, '10 ILIR', 3531, 237, 15),
	(47624, '11 ILIR', 3531, 237, 15),
	(47625, '2 ILIR', 3531, 237, 15),
	(47626, '3 ILIR', 3531, 237, 15),
	(47627, '5 ILIR', 3531, 237, 15),
	(47628, '8 ILIR', 3531, 237, 15),
	(47629, '9 ILIR', 3531, 237, 15),
	(47630, 'DUKU', 3531, 237, 15),
	(47631, 'KUTO BATU', 3531, 237, 15),
	(47632, 'LAWANG KIDUL', 3531, 237, 15),
	(47633, 'SUNGAI BUAH', 3531, 237, 15),
	(47634, 'BUKIT SANGKAL', 3532, 237, 15),
	(47635, 'KALIDONI', 3532, 237, 15),
	(47636, 'SEI LAIS', 3532, 237, 15),
	(47637, 'SEI SELAYUR', 3532, 237, 15),
	(47638, 'SEI SELINCAH', 3532, 237, 15),
	(47639, '20 ILIR II', 3533, 237, 15),
	(47640, 'AIR BALUI', 3533, 237, 15),
	(47641, 'ARIO KEMUNING', 3533, 237, 15),
	(47642, 'BATU AMPAR', 3533, 237, 15),
	(47643, 'DUSUN TUK JIMUN', 3533, 237, 15),
	(47644, 'KEMUNING MUDA', 3533, 237, 15),
	(47645, 'KEMUNING TUA', 3533, 237, 15),
	(47646, 'KERITANG', 3533, 237, 15),
	(47647, 'LIMAU MANIS', 3533, 237, 15),
	(47648, 'LUBUK BESAR', 3533, 237, 15),
	(47649, 'PAHLAWAN', 3533, 237, 15),
	(47650, 'PIPA REJA', 3533, 237, 15),
	(47651, 'SEKARA', 3533, 237, 15),
	(47652, 'SEKAYAN', 3533, 237, 15),
	(47653, 'SEKIP JAYA', 3533, 237, 15),
	(47654, 'SELENSEN', 3533, 237, 15),
	(47655, 'TALANG AMAN', 3533, 237, 15),
	(47656, 'TALANG JANGKANG', 3533, 237, 15),
	(47657, 'KARYA JAYA', 3534, 237, 15),
	(47658, 'KEMANG AGUNG', 3534, 237, 15),
	(47659, 'KEMAS RINDO', 3534, 237, 15),
	(47660, 'KERAMASAN', 3534, 237, 15),
	(47661, 'KERTAPATI', 3534, 237, 15),
	(47662, 'OGAN BARU', 3534, 237, 15),
	(47663, 'BAGUS KUNING', 3535, 237, 15),
	(47664, 'KOMPERTA', 3535, 237, 15),
	(47665, 'PLAJU DARAT', 3535, 237, 15),
	(47666, 'PLAJU ILIR', 3535, 237, 15),
	(47667, 'PLAJU ULU', 3535, 237, 15),
	(47668, 'TALANG BUBUK', 3535, 237, 15),
	(47669, 'TALANG PUTRI', 3535, 237, 15),
	(47670, 'SAKO', 3536, 237, 15),
	(47671, 'SIALANG', 3536, 237, 15),
	(47672, 'SUKAMAJU', 3536, 237, 15),
	(47673, '1 ULU', 3537, 237, 15),
	(47674, '15 ULU', 3537, 237, 15),
	(47675, '2 ULU', 3537, 237, 15),
	(47676, '3-4 ULU', 3537, 237, 15),
	(47677, '5 ULU', 3537, 237, 15),
	(47678, '7 ULU', 3537, 237, 15),
	(47679, '8 ULU', 3537, 237, 15),
	(47680, '9/10 ULU', 3537, 237, 15),
	(47681, 'SILABERANTI', 3537, 237, 15),
	(47682, 'TUAN KENTANG', 3537, 237, 15),
	(47683, '11 ULU', 3538, 237, 15),
	(47684, '12 ULU', 3538, 237, 15),
	(47685, '13 ULU', 3538, 237, 15),
	(47686, '14 ULU', 3538, 237, 15),
	(47687, '16 ULU', 3538, 237, 15),
	(47688, 'SENTOSA', 3538, 237, 15),
	(47689, 'TANGGA TAKAT', 3538, 237, 15),
	(47690, 'KARYA MULYA', 3539, 237, 15),
	(47691, 'LEBONG/LEBUNG GAJAH', 3539, 237, 15),
	(47692, 'SAKO BARU', 3539, 237, 15),
	(47693, 'SRI MULYA (SRIMULYO)', 3539, 237, 15),
	(47694, 'SUKA MULYA', 3539, 237, 15),
	(47695, 'JAMBE (TALANG JAMBE)', 3540, 237, 15),
	(47696, 'KEBUN BUNGA', 3540, 237, 15),
	(47697, 'SUKA BANGUN', 3540, 237, 15),
	(47698, 'SUKAJAYA', 3540, 237, 15),
	(47699, 'SUKARAMI', 3540, 237, 15),
	(47700, 'SUKODADI', 3540, 237, 15),
	(47701, 'TALANG BETUTU', 3540, 237, 15),
	(47702, 'CAMBAI', 3541, 238, 15),
	(47703, 'MUARA SUNGAI', 3541, 238, 15),
	(47704, 'PANGKUL', 3541, 238, 15),
	(47705, 'SINDUR', 3541, 238, 15),
	(47706, 'SUNGAI MEDANG', 3541, 238, 15),
	(47707, 'GUNUNG KEMALA', 3542, 238, 15),
	(47708, 'MUNTANG TAPUS', 3542, 238, 15),
	(47709, 'PATIH GALUNG', 3542, 238, 15),
	(47710, 'PAYU PUTAT', 3542, 238, 15),
	(47711, 'PRABUMULIH', 3542, 238, 15),
	(47712, 'TANJUNG TELANG', 3542, 238, 15),
	(47713, 'MAJASARI', 3543, 238, 15),
	(47714, 'SUKARAJA', 3543, 238, 15),
	(47715, 'TANJUNG MENANG', 3543, 238, 15),
	(47716, 'TANJUNG RAMAN', 3543, 238, 15),
	(47717, 'GUNUNG IBUL', 3544, 238, 15),
	(47718, 'GUNUNG IBUL BARAT', 3544, 238, 15),
	(47719, 'KARANG JAYA', 3544, 238, 15),
	(47720, 'KARANG RAJA', 3544, 238, 15),
	(47721, 'MUARA DUA', 3544, 238, 15),
	(47722, 'PRABU JAYA', 3544, 238, 15),
	(47723, 'SUKAJADI', 3544, 238, 15),
	(47724, 'TUGU KECIL', 3544, 238, 15),
	(47725, 'ANAK PETAI', 3545, 238, 15),
	(47726, 'MANGGA BESAR', 3545, 238, 15),
	(47727, 'PASAR PRABUMULIH', 3545, 238, 15),
	(47728, 'PASAR PRABUMULIH II', 3545, 238, 15),
	(47729, 'WONOSARI', 3545, 238, 15),
	(47730, 'JUNGAI', 3546, 238, 15),
	(47731, 'KARANG BINDU', 3546, 238, 15),
	(47732, 'KARANGAN', 3546, 238, 15),
	(47733, 'KARYA MULYA', 3546, 238, 15),
	(47734, 'KEMBANG TANDUK', 3546, 238, 15),
	(47735, 'RAMBANG SENULING', 3546, 238, 15),
	(47736, 'SINAR RAMBANG', 3546, 238, 15),
	(47737, 'TALANG BATU', 3546, 238, 15),
	(47738, 'TANJUNG RAMBANG', 3546, 238, 15),
	(47739, 'BAKAM', 3547, 239, 16),
	(47740, 'BUKIT LAYANG', 3547, 239, 16),
	(47741, 'DALIL', 3547, 239, 16),
	(47742, 'KAPUK', 3547, 239, 16),
	(47743, 'MABAT', 3547, 239, 16),
	(47744, 'MANGKA', 3547, 239, 16),
	(47745, 'MARAS SENANG', 3547, 239, 16),
	(47746, 'NEKNANG', 3547, 239, 16),
	(47747, 'TIANG TARA', 3547, 239, 16),
	(47748, 'AIR JUKUNG', 3548, 239, 16),
	(47749, 'BINTET', 3548, 239, 16),
	(47750, 'BUKIT KETOK', 3548, 239, 16),
	(47751, 'GUNUNG MUDA', 3548, 239, 16),
	(47752, 'GUNUNG PELAWAN', 3548, 239, 16),
	(47753, 'KUTO PANJI', 3548, 239, 16),
	(47754, 'LUMUT', 3548, 239, 16),
	(47755, 'RIDING PANJANG', 3548, 239, 16),
	(47756, 'AIR BULUH', 3549, 239, 16),
	(47757, 'AIR DUREN', 3549, 239, 16),
	(47758, 'CENGKONG ABANG', 3549, 239, 16),
	(47759, 'KACE', 3549, 239, 16),
	(47760, 'KACE TIMUR', 3549, 239, 16),
	(47761, 'KEMUJA', 3549, 239, 16),
	(47762, 'KOTA KAPUR', 3549, 239, 16),
	(47763, 'LABUH AIR PANDAN', 3549, 239, 16),
	(47764, 'MENDO (MENDUK)', 3549, 239, 16),
	(47765, 'PAYABENUA', 3549, 239, 16),
	(47766, 'PENAGAN', 3549, 239, 16),
	(47767, 'PETALING', 3549, 239, 16),
	(47768, 'PETALING BANJAR', 3549, 239, 16),
	(47769, 'RUKAM', 3549, 239, 16),
	(47770, 'ZED', 3549, 239, 16),
	(47771, 'AIR ANYIR', 3550, 239, 16),
	(47772, 'BALUN IJUK', 3550, 239, 16),
	(47773, 'BATURUSA', 3550, 239, 16),
	(47774, 'DWI MAKMUR', 3550, 239, 16),
	(47775, 'JADA BAHRIN', 3550, 239, 16),
	(47776, 'JURUNG', 3550, 239, 16),
	(47777, 'KIMAK', 3550, 239, 16),
	(47778, 'MERAWANG', 3550, 239, 16),
	(47779, 'PAGARAWAN', 3550, 239, 16),
	(47780, 'RIDING PANJANG', 3550, 239, 16),
	(47781, 'AIR DUREN', 3551, 239, 16),
	(47782, 'AIR RUAI', 3551, 239, 16),
	(47783, 'KARYA MAKMUR', 3551, 239, 16),
	(47784, 'PEMALI', 3551, 239, 16),
	(47785, 'PENYAMUN', 3551, 239, 16),
	(47786, 'SEMPAN', 3551, 239, 16),
	(47787, 'KAYU BESI', 3552, 239, 16),
	(47788, 'KOTA WARINGIN', 3552, 239, 16),
	(47789, 'LABU', 3552, 239, 16),
	(47790, 'NIBUNG', 3552, 239, 16),
	(47791, 'PUDINGBESAR', 3552, 239, 16),
	(47792, 'SAING', 3552, 239, 16),
	(47793, 'TANAH BAWAH', 3552, 239, 16),
	(47794, 'BANYU ASIN', 3553, 239, 16),
	(47795, 'BERBURA', 3553, 239, 16),
	(47796, 'CIT', 3553, 239, 16),
	(47797, 'DENIANG', 3553, 239, 16),
	(47798, 'MAPUR', 3553, 239, 16),
	(47799, 'PANGKAL NIUR', 3553, 239, 16),
	(47800, 'PUGUL', 3553, 239, 16),
	(47801, 'RIAU', 3553, 239, 16),
	(47802, 'SILIP', 3553, 239, 16),
	(47803, 'KENANGA', 3554, 239, 16),
	(47804, 'KUDAI (KUDAY)', 3554, 239, 16),
	(47805, 'PARIT PADANG', 3554, 239, 16),
	(47806, 'REBO', 3554, 239, 16),
	(47807, 'SINAR BARU', 3554, 239, 16),
	(47808, 'SRIMENANTI', 3554, 239, 16),
	(47809, 'SUNGAILIAT', 3554, 239, 16),
	(47810, 'AIR KUANG', 3555, 240, 16),
	(47811, 'JEBUS', 3555, 240, 16),
	(47812, 'KETAP', 3555, 240, 16),
	(47813, 'LIMBUNG', 3555, 240, 16),
	(47814, 'MISLAK', 3555, 240, 16),
	(47815, 'PEBUAR', 3555, 240, 16),
	(47816, 'RANGGI ASAM', 3555, 240, 16),
	(47817, 'RUKAM', 3555, 240, 16),
	(47818, 'SINAR MANIK', 3555, 240, 16),
	(47819, 'SUNGAI BULUH', 3555, 240, 16),
	(47820, 'TUMBAK PETAR', 3555, 240, 16),
	(47821, 'AIR BULIN', 3556, 240, 16),
	(47822, 'BERUAS', 3556, 240, 16),
	(47823, 'DENDANG', 3556, 240, 16),
	(47824, 'KACUNG', 3556, 240, 16),
	(47825, 'KAYUARANG', 3556, 240, 16),
	(47826, 'KELAPA', 3556, 240, 16),
	(47827, 'MANCUNG', 3556, 240, 16),
	(47828, 'PANGKAL BERAS', 3556, 240, 16),
	(47829, 'PUSUK', 3556, 240, 16),
	(47830, 'SINAR SARI', 3556, 240, 16),
	(47831, 'TEBING', 3556, 240, 16),
	(47832, 'TERENTANG', 3556, 240, 16),
	(47833, 'TUGANG', 3556, 240, 16),
	(47834, 'TUIK', 3556, 240, 16),
	(47835, 'AIR BELO', 3557, 240, 16),
	(47836, 'AIR LIMAU', 3557, 240, 16),
	(47837, 'AIR PUTIH', 3557, 240, 16),
	(47838, 'BELOLAUT', 3557, 240, 16),
	(47839, 'SUNGAI BARU', 3557, 240, 16),
	(47840, 'SUNGAI DAENG', 3557, 240, 16),
	(47841, 'TANJUNG', 3557, 240, 16),
	(47842, 'AIR GANTANG', 3558, 240, 16),
	(47843, 'BAKIT', 3558, 240, 16),
	(47844, 'CUPAT', 3558, 240, 16),
	(47845, 'KAPIT', 3558, 240, 16),
	(47846, 'KELABAT', 3558, 240, 16),
	(47847, 'PUPUT', 3558, 240, 16),
	(47848, 'SEKAR BIRU', 3558, 240, 16),
	(47849, 'SEMULUT', 3558, 240, 16),
	(47850, 'TELAK', 3558, 240, 16),
	(47851, 'TELUK LIMAU', 3558, 240, 16),
	(47852, 'AIR MENDUYUNG', 3559, 240, 16),
	(47853, 'AIR NYATOH', 3559, 240, 16),
	(47854, 'BERANG', 3559, 240, 16),
	(47855, 'BUKIT TERAK', 3559, 240, 16),
	(47856, 'IBUL', 3559, 240, 16),
	(47857, 'KUNDI', 3559, 240, 16),
	(47858, 'MAYANG', 3559, 240, 16),
	(47859, 'PANGEK', 3559, 240, 16),
	(47860, 'PELANGAS', 3559, 240, 16),
	(47861, 'PERADONG', 3559, 240, 16),
	(47862, 'RAMBAT', 3559, 240, 16),
	(47863, 'SIMPANG GUNG', 3559, 240, 16),
	(47864, 'SIMPANG TIGA', 3559, 240, 16),
	(47865, 'AIR LINTANG', 3560, 240, 16),
	(47866, 'BENTENG KOTA/KUTA', 3560, 240, 16),
	(47867, 'BUYAN KELUMBI', 3560, 240, 16),
	(47868, 'PENYAMPAK', 3560, 240, 16),
	(47869, 'SANGKU', 3560, 240, 16),
	(47870, 'SIMPANG YUL', 3560, 240, 16),
	(47871, 'SINAR SURYA', 3560, 240, 16),
	(47872, 'TANJUNGNIUR', 3560, 240, 16),
	(47873, 'TEMPILANG', 3560, 240, 16),
	(47874, 'AIR BARA', 3561, 241, 16),
	(47875, 'AIR GEGAS', 3561, 241, 16),
	(47876, 'BENCAH', 3561, 241, 16),
	(47877, 'DELAS', 3561, 241, 16),
	(47878, 'NANGKA', 3561, 241, 16),
	(47879, 'NYELANDING', 3561, 241, 16),
	(47880, 'PERGAM', 3561, 241, 16),
	(47881, 'RANGGAS', 3561, 241, 16),
	(47882, 'SIDOHARJO', 3561, 241, 16),
	(47883, 'TEPUS', 3561, 241, 16),
	(47884, 'CELAGEN', 3562, 241, 16),
	(47885, 'PONGOK', 3562, 241, 16),
	(47886, 'KUMBUNG', 3563, 241, 16),
	(47887, 'PENUTUK', 3563, 241, 16),
	(47888, 'TANJUNG LABU', 3563, 241, 16),
	(47889, 'TANJUNG SANGKAR', 3563, 241, 16),
	(47890, 'BATUKARANG', 3564, 241, 16),
	(47891, 'BEDENGUNG', 3564, 241, 16),
	(47892, 'CIMBANG', 3564, 241, 16),
	(47893, 'GURUKINAYAN', 3564, 241, 16),
	(47894, 'IRAT', 3564, 241, 16),
	(47895, 'MALIK', 3564, 241, 16),
	(47896, 'NADUNG', 3564, 241, 16),
	(47897, 'PAKU', 3564, 241, 16),
	(47898, 'PANGKALBULUH', 3564, 241, 16),
	(47899, 'PAYUNG', 3564, 241, 16),
	(47900, 'PAYUNG', 3564, 241, 16),
	(47901, 'RANGGUNG', 3564, 241, 16),
	(47902, 'RIMOKAYU', 3564, 241, 16),
	(47903, 'SELANDI', 3564, 241, 16),
	(47904, 'SENGIR', 3564, 241, 16),
	(47905, 'SUKA MERIAH', 3564, 241, 16),
	(47906, 'UJUNG PAYUNG', 3564, 241, 16),
	(47907, 'BATU BETUMPANG', 3565, 241, 16),
	(47908, 'FAJAR INDAH (TRANS 1)', 3565, 241, 16),
	(47909, 'PANCA TUNGGAL (TRANS 3)', 3565, 241, 16),
	(47910, 'SUKAJAYA', 3565, 241, 16),
	(47911, 'SUMBER JAYA PERMAI (TRANS 4)', 3565, 241, 16),
	(47912, 'BANGKA KOTA', 3566, 241, 16),
	(47913, 'GUDANG', 3566, 241, 16),
	(47914, 'JELUTUNG II', 3566, 241, 16),
	(47915, 'PERMIS', 3566, 241, 16),
	(47916, 'RAJIK', 3566, 241, 16),
	(47917, 'SEBAGIN', 3566, 241, 16),
	(47918, 'SIMPANG RIMBA', 3566, 241, 16),
	(47919, 'BIKANG', 3567, 241, 16),
	(47920, 'GADUNG', 3567, 241, 16),
	(47921, 'JERIJI', 3567, 241, 16),
	(47922, 'KAPOSANG/KEPOSANG', 3567, 241, 16),
	(47923, 'KEPOH', 3567, 241, 16),
	(47924, 'RIAS', 3567, 241, 16),
	(47925, 'RINDIK', 3567, 241, 16),
	(47926, 'SERDANG', 3567, 241, 16),
	(47927, 'TANJUNG KETAPANG', 3567, 241, 16),
	(47928, 'TELADAN', 3567, 241, 16),
	(47929, 'TOBOALI', 3567, 241, 16),
	(47930, 'BUKIT TERAP', 3568, 241, 16),
	(47931, 'PASIR PUTIH', 3568, 241, 16),
	(47932, 'SADAI', 3568, 241, 16),
	(47933, 'TIRAM', 3568, 241, 16),
	(47934, 'TUKAK', 3568, 241, 16),
	(47935, 'ARUNG DALAM', 3569, 242, 16),
	(47936, 'BEROK', 3569, 242, 16),
	(47937, 'GUNTUNG', 3569, 242, 16),
	(47938, 'KOBA', 3569, 242, 16),
	(47939, 'KURAU (TIMUR)', 3569, 242, 16),
	(47940, 'KURAU BARAT', 3569, 242, 16),
	(47941, 'NIBUNG', 3569, 242, 16),
	(47942, 'PADANG MULYA', 3569, 242, 16),
	(47943, 'PENYAK', 3569, 242, 16),
	(47944, 'SIMPANG PERLANG', 3569, 242, 16),
	(47945, 'TERENTANG III', 3569, 242, 16),
	(47946, 'BATU BERIGA', 3570, 242, 16),
	(47947, 'BELIMBING', 3570, 242, 16),
	(47948, 'KULUR', 3570, 242, 16),
	(47949, 'KULUR HILIR', 3570, 242, 16),
	(47950, 'LUBUK BESAR', 3570, 242, 16),
	(47951, 'LUBUK LINGKUK', 3570, 242, 16),
	(47952, 'LUBUK PABRIK', 3570, 242, 16),
	(47953, 'PERLANG', 3570, 242, 16),
	(47954, 'TRUBUS', 3570, 242, 16),
	(47955, 'BASKARA BAKTI/BHAKTI', 3571, 242, 16),
	(47956, 'BELILIK', 3571, 242, 16),
	(47957, 'BUKIT KIJANG', 3571, 242, 16),
	(47958, 'CAMBAI', 3571, 242, 16),
	(47959, 'CAMBAI SELATAN', 3571, 242, 16),
	(47960, 'JELUTUNG', 3571, 242, 16),
	(47961, 'KAYU BESI', 3571, 242, 16),
	(47962, 'NAMANG', 3571, 242, 16),
	(47963, 'AIR MESU TIMUR', 3572, 242, 16),
	(47964, 'AIRMESU', 3572, 242, 16),
	(47965, 'BATU BELUBANG', 3572, 242, 16),
	(47966, 'BELULUK', 3572, 242, 16),
	(47967, 'BENTENG', 3572, 242, 16),
	(47968, 'DUL', 3572, 242, 16),
	(47969, 'JERUK', 3572, 242, 16),
	(47970, 'KEBINTIK', 3572, 242, 16),
	(47971, 'MANGKOL', 3572, 242, 16),
	(47972, 'PADANG BARU', 3572, 242, 16),
	(47973, 'PELINDANG (SELINDUNG)', 3572, 242, 16),
	(47974, 'TANJUNG GUNUNG', 3572, 242, 16),
	(47975, 'BERUAS', 3573, 242, 16),
	(47976, 'CELUAK', 3573, 242, 16),
	(47977, 'KATIS', 3573, 242, 16),
	(47978, 'PASIR GARAM', 3573, 242, 16),
	(47979, 'PINANG SEBATANG', 3573, 242, 16),
	(47980, 'PUPUT', 3573, 242, 16),
	(47981, 'SIMPANG KATIS', 3573, 242, 16),
	(47982, 'SUNGKAP', 3573, 242, 16),
	(47983, 'TERAK', 3573, 242, 16),
	(47984, 'TERU', 3573, 242, 16),
	(47985, 'KEMINGKING', 3574, 242, 16),
	(47986, 'KERAKAS', 3574, 242, 16),
	(47987, 'KERANTAI', 3574, 242, 16),
	(47988, 'KERETAK', 3574, 242, 16),
	(47989, 'KERETAK ATAS', 3574, 242, 16),
	(47990, 'LAMPUR', 3574, 242, 16),
	(47991, 'MELABUN', 3574, 242, 16),
	(47992, 'MUNGGU', 3574, 242, 16),
	(47993, 'RAMADHON (ROMADHONI)', 3574, 242, 16),
	(47994, 'SARANGMANDI', 3574, 242, 16),
	(47995, 'SUNGAI SELAN', 3574, 242, 16),
	(47996, 'SUNGAI SELAN ATAS', 3574, 242, 16),
	(47997, 'TANJUNG PURA', 3574, 242, 16),
	(47998, 'BANTAN', 3575, 243, 16),
	(47999, 'GUNUNG RITING', 3575, 243, 16),
	(48000, 'KEMBIRI', 3575, 243, 16)
	`)
}
