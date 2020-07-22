package kel50

import "github.com/danangkonang/rest-api/config"

func Kel52() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
	(52001, 'SEMELAGI BESAR', 3923, 269, 18),
	(52002, 'TWI MENTIBAR', 3923, 269, 18),
	(52003, 'BUDUK SEMPADANG', 3924, 269, 18),
	(52004, 'GELIK', 3924, 269, 18),
	(52005, 'SELAKAU TUA', 3924, 269, 18),
	(52006, 'SERANGGAM', 3924, 269, 18),
	(52007, 'SEBURING', 3925, 269, 18),
	(52008, 'SEMPARUK', 3925, 269, 18),
	(52009, 'SEPADU', 3925, 269, 18),
	(52010, 'SEPINGGAN', 3925, 269, 18),
	(52011, 'SINGARAYA', 3925, 269, 18),
	(52012, 'ADINUSO', 3926, 269, 18),
	(52013, 'BALAI GEMURUH', 3926, 269, 18),
	(52014, 'BUKIT MULYA', 3926, 269, 18),
	(52015, 'CLAPAR', 3926, 269, 18),
	(52016, 'CLUWUK', 3926, 269, 18),
	(52017, 'DURENOMBO', 3926, 269, 18),
	(52018, 'GONDANG', 3926, 269, 18),
	(52019, 'JATISARI', 3926, 269, 18),
	(52020, 'JOLOSEKTI', 3926, 269, 18),
	(52021, 'JRAKAHPAYUNG', 3926, 269, 18),
	(52022, 'KALIMANGGIS', 3926, 269, 18),
	(52023, 'KARANGTENGAH', 3926, 269, 18),
	(52024, 'KEBORANGAN', 3926, 269, 18),
	(52025, 'KEBUMEN', 3926, 269, 18),
	(52026, 'KEMIRI BARAT', 3926, 269, 18),
	(52027, 'KEMIRI TIMUR', 3926, 269, 18),
	(52028, 'KUMEJING', 3926, 269, 18),
	(52029, 'KURIPAN', 3926, 269, 18),
	(52030, 'MADAK', 3926, 269, 18),
	(52031, 'MANGUNHARJO', 3926, 269, 18),
	(52032, 'MENJANGAN', 3926, 269, 18),
	(52033, 'MENSADE', 3926, 269, 18),
	(52034, 'MUKTI RAHARJA', 3926, 269, 18),
	(52035, 'SABUNG', 3926, 269, 18),
	(52036, 'SEI/SUNGAI DEDEN', 3926, 269, 18),
	(52037, 'SEI/SUNGAI SAPAK', 3926, 269, 18),
	(52038, 'SEMPURNA', 3926, 269, 18),
	(52039, 'SENGON', 3926, 269, 18),
	(52040, 'SIBERUK', 3926, 269, 18),
	(52041, 'SUBAH', 3926, 269, 18),
	(52042, 'TEBUAH ELOK', 3926, 269, 18),
	(52043, 'TENGGULANGHARJO', 3926, 269, 18),
	(52044, 'ARUNG PARAK', 3927, 269, 18),
	(52045, 'MERABUAN', 3927, 269, 18),
	(52046, 'MERPATI', 3927, 269, 18),
	(52047, 'PANCUR', 3927, 269, 18),
	(52048, 'SEMATA', 3927, 269, 18),
	(52049, 'SIMPANG EMPAT', 3927, 269, 18),
	(52050, 'TANGARAN', 3927, 269, 18),
	(52051, 'BATU MAKJAGE', 3928, 269, 18),
	(52052, 'BEKUT', 3928, 269, 18),
	(52053, 'BUKIT SEGOLER', 3928, 269, 18),
	(52054, 'DUNGUN PERAPAKAN', 3928, 269, 18),
	(52055, 'MAK RAMPAI', 3928, 269, 18),
	(52056, 'MAK TANGGUK', 3928, 269, 18),
	(52057, 'MARIBAS', 3928, 269, 18),
	(52058, 'MATANG LABUNG', 3928, 269, 18),
	(52059, 'MEKAR SEKUNTUM', 3928, 269, 18),
	(52060, 'MENSERE', 3928, 269, 18),
	(52061, 'PANGKALAN KONGSI', 3928, 269, 18),
	(52062, 'PUSAKA', 3928, 269, 18),
	(52063, 'SEBERKAT', 3928, 269, 18),
	(52064, 'SEGARAU PARIT', 3928, 269, 18),
	(52065, 'SEGEDONG', 3928, 269, 18),
	(52066, 'SEI/SUNGAI KELAMBU', 3928, 269, 18),
	(52067, 'SEJIRAM', 3928, 269, 18),
	(52068, 'SEMPALAI', 3928, 269, 18),
	(52069, 'SERET AYON', 3928, 269, 18),
	(52070, 'SERINDANG', 3928, 269, 18),
	(52071, 'SERUMPUN BULUH', 3928, 269, 18),
	(52072, 'TEBAS KUALA', 3928, 269, 18),
	(52073, 'TEBAS SUNGAI', 3928, 269, 18),
	(52074, 'CEPALA', 3929, 269, 18),
	(52075, 'MATANG SEGARAU', 3929, 269, 18),
	(52076, 'MERUBONG', 3929, 269, 18),
	(52077, 'RAMBAYAN', 3929, 269, 18),
	(52078, 'SARI MAKMUR', 3929, 269, 18),
	(52079, 'SEMPADIAN', 3929, 269, 18),
	(52080, 'TEKARANG', 3929, 269, 18),
	(52081, 'BERLIMANG', 3930, 269, 18),
	(52082, 'KUALA PANGKALAN KERAMAT', 3930, 269, 18),
	(52083, 'KUBANGGA', 3930, 269, 18),
	(52084, 'LELA', 3930, 269, 18),
	(52085, 'MATANG SEGANTAR', 3930, 269, 18),
	(52086, 'MEKAR SEKUNTUM', 3930, 269, 18),
	(52087, 'MULIA', 3930, 269, 18),
	(52088, 'PEDADA', 3930, 269, 18),
	(52089, 'PIPIT TEJA', 3930, 269, 18),
	(52090, 'PURINGAN', 3930, 269, 18),
	(52091, 'SAMUSTIDA', 3930, 269, 18),
	(52092, 'SAYANG SEDAYU', 3930, 269, 18),
	(52093, 'SEBAGU', 3930, 269, 18),
	(52094, 'SEI/SUNGAI BARU', 3930, 269, 18),
	(52095, 'SEI/SUNGAI KUMPAI', 3930, 269, 18),
	(52096, 'SEI/SUNGAI SERABEK', 3930, 269, 18),
	(52097, 'SEKURA', 3930, 269, 18),
	(52098, 'SENGAWANG', 3930, 269, 18),
	(52099, 'SEPADU', 3930, 269, 18),
	(52100, 'TAMBATAN', 3930, 269, 18),
	(52101, 'TANJUNG KERACUT', 3930, 269, 18),
	(52102, 'TELUK KASEH', 3930, 269, 18),
	(52103, 'TELUK KEMBANG', 3930, 269, 18),
	(52104, 'TRI MANDAYAN', 3930, 269, 18),
	(52105, 'BULU BALA', 3931, 270, 18),
	(52106, 'COWET (CUWET)', 3931, 270, 18),
	(52107, 'EMPIRANG UJUNG', 3931, 270, 18),
	(52108, 'HILIR', 3931, 270, 18),
	(52109, 'KAWING (MAK KAWING)', 3931, 270, 18),
	(52110, 'KEBADU', 3931, 270, 18),
	(52111, 'PADI KAYE', 3931, 270, 18),
	(52112, 'SEMONCOL', 3931, 270, 18),
	(52113, 'SENYABANG', 3931, 270, 18),
	(52114, 'TAE', 3931, 270, 18),
	(52115, 'TEMIANG MALI', 3931, 270, 18),
	(52116, 'TEMIANG TABA', 3931, 270, 18),
	(52117, 'BERENG BERKAWAT', 3932, 270, 18),
	(52118, 'KASROMEGO', 3932, 270, 18),
	(52119, 'MAWANG MUDA', 3932, 270, 18),
	(52120, 'SEI/SUNGAI ILAI', 3932, 270, 18),
	(52121, 'THANG RAYA', 3932, 270, 18),
	(52122, 'BAHTA', 3933, 270, 18),
	(52123, 'BANTAI', 3933, 270, 18),
	(52124, 'BONTI', 3933, 270, 18),
	(52125, 'EMPODIS', 3933, 270, 18),
	(52126, 'KAMPUH', 3933, 270, 18),
	(52127, 'MAJEL', 3933, 270, 18),
	(52128, 'SAMI', 3933, 270, 18),
	(52129, 'TUNGGUL BOYOK', 3933, 270, 18),
	(52130, 'UPE', 3933, 270, 18),
	(52131, 'ENTIKONG', 3934, 270, 18),
	(52132, 'NEKAN', 3934, 270, 18),
	(52133, 'PALA ASANG (PALAPASANG)', 3934, 270, 18),
	(52134, 'SEMANGIT (SEMANGET)', 3934, 270, 18),
	(52135, 'SURUH TEMBAWANG', 3934, 270, 18),
	(52136, 'BALAI SEBUT', 3935, 270, 18),
	(52137, 'EMPIYANG', 3935, 270, 18),
	(52138, 'JANGKANG BENUA', 3935, 270, 18),
	(52139, 'KETORI', 3935, 270, 18),
	(52140, 'PISANG', 3935, 270, 18),
	(52141, 'SAPE', 3935, 270, 18),
	(52142, 'SELAMPUNG', 3935, 270, 18),
	(52143, 'SEMIRAU', 3935, 270, 18),
	(52144, 'SEMOMBAT', 3935, 270, 18),
	(52145, 'TANGGUNG', 3935, 270, 18),
	(52146, 'TERATI', 3935, 270, 18),
	(52147, 'BELANGIN', 3936, 270, 18),
	(52148, 'BERINGIN', 3936, 270, 18),
	(52149, 'BOTUH LINTANG', 3936, 270, 18),
	(52150, 'BUNUT', 3936, 270, 18),
	(52151, 'ENTAKAI', 3936, 270, 18),
	(52152, 'ILIR KOTA', 3936, 270, 18),
	(52153, 'KAMBONG', 3936, 270, 18),
	(52154, 'LAPE', 3936, 270, 18),
	(52155, 'LINTANG KAPUAS', 3936, 270, 18),
	(52156, 'LINTANG PELAMAN', 3936, 270, 18),
	(52157, 'MENGKIANG', 3936, 270, 18),
	(52158, 'NANGA BIANG', 3936, 270, 18),
	(52159, 'PANA', 3936, 270, 18),
	(52160, 'PENYELADI', 3936, 270, 18),
	(52161, 'PENYELIMAU HILIR', 3936, 270, 18),
	(52162, 'PENYELIMAU JAYA', 3936, 270, 18),
	(52163, 'RAMBIN', 3936, 270, 18),
	(52164, 'SEI/SUNGAI ALAI', 3936, 270, 18),
	(52165, 'SEI/SUNGAI BATU', 3936, 270, 18),
	(52166, 'SEI/SUNGAI MAWANG', 3936, 270, 18),
	(52167, 'SEI/SUNGAI MUNTIK', 3936, 270, 18),
	(52168, 'SEI/SUNGAI SENGKUANG', 3936, 270, 18),
	(52169, 'SEMERANGKAI', 3936, 270, 18),
	(52170, 'TANJUNG KAPUAS', 3936, 270, 18),
	(52171, 'TANJUNG SEKAYAM', 3936, 270, 18),
	(52172, 'TAPANG DULANG', 3936, 270, 18),
	(52173, 'KELOMPU', 3937, 270, 18),
	(52174, 'KUALA DUA', 3937, 270, 18),
	(52175, 'MOBUI', 3937, 270, 18),
	(52176, 'SEBONGKUH', 3937, 270, 18),
	(52177, 'SEBUDUH', 3937, 270, 18),
	(52178, 'SEJUAH', 3937, 270, 18),
	(52179, 'SEMAYANG', 3937, 270, 18),
	(52180, 'TANAP', 3937, 270, 18),
	(52181, 'TANJUNG BUNGA', 3937, 270, 18),
	(52182, 'TANJUNG MERPATI', 3937, 270, 18),
	(52183, 'TUNGGAL BHAKTI', 3937, 270, 18),
	(52184, 'BAKTI JAYA', 3938, 270, 18),
	(52185, 'BALAI TINGGI', 3938, 270, 18),
	(52186, 'BARU LOMBAK', 3938, 270, 18),
	(52187, 'CUPANG', 3938, 270, 18),
	(52188, 'ENGGADAI', 3938, 270, 18),
	(52189, 'HARAPAN MAKMUR', 3938, 270, 18),
	(52190, 'KUALA BUAYAN', 3938, 270, 18),
	(52191, 'KUALA ROSAN', 3938, 270, 18),
	(52192, 'KUNYIL', 3938, 270, 18),
	(52193, 'LALANG', 3938, 270, 18),
	(52194, 'MELAWI MAKMUR', 3938, 270, 18),
	(52195, 'MELIAU HILIR', 3938, 270, 18),
	(52196, 'MELIAU HULU', 3938, 270, 18),
	(52197, 'MELOBOK', 3938, 270, 18),
	(52198, 'MERANGGAU', 3938, 270, 18),
	(52199, 'MUKTI JAYA', 3938, 270, 18),
	(52200, 'PAMPANG DUA', 3938, 270, 18),
	(52201, 'SEI/SUNGAI KEMBAYAU', 3938, 270, 18),
	(52202, 'SEI/SUNGAI MAYAN', 3938, 270, 18),
	(52203, 'ENGKODE', 3939, 270, 18),
	(52204, 'INGGIS', 3939, 270, 18),
	(52205, 'KEDUKUL', 3939, 270, 18),
	(52206, 'LAYAK OMANG', 3939, 270, 18),
	(52207, 'SEI/SUNGAI MAWANG', 3939, 270, 18),
	(52208, 'SEMANGGIS RAYA', 3939, 270, 18),
	(52209, 'SEMUNTAI', 3939, 270, 18),
	(52210, 'SERAMBAI JAYA', 3939, 270, 18),
	(52211, 'TRIMULYA', 3939, 270, 18),
	(52212, 'EMPOTO', 3940, 270, 18),
	(52213, 'IDAS', 3940, 270, 18),
	(52214, 'NOYAN', 3940, 270, 18),
	(52215, 'SEI/SUNGAI DANGIN', 3940, 270, 18),
	(52216, 'SEMONGAN', 3940, 270, 18),
	(52217, 'DOSAN', 3941, 270, 18),
	(52218, 'EMBALA', 3941, 270, 18),
	(52219, 'GUNAM', 3941, 270, 18),
	(52220, 'HIBUN', 3941, 270, 18),
	(52221, 'MAJU KARYA', 3941, 270, 18),
	(52222, 'MARIGIN JAYA', 3941, 270, 18),
	(52223, 'MARITA', 3941, 270, 18),
	(52224, 'PALEM JAYA', 3941, 270, 18),
	(52225, 'PANDU RAYA', 3941, 270, 18),
	(52226, 'PUSAT DAMAI', 3941, 270, 18),
	(52227, 'RAHAYU', 3941, 270, 18),
	(52228, 'SEBARA', 3941, 270, 18),
	(52229, 'SUKA GRUNDI', 3941, 270, 18),
	(52230, 'SUKA MULYA', 3941, 270, 18),
	(52231, 'BALAI KARANGAN', 3942, 270, 18),
	(52232, 'BUNGKANG', 3942, 270, 18),
	(52233, 'ENGKAHAN', 3942, 270, 18),
	(52234, 'KENAMAN', 3942, 270, 18),
	(52235, 'LUBUK SABUK', 3942, 270, 18),
	(52236, 'MALENGGANG', 3942, 270, 18),
	(52237, 'PENGADANG', 3942, 270, 18),
	(52238, 'RAUT MUARA', 3942, 270, 18),
	(52239, 'SEI/SUNGAI TEKAM', 3942, 270, 18),
	(52240, 'SOTOK', 3942, 270, 18),
	(52241, 'BALAI INGIN', 3943, 270, 18),
	(52242, 'BEGINJAN', 3943, 270, 18),
	(52243, 'CEMPEDAK', 3943, 270, 18),
	(52244, 'EMBERAS', 3943, 270, 18),
	(52245, 'KAWAT', 3943, 270, 18),
	(52246, 'LALANG', 3943, 270, 18),
	(52247, 'MELUGAI', 3943, 270, 18),
	(52248, 'PEDALAMAN', 3943, 270, 18),
	(52249, 'PULAU TAYAN UTARA', 3943, 270, 18),
	(52250, 'SEBEMBAN', 3943, 270, 18),
	(52251, 'SEI/SUNGAI JAMAN', 3943, 270, 18),
	(52252, 'SEJOTANG', 3943, 270, 18),
	(52253, 'SUBAH', 3943, 270, 18),
	(52254, 'TANJUNG BUNUT', 3943, 270, 18),
	(52255, 'TEBANG BENUA', 3943, 270, 18),
	(52256, 'BERAKAK', 3944, 270, 18),
	(52257, 'BINJAI', 3944, 270, 18),
	(52258, 'ENGKASAN', 3944, 270, 18),
	(52259, 'JANJANG', 3944, 270, 18),
	(52260, 'KEDAKAS', 3944, 270, 18),
	(52261, 'MANDONG', 3944, 270, 18),
	(52262, 'MENYABO', 3944, 270, 18),
	(52263, 'PANDAN SEMBUAT', 3944, 270, 18),
	(52264, 'PERUAN DALAM', 3944, 270, 18),
	(52265, 'RIYAI', 3944, 270, 18),
	(52266, 'SOSOK', 3944, 270, 18),
	(52267, 'BAGAN ASAM', 3945, 270, 18),
	(52268, 'BALAI BELUNGAI', 3945, 270, 18),
	(52269, 'BELUNGAI DALAM', 3945, 270, 18),
	(52270, 'KAMPUNG BARU', 3945, 270, 18),
	(52271, 'LUMUT', 3945, 270, 18),
	(52272, 'SANSAT', 3945, 270, 18),
	(52273, 'TERAJU (TERAJU BARAT)', 3945, 270, 18),
	(52274, 'BEDILAN', 3946, 271, 18),
	(52275, 'BELITANG DUA', 3946, 271, 18),
	(52276, 'BELITANG SATU', 3946, 271, 18),
	(52277, 'GEDUNG REJO', 3946, 271, 18),
	(52278, 'GUMAWANG', 3946, 271, 18),
	(52279, 'GUNUNG MAS', 3946, 271, 18),
	(52280, 'HARJO WINANGUN', 3946, 271, 18),
	(52281, 'KARANG KEMIRI', 3946, 271, 18),
	(52282, 'MABOH PERMAI', 3946, 271, 18),
	(52283, 'MENUA PRAMA', 3946, 271, 18),
	(52284, 'NANGA ANSAR', 3946, 271, 18),
	(52285, 'PADAK', 3946, 271, 18),
	(52286, 'PUJORAHAYU', 3946, 271, 18),
	(52287, 'SERBAGUNA', 3946, 271, 18),
	(52288, 'SETUNTUNG', 3946, 271, 18),
	(52289, 'SIDO MAKMUR', 3946, 271, 18),
	(52290, 'SIDO RAHAYU', 3946, 271, 18),
	(52291, 'SIDODADI', 3946, 271, 18),
	(52292, 'SIDOGEDE', 3946, 271, 18),
	(52293, 'SIDOMULYO', 3946, 271, 18),
	(52294, 'SUKAJADI', 3946, 271, 18),
	(52295, 'SUKARAMI', 3946, 271, 18),
	(52296, 'SUKOSARI', 3946, 271, 18),
	(52297, 'SUMBER SUKO', 3946, 271, 18),
	(52298, 'SUMBER SUKO JAYA', 3946, 271, 18),
	(52299, 'TANJUNG RAYA', 3946, 271, 18),
	(52300, 'TAWANG REJO', 3946, 271, 18),
	(52301, 'TEGAL REJO', 3946, 271, 18),
	(52302, 'TRIYOSO', 3946, 271, 18),
	(52303, 'EMPAJAK', 3947, 271, 18),
	(52304, 'ENTABUK', 3947, 271, 18),
	(52305, 'KUMPANG BIS', 3947, 271, 18),
	(52306, 'MENAWAI TEKAM', 3947, 271, 18),
	(52307, 'MERBANG', 3947, 271, 18),
	(52308, 'SEI/SUNGAI AYAK I', 3947, 271, 18),
	(52309, 'SEI/SUNGAI AYAK III', 3947, 271, 18),
	(52310, 'SEMADU', 3947, 271, 18),
	(52311, 'TAPANG PULAU', 3947, 271, 18),
	(52312, 'BALAI SEPUAK', 3948, 271, 18),
	(52313, 'BATUK MULAU', 3948, 271, 18),
	(52314, 'BUKIT RAMBAT', 3948, 271, 18),
	(52315, 'IJUK (IJOK)', 3948, 271, 18),
	(52316, 'KUMPANG ILONG', 3948, 271, 18),
	(52317, 'MENGARET', 3948, 271, 18),
	(52318, 'PAKET/PAKIT MULAU', 3948, 271, 18),
	(52319, 'SEBETUNG', 3948, 271, 18),
	(52320, 'SEBURUK', 3948, 271, 18),
	(52321, 'SEI/SUNGAI ANTU HULU', 3948, 271, 18),
	(52322, 'SEI/SUNGAI TAPAH', 3948, 271, 18),
	(52323, 'TABUK HULU', 3948, 271, 18),
	(52324, 'TERDUK DAMPAK', 3948, 271, 18),
	(52325, 'BATU PAHAT', 3949, 271, 18),
	(52326, 'CENAYAN', 3949, 271, 18),
	(52327, 'KARANG BETUNG', 3949, 271, 18),
	(52328, 'LANDAU APIN', 3949, 271, 18),
	(52329, 'LANDAU KUMPAI', 3949, 271, 18),
	(52330, 'LEMBAH BERINGIN', 3949, 271, 18),
	(52331, 'NANGA MAHAP', 3949, 271, 18),
	(52332, 'NANGA SURI', 3949, 271, 18),
	(52333, 'SEBABAS (SE BABAS)', 3949, 271, 18),
	(52334, 'TAMANG', 3949, 271, 18),
	(52335, 'TELUK KEBAU', 3949, 271, 18),
	(52336, 'TEMBAGA', 3949, 271, 18),
	(52337, 'TEMBESUK', 3949, 271, 18),
	(52338, 'LUBUK TAJAU', 3950, 271, 18),
	(52339, 'MERAGUN', 3950, 271, 18),
	(52340, 'NANGA ENGKULUN', 3950, 271, 18),
	(52341, 'NANGA KIUNGKANG', 3950, 271, 18),
	(52342, 'NANGA KOMAN', 3950, 271, 18),
	(52343, 'NANGA MENTUKAK', 3950, 271, 18),
	(52344, 'NANGA MONGKO', 3950, 271, 18),
	(52345, 'NANGA TAMAN', 3950, 271, 18),
	(52346, 'PANTOK', 3950, 271, 18),
	(52347, 'RIRANG JATI', 3950, 271, 18),
	(52348, 'SEI/SUNGAI LAWAK', 3950, 271, 18),
	(52349, 'SENANGAK', 3950, 271, 18),
	(52350, 'TAPANG TINGANG', 3950, 271, 18),
	(52351, 'BOKAK SEBUMBUN', 3951, 271, 18),
	(52352, 'ENGKERSIK', 3951, 271, 18),
	(52353, 'ENSALANG', 3951, 271, 18),
	(52354, 'GONIS TEKAM', 3951, 271, 18),
	(52355, 'LANDAU KODAH', 3951, 271, 18),
	(52356, 'MERAPI', 3951, 271, 18),
	(52357, 'MUNGGUK', 3951, 271, 18),
	(52358, 'PENITI', 3951, 271, 18),
	(52359, 'SEBERANG KAPUAS', 3951, 271, 18),
	(52360, 'SEI/SUNGAI KUNYIT', 3951, 271, 18),
	(52361, 'SEI/SUNGAI RINGIN', 3951, 271, 18),
	(52362, 'SELALONG', 3951, 271, 18),
	(52363, 'SEMABI', 3951, 271, 18),
	(52364, 'SERARAS', 3951, 271, 18),
	(52365, 'TANJUNG', 3951, 271, 18),
	(52366, 'TAPANG SEMADAK', 3951, 271, 18),
	(52367, 'TIMPUK', 3951, 271, 18),
	(52368, 'BOTI', 3952, 271, 18),
	(52369, 'CUPANG GADING', 3952, 271, 18),
	(52370, 'MONDI', 3952, 271, 18),
	(52371, 'NANGA BIABAN', 3952, 271, 18),
	(52372, 'NANGA MANTERAP', 3952, 271, 18),
	(52373, 'NANGA PEMUBUH', 3952, 271, 18),
	(52374, 'PERONGKAN', 3952, 271, 18),
	(52375, 'RAWAK HILIR', 3952, 271, 18),
	(52376, 'RAWAK HULU', 3952, 271, 18),
	(52377, 'SEI/SUNGAI SAMBANG', 3952, 271, 18),
	(52378, 'SEKONAU', 3952, 271, 18),
	(52379, 'SETAWAR', 3952, 271, 18),
	(52380, 'SUNSONG', 3952, 271, 18),
	(52381, 'TAPANG PERODAH', 3952, 271, 18),
	(52382, 'TINTING BOYOK', 3952, 271, 18),
	(52383, 'KUALA', 3953, 272, 18),
	(52384, 'MELAYU', 3953, 272, 18),
	(52385, 'PASIRAN', 3953, 272, 18),
	(52386, 'TENGAH', 3953, 272, 18),
	(52387, 'PANGMILANG', 3954, 272, 18),
	(52388, 'SAGATANI', 3954, 272, 18),
	(52389, 'SEDAU', 3954, 272, 18),
	(52390, 'SIJANGKUNG', 3954, 272, 18),
	(52391, 'BUKIT BATU', 3955, 272, 18),
	(52392, 'CONDONG', 3955, 272, 18),
	(52393, 'JAWA', 3955, 272, 18),
	(52394, 'ROBAN', 3955, 272, 18),
	(52395, 'SEI/SUNGAI WIE', 3955, 272, 18),
	(52396, 'SEKIP LAMA', 3955, 272, 18),
	(52397, 'BAGAK SAHWA', 3956, 272, 18),
	(52398, 'MAYASOPA', 3956, 272, 18),
	(52399, 'NYARUNGKOP', 3956, 272, 18),
	(52400, 'PAJINTAN', 3956, 272, 18),
	(52401, 'SANGGAU KOLOR', 3956, 272, 18),
	(52402, 'NARAM', 3957, 272, 18),
	(52403, 'SEI/SUNGAI BULAN', 3957, 272, 18),
	(52404, 'SEI/SUNGAI GARAM HILIR', 3957, 272, 18),
	(52405, 'SEI/SUNGAI RASAU', 3957, 272, 18),
	(52406, 'SEMELAGI KECIL', 3957, 272, 18),
	(52407, 'SETAPUK BESAR', 3957, 272, 18),
	(52408, 'SETAPUK KECIL', 3957, 272, 18),
	(52409, 'AMPAR BEDANG', 3958, 273, 18),
	(52410, 'BINJAI HILIR', 3958, 273, 18),
	(52411, 'BINJAI HULU', 3958, 273, 18),
	(52412, 'DAK JAYA', 3958, 273, 18),
	(52413, 'EMPAKA KABIAU RAYA', 3958, 273, 18),
	(52414, 'MENSIKU', 3958, 273, 18),
	(52415, 'SEI/SUNGAI RISAP', 3958, 273, 18),
	(52416, 'SIMBA RAYA', 3958, 273, 18),
	(52417, 'SUNGAI RISAP MENSIKU BERSATU', 3958, 273, 18),
	(52418, 'TELAGA DUA', 3958, 273, 18),
	(52419, 'TELAGA SATU', 3958, 273, 18),
	(52420, 'BARAS', 3959, 273, 18),
	(52421, 'DEDAI KANAN', 3959, 273, 18),
	(52422, 'EMPACI', 3959, 273, 18),
	(52423, 'EMPARU BARU', 3959, 273, 18),
	(52424, 'GANIS', 3959, 273, 18),
	(52425, 'KUMPANG', 3959, 273, 18),
	(52426, 'LUNDANG BARU', 3959, 273, 18),
	(52427, 'MANGAT BARU', 3959, 273, 18),
	(52428, 'MENAONG BARU', 3959, 273, 18),
	(52429, 'MEREMPIT BARU', 3959, 273, 18),
	(52430, 'NANGA DEDAI', 3959, 273, 18),
	(52431, 'NANGA JETAK', 3959, 273, 18),
	(52432, 'PENGKADAN BARU', 3959, 273, 18),
	(52433, 'PENGKADAN SUNGAI RUPA', 3959, 273, 18),
	(52434, 'PENYAK LALANG', 3959, 273, 18),
	(52435, 'RIGUK', 3959, 273, 18),
	(52436, 'SEI/SUNGAI MALI', 3959, 273, 18),
	(52437, 'SUNGAI TAPANG', 3959, 273, 18),
	(52438, 'TAUK', 3959, 273, 18),
	(52439, 'UMIN JAYA', 3959, 273, 18),
	(52440, 'BUKIT SEGALOH', 3960, 273, 18),
	(52441, 'DATA DIAN', 3960, 273, 18),
	(52442, 'ENGKERANGAN', 3960, 273, 18),
	(52443, 'JAMBU', 3960, 273, 18),
	(52444, 'JAYA SAKTI', 3960, 273, 18),
	(52445, 'KARYA BARU', 3960, 273, 18),
	(52446, 'KERAPA SEPAN', 3960, 273, 18),
	(52447, 'LALANG INGGAR', 3960, 273, 18),
	(52448, 'LANDAU BERINGIN', 3960, 273, 18),
	(52449, 'LINGGAM PERMAI', 3960, 273, 18),
	(52450, 'LONG METUN (LONG METUN II)', 3960, 273, 18),
	(52451, 'LONG PIPA', 3960, 273, 18),
	(52452, 'MEKAR MANDIRI', 3960, 273, 18),
	(52453, 'MELINGKAT', 3960, 273, 18),
	(52454, 'MENTUNAI', 3960, 273, 18),
	(52455, 'METUN (LONG METUN I)', 3960, 273, 18),
	(52456, 'NANGA MAU', 3960, 273, 18),
	(52457, 'NANGA TIKAN', 3960, 273, 18),
	(52458, 'NATAI LESUNG', 3960, 273, 18),
	(52459, 'NATAI TEBEDAK', 3960, 273, 18),
	(52460, 'NYANGKOM', 3960, 273, 18),
	(52461, 'PAKAK', 3960, 273, 18),
	(52462, 'PAOH DESA', 3960, 273, 18),
	(52463, 'PELAIK', 3960, 273, 18),
	(52464, 'SEI/SUNGAI ANAI', 3960, 273, 18),
	(52465, 'SEI/SUNGAI BUAYA', 3960, 273, 18),
	(52466, 'SUNGAI GARONG', 3960, 273, 18),
	(52467, 'SUNGAI MENUANG', 3960, 273, 18),
	(52468, 'SUNGAI SINTANG', 3960, 273, 18),
	(52469, 'TERTUNG MAU', 3960, 273, 18),
	(52470, 'TUGUK', 3960, 273, 18),
	(52471, 'BULUH MERINDU', 3961, 273, 18),
	(52472, 'EMPAKAN', 3961, 273, 18),
	(52473, 'EMPOYANG', 3961, 273, 18),
	(52474, 'ENTOGONG', 3961, 273, 18),
	(52475, 'KERAPUK JAYA', 3961, 273, 18),
	(52476, 'LANDAU BARA', 3961, 273, 18),
	(52477, 'LINTANG TAMBUK', 3961, 273, 18),
	(52478, 'LONG BETAOH', 3961, 273, 18),
	(52479, 'LONG NAWANG', 3961, 273, 18),
	(52480, 'LONG PAYAU', 3961, 273, 18),
	(52481, 'LONG TEMUNYAT/TEMUYAT', 3961, 273, 18),
	(52482, 'MAPAN JAYA', 3961, 273, 18),
	(52483, 'MERAH ARAI', 3961, 273, 18),
	(52484, 'MERAHAU PERMAI', 3961, 273, 18),
	(52485, 'NANGA ABAI', 3961, 273, 18),
	(52486, 'NANGA LAAR', 3961, 273, 18),
	(52487, 'NANGA MASAU', 3961, 273, 18),
	(52488, 'NANGA ORAN', 3961, 273, 18),
	(52489, 'NANGA PAYAK', 3961, 273, 18),
	(52490, 'NANGA TEBIDAH', 3961, 273, 18),
	(52491, 'NANGA TONGGOI', 3961, 273, 18),
	(52492, 'NANGA TORAN', 3961, 273, 18),
	(52493, 'NANGA UNGAI', 3961, 273, 18),
	(52494, 'NANGKAK LESTARI', 3961, 273, 18),
	(52495, 'NAWANG BARU', 3961, 273, 18),
	(52496, 'RIAM MUNTIK', 3961, 273, 18),
	(52497, 'RIAM PANJANG', 3961, 273, 18),
	(52498, 'TANAH MERAH', 3961, 273, 18),
	(52499, 'TANJUNG BUNGA', 3961, 273, 18),
	(52500, 'TANJUNG LALAU', 3961, 273, 18),
	(52501, 'TANJUNG MIRU', 3961, 273, 18),
	(52502, 'TAPANG MANUA', 3961, 273, 18),
	(52503, 'TONAK GONEH', 3961, 273, 18),
	(52504, 'TOPAN NANGA', 3961, 273, 18),
	(52505, 'BANING PANJANG', 3962, 273, 18),
	(52506, 'BENGKUANG', 3962, 273, 18),
	(52507, 'ENSAID PANJANG', 3962, 273, 18),
	(52508, 'GEMBA RAYA', 3962, 273, 18),
	(52509, 'KARYA JAYA BHAKTI', 3962, 273, 18),
	(52510, 'KEBONG', 3962, 273, 18),
	(52511, 'LANDAU KODAM', 3962, 273, 18),
	(52512, 'MANDIRI JAYA', 3962, 273, 18),
	(52513, 'MERPAK', 3962, 273, 18),
	(52514, 'NANGA LEBANG', 3962, 273, 18),
	(52515, 'PELIMPING', 3962, 273, 18),
	(52516, 'SEI/SUNGAI MARAM', 3962, 273, 18),
	(52517, 'SEI/SUNGAI PUKAT', 3962, 273, 18),
	(52518, 'SEPAN LEBANG', 3962, 273, 18),
	(52519, 'SUNGAI LABI', 3962, 273, 18),
	(52520, 'SUNGAI LAIS', 3962, 273, 18),
	(52521, 'AIR NYURUK', 3963, 273, 18),
	(52522, 'BATU AMPAR', 3963, 273, 18),
	(52523, 'BAUNG SENGATAP', 3963, 273, 18),
	(52524, 'BELUH MULYO', 3963, 273, 18),
	(52525, 'BETUNG PERMAI', 3963, 273, 18),
	(52526, 'BUKIT SIDIN PERMAI', 3963, 273, 18),
	(52527, 'KENUAK', 3963, 273, 18),
	(52528, 'NANGA KETUNGAU', 3963, 273, 18),
	(52529, 'NANGA MERKAK', 3963, 273, 18),
	(52530, 'NANGA SEJIRAK', 3963, 273, 18),
	(52531, 'SEI/SUNGAI DERAS', 3963, 273, 18),
	(52532, 'SEI/SUNGAI MALI', 3963, 273, 18),
	(52533, 'SEMAJAU MEKAR', 3963, 273, 18),
	(52534, 'SEMUNTAI', 3963, 273, 18),
	(52535, 'SENIBUNG', 3963, 273, 18),
	(52536, 'SETUNGKUP', 3963, 273, 18),
	(52537, 'TANJUNG BAUNG', 3963, 273, 18),
	(52538, 'BEKUAN LUYANG', 3964, 273, 18),
	(52539, 'EMPUNAK TAPANG KELADAN', 3964, 273, 18),
	(52540, 'EMPURA', 3964, 273, 18),
	(52541, 'JASA', 3964, 273, 18),
	(52542, 'MUAKAN PETINGGI', 3964, 273, 18),
	(52543, 'NANGA BAYAN', 3964, 273, 18),
	(52544, 'NANGA BUGAU', 3964, 273, 18),
	(52545, 'NANGA SEBAWANG', 3964, 273, 18),
	(52546, 'RASAU', 3964, 273, 18),
	(52547, 'SEBADAK', 3964, 273, 18),
	(52548, 'SEBETUNG PALUK', 3964, 273, 18),
	(52549, 'SEI/SUNGAI SERIA', 3964, 273, 18),
	(52550, 'SEKAIH', 3964, 273, 18),
	(52551, 'SENANING', 3964, 273, 18),
	(52552, 'SEPILUK', 3964, 273, 18),
	(52553, 'SUAK MEDANG', 3964, 273, 18),
	(52554, 'SUNGAI BUGAU', 3964, 273, 18),
	(52555, 'SUNGAI PISAU', 3964, 273, 18),
	(52556, 'ARGO MULYO', 3965, 273, 18),
	(52557, 'BEGELANG JAYA', 3965, 273, 18),
	(52558, 'GUT JAYA BAKTI', 3965, 273, 18),
	(52559, 'KAYU DUJUNG', 3965, 273, 18),
	(52560, 'KERTA SARI', 3965, 273, 18),
	(52561, 'LANDAU BUAYA', 3965, 273, 18),
	(52562, 'MARGA HAYU', 3965, 273, 18),
	(52563, 'MUNGGUK GELOMBANG', 3965, 273, 18),
	(52564, 'MUNGGUK LAWANG', 3965, 273, 18),
	(52565, 'NANGA KELAPAN', 3965, 273, 18),
	(52566, 'PANDING JAYA', 3965, 273, 18),
	(52567, 'PANGGI AGUNG', 3965, 273, 18),
	(52568, 'SENANGAN KECIL', 3965, 273, 18),
	(52569, 'SUMBER SARI', 3965, 273, 18),
	(52570, 'SUNGAI AREH', 3965, 273, 18),
	(52571, 'SWADAYA', 3965, 273, 18),
	(52572, 'TANJUNG SARI', 3965, 273, 18),
	(52573, 'TERTA KARYA', 3965, 273, 18),
	(52574, 'WANA BHAKTI', 3965, 273, 18),
	(52575, 'WIRAYUDA', 3965, 273, 18),
	(52576, 'BANGUN', 3966, 273, 18),
	(52577, 'BEDAYAN', 3966, 273, 18),
	(52578, 'BERNAYAU', 3966, 273, 18),
	(52579, 'BULUH KUNING', 3966, 273, 18),
	(52580, 'ENSABANG', 3966, 273, 18),
	(52581, 'GERNIS JAYA', 3966, 273, 18),
	(52582, 'KEMANTAN', 3966, 273, 18),
	(52583, 'KENYAUK', 3966, 273, 18),
	(52584, 'LANDAU PANJANG', 3966, 273, 18),
	(52585, 'LENGKENAT', 3966, 273, 18),
	(52586, 'MAIT HILIR', 3966, 273, 18),
	(52587, 'MANIS RAYA', 3966, 273, 18),
	(52588, 'NANGA LIBAU', 3966, 273, 18),
	(52589, 'NANGA PARI', 3966, 273, 18),
	(52590, 'NANGA SEPAUK', 3966, 273, 18),
	(52591, 'PAOH BENUA', 3966, 273, 18),
	(52592, 'PENINSUNG', 3966, 273, 18),
	(52593, 'SEKUBANG', 3966, 273, 18),
	(52594, 'SEKUJAM TIMBAI', 3966, 273, 18),
	(52595, 'SEMUNTAI', 3966, 273, 18),
	(52596, 'SEPULUT', 3966, 273, 18),
	(52597, 'SINAR PEKAYAU', 3966, 273, 18),
	(52598, 'SIRANG SETAMBANG', 3966, 273, 18),
	(52599, 'SUKAU BERSATU', 3966, 273, 18),
	(52600, 'SUNGAI RAYA', 3966, 273, 18),
	(52601, 'SUNGAI SEGAK', 3966, 273, 18),
	(52602, 'TANJUNG BALAI', 3966, 273, 18),
	(52603, 'TANJUNG HULU', 3966, 273, 18),
	(52604, 'TANJUNG RIA', 3966, 273, 18),
	(52605, 'TAWANG SARI', 3966, 273, 18),
	(52606, 'TEMAWANG BULAI', 3966, 273, 18),
	(52607, 'TEMAWANG MUNTAI', 3966, 273, 18),
	(52608, 'TEMIANG KAPUAS', 3966, 273, 18),
	(52609, 'BARAS NABUN', 3967, 273, 18),
	(52610, 'BEDAHA', 3967, 273, 18),
	(52611, 'BEGORI', 3967, 273, 18),
	(52612, 'BUNTUT PONTE', 3967, 273, 18),
	(52613, 'GURUNG SENGHIANG', 3967, 273, 18),
	(52614, 'KARYA JAYA', 3967, 273, 18),
	(52615, 'MENTAJOI', 3967, 273, 18),
	(52616, 'MERAKO JAYA', 3967, 273, 18),
	(52617, 'NANGA BIHE', 3967, 273, 18),
	(52618, 'NANGA JELUNDUNG', 3967, 273, 18),
	(52619, 'NANGA MENTATAI', 3967, 273, 18),
	(52620, 'NANGA RIYOI', 3967, 273, 18),
	(52621, 'NANGA SEGULANG', 3967, 273, 18),
	(52622, 'NANGA SERAWAI', 3967, 273, 18),
	(52623, 'NANGA TEKUNGAI', 3967, 273, 18),
	(52624, 'PAGAR LEBATA', 3967, 273, 18),
	(52625, 'RANTAU MALAM', 3967, 273, 18),
	(52626, 'SAWANG SENGHIANG', 3967, 273, 18),
	(52627, 'TAHAI PERMAI', 3967, 273, 18),
	(52628, 'TANJUNG HARAPAN', 3967, 273, 18),
	(52629, 'TANJUNG RAYA', 3967, 273, 18),
	(52630, 'TELUK HARAPAN', 3967, 273, 18),
	(52631, 'TONTANG', 3967, 273, 18),
	(52632, 'TUNAS HARAPAN', 3967, 273, 18),
	(52633, 'ANGGAH JAYA', 3968, 273, 18),
	(52634, 'BANING KOTA', 3968, 273, 18),
	(52635, 'KAPUAS KANAN HILIR', 3968, 273, 18),
	(52636, 'KAPUAS KANAN HULU', 3968, 273, 18),
	(52637, 'KAPUAS KIRI HILIR', 3968, 273, 18),
	(52638, 'KAPUAS KIRI HULU', 3968, 273, 18),
	(52639, 'LADANG', 3968, 273, 18),
	(52640, 'LALANG BARU', 3968, 273, 18),
	(52641, 'MERTI GUNA', 3968, 273, 18),
	(52642, 'MUNGGUK BANTOK', 3968, 273, 18),
	(52643, 'SUNGAI ANA', 3968, 273, 18),
	(52644, 'TANJUNG KELANSAM', 3968, 273, 18),
	(52645, 'TANJUNG PURI', 3968, 273, 18),
	(52646, 'TELUK KELANSAM', 3968, 273, 18),
	(52647, 'TERTONG', 3968, 273, 18),
	(52648, 'BANCOH', 3969, 273, 18),
	(52649, 'BAYA (BAYA BETONG)', 3969, 273, 18),
	(52650, 'BONET ENGKABANG', 3969, 273, 18),
	(52651, 'BONET LAMA', 3969, 273, 18),
	(52652, 'GURUNG KEMPADIK', 3969, 273, 18),
	(52653, 'KAJANG BARU', 3969, 273, 18),
	(52654, 'LAMAN NATAI', 3969, 273, 18),
	(52655, 'LEBAK UBAH', 3969, 273, 18),
	(52656, 'MANTER', 3969, 273, 18),
	(52657, 'MELAYANG SARI', 3969, 273, 18),
	(52658, 'MERARAI DUA', 3969, 273, 18),
	(52659, 'MERARAI SATU', 3969, 273, 18),
	(52660, 'NOBAL', 3969, 273, 18),
	(52661, 'PENJERNANG', 3969, 273, 18),
	(52662, 'PENJERNANG HULU', 3969, 273, 18),
	(52663, 'PEREMBANG', 3969, 273, 18),
	(52664, 'RANSI DAKAN', 3969, 273, 18),
	(52665, 'RARAI', 3969, 273, 18),
	(52666, 'RIAM KIJANG', 3969, 273, 18),
	(52667, 'SARAI', 3969, 273, 18),
	(52668, 'SEI/SUNGAI UKOI', 3969, 273, 18),
	(52669, 'BALAI HARAPAN', 3970, 273, 18),
	(52670, 'BENUA BARU', 3970, 273, 18),
	(52671, 'BENUA KENCANA', 3970, 273, 18),
	(52672, 'GURUNG MALI', 3970, 273, 18),
	(52673, 'JAYA MENTARI', 3970, 273, 18),
	(52674, 'KENYABUR BARU', 3970, 273, 18),
	(52675, 'KUALA TIGA', 3970, 273, 18),
	(52676, 'KUPAN JAYA', 3970, 273, 18),
	(52677, 'MENGKURAT BARU', 3970, 273, 18),
	(52678, 'MENSIAP BARU', 3970, 273, 18),
	(52679, 'MERTI JAYA', 3970, 273, 18),
	(52680, 'NANGA TEMPUNAK', 3970, 273, 18),
	(52681, 'PAGAL BARU', 3970, 273, 18),
	(52682, 'PANGKAL BARU', 3970, 273, 18),
	(52683, 'PARIBANG BARU', 3970, 273, 18),
	(52684, 'PUDAU BERSATU', 3970, 273, 18),
	(52685, 'PULAU JAYA', 3970, 273, 18),
	(52686, 'REPAK SARI', 3970, 273, 18),
	(52687, 'RIAM BATU', 3970, 273, 18),
	(52688, 'SUKA JAYA', 3970, 273, 18),
	(52689, 'SUNGAI BULUH', 3970, 273, 18),
	(52690, 'TANJUNG PERADA', 3970, 273, 18),
	(52691, 'TEMPUNAK KAPUAS', 3970, 273, 18),
	(52692, 'TINUM BARU', 3970, 273, 18),
	(52693, 'BATAMPANG', 3971, 274, 19),
	(52694, 'BATILAP', 3971, 274, 19),
	(52695, 'DAMPARAN', 3971, 274, 19),
	(52696, 'KALANIS', 3971, 274, 19),
	(52697, 'LEHAI', 3971, 274, 19),
	(52698, 'MAHAJANDAU/MAHANJAU', 3971, 274, 19),
	(52699, 'MANGKATIR', 3971, 274, 19),
	(52700, 'MENGKATIP', 3971, 274, 19),
	(52701, 'SEI/SUNGAI JAYA', 3971, 274, 19),
	(52702, 'TELUK TIMBAU', 3971, 274, 19),
	(52703, 'BARU', 3972, 274, 19),
	(52704, 'BUNTOK KOTA', 3972, 274, 19),
	(52705, 'DANAU GANTING', 3972, 274, 19),
	(52706, 'DANAU MASURA', 3972, 274, 19),
	(52707, 'DANAU SADAR', 3972, 274, 19),
	(52708, 'DANGKA', 3972, 274, 19),
	(52709, 'HILIR SPER', 3972, 274, 19),
	(52710, 'JELAPAT', 3972, 274, 19),
	(52711, 'KALAHIEN', 3972, 274, 19),
	(52712, 'LEMBENG', 3972, 274, 19),
	(52713, 'MABUAN', 3972, 274, 19),
	(52714, 'MADARA', 3972, 274, 19),
	(52715, 'MANGARIS', 3972, 274, 19),
	(52716, 'MUARA RIPUNG', 3972, 274, 19),
	(52717, 'MUARA TALANG', 3972, 274, 19),
	(52718, 'MURUNG PAKEN', 3972, 274, 19),
	(52719, 'PAMAIT', 3972, 274, 19),
	(52720, 'PAMANGKA', 3972, 274, 19),
	(52721, 'PARARAPAK', 3972, 274, 19),
	(52722, 'PENDA ASAM', 3972, 274, 19),
	(52723, 'SABABILAH', 3972, 274, 19),
	(52724, 'SANGGU', 3972, 274, 19),
	(52725, 'TANJUNG JAWA', 3972, 274, 19),
	(52726, 'TELANG ANDRAU', 3972, 274, 19),
	(52727, 'TELUK MAMPUN', 3972, 274, 19),
	(52728, 'TELUK TELAGA', 3972, 274, 19),
	(52729, 'TETEI LANAN', 3972, 274, 19),
	(52730, 'BANTAI BAMBURE', 3973, 274, 19),
	(52731, 'BUNDAR', 3973, 274, 19),
	(52732, 'DANAU BAMBURE', 3973, 274, 19),
	(52733, 'GUNUNG RANTAU', 3973, 274, 19),
	(52734, 'HINGAN', 3973, 274, 19),
	(52735, 'HULU TAMPANG', 3973, 274, 19),
	(52736, 'MAJUNDRE', 3973, 274, 19),
	(52737, 'MARAWAN BARU', 3973, 274, 19),
	(52738, 'MARAWAN LAMA', 3973, 274, 19),
	(52739, 'MARUGA', 3973, 274, 19),
	(52740, 'PANARUKAN', 3973, 274, 19),
	(52741, 'PENDANG', 3973, 274, 19),
	(52742, 'RAMPA MEA', 3973, 274, 19),
	(52743, 'REONG', 3973, 274, 19),
	(52744, 'SEI/SUNGAI TELANG', 3973, 274, 19),
	(52745, 'TAMPARAK', 3973, 274, 19),
	(52746, 'TAMPARAK LAYUNG', 3973, 274, 19),
	(52747, 'TARUSAN', 3973, 274, 19),
	(52748, 'TELEKOI (TALEKOI)', 3973, 274, 19),
	(52749, 'BARUANG/EKENG', 3974, 274, 19),
	(52750, 'BINTANG ARA', 3974, 274, 19),
	(52751, 'BIPAK KALI', 3974, 274, 19),
	(52752, 'GAGUTUR', 3974, 274, 19),
	(52753, 'KAYUMBAN', 3974, 274, 19),
	(52754, 'MALUNGAI RAYA', 3974, 274, 19),
	(52755, 'MARGA JAYA', 3974, 274, 19),
	(52756, 'MUARA SINGAN', 3974, 274, 19),
	(52757, 'MUKA HAJI', 3974, 274, 19),
	(52758, 'NGURIT', 3974, 274, 19),
	(52759, 'PALO REJO', 3974, 274, 19),
	(52760, 'PATAS I', 3974, 274, 19),
	(52761, 'PATAS II', 3974, 274, 19),
	(52762, 'RUHING RAYA', 3974, 274, 19),
	(52763, 'SARIMBUAH', 3974, 274, 19),
	(52764, 'SEI/SUNGAI PAKEN', 3974, 274, 19),
	(52765, 'SIRE', 3974, 274, 19),
	(52766, 'TABAK KANILAN', 3974, 274, 19),
	(52767, 'UGANG SAYU', 3974, 274, 19),
	(52768, 'WAYUN', 3974, 274, 19),
	(52769, 'WUNGKUR BARU', 3974, 274, 19),
	(52770, 'RANGGA ILUNG', 3975, 274, 19),
	(52771, 'RANTAU BAHUANG', 3975, 274, 19),
	(52772, 'RANTAU KUJANG', 3975, 274, 19),
	(52773, 'TABATAN', 3975, 274, 19),
	(52774, 'TAMPULANG', 3975, 274, 19),
	(52775, 'BABAI', 3976, 274, 19),
	(52776, 'BANGKUANG', 3976, 274, 19),
	(52777, 'BINTANG KURUNG', 3976, 274, 19),
	(52778, 'JANGGI', 3976, 274, 19),
	(52779, 'MALITIN', 3976, 274, 19),
	(52780, 'MUARA ARAI', 3976, 274, 19),
	(52781, 'SELAT BARU', 3976, 274, 19),
	(52782, 'TALIO', 3976, 274, 19),
	(52783, 'TAMPIJAK', 3976, 274, 19),
	(52784, 'TELUK BETUNG', 3976, 274, 19),
	(52785, 'TELUK SAMPUDAU', 3976, 274, 19),
	(52786, 'AMPARI', 3977, 275, 19),
	(52787, 'APAR BATU', 3977, 275, 19),
	(52788, 'BANGKIRAYEN', 3977, 275, 19),
	(52789, 'BIWAN', 3977, 275, 19),
	(52790, 'DANAU', 3977, 275, 19),
	(52791, 'HAYAPING', 3977, 275, 19),
	(52792, 'JANAH JARI', 3977, 275, 19),
	(52793, 'JANAH MANSIWUI', 3977, 275, 19),
	(52794, 'PIANGGU', 3977, 275, 19),
	(52795, 'TANGKAN', 3977, 275, 19),
	(52796, 'WUNGKUR NANAKAN', 3977, 275, 19),
	(52797, 'BAGOK', 3978, 275, 19),
	(52798, 'BAMBAN', 3978, 275, 19),
	(52799, 'BANYU LANDAS', 3978, 275, 19),
	(52800, 'GUDANG SENG', 3978, 275, 19),
	(52801, 'KANDRIS', 3978, 275, 19),
	(52802, 'TANIRAN PASAR PANAS', 3978, 275, 19),
	(52803, 'TEWAH PUPUH', 3978, 275, 19),
	(52804, 'AMPAH DUA', 3979, 275, 19),
	(52805, 'AMPAH KOTA', 3979, 275, 19),
	(52806, 'MUARA AWANG', 3979, 275, 19),
	(52807, 'NETAPIN / NETAMPIN', 3979, 275, 19),
	(52808, 'PUTAI', 3979, 275, 19),
	(52809, 'RODOK', 3979, 275, 19),
	(52810, 'SAING', 3979, 275, 19),
	(52811, 'SUMBER GARONGGONG', 3979, 275, 19),
	(52812, 'ANGKARAP', 3980, 275, 19),
	(52813, 'DIDI', 3980, 275, 19),
	(52814, 'DORONG', 3980, 275, 19),
	(52815, 'GUMPA', 3980, 275, 19),
	(52816, 'HARARA', 3980, 275, 19),
	(52817, 'HARINGEN', 3980, 275, 19),
	(52818, 'JAAR', 3980, 275, 19),
	(52819, 'JAWETEN', 3980, 275, 19),
	(52820, 'KARANG LANGIT', 3980, 275, 19),
	(52821, 'MAGANTIS', 3980, 275, 19),
	(52822, 'MARAGUT', 3980, 275, 19),
	(52823, 'MATABU', 3980, 275, 19),
	(52824, 'MATARAH', 3980, 275, 19),
	(52825, 'PULAU PATAI', 3980, 275, 19),
	(52826, 'SARAPAT (EX. HARARA)', 3980, 275, 19),
	(52827, 'SUMUR', 3980, 275, 19),
	(52828, 'TAMIANG LAYANG', 3980, 275, 19),
	(52829, 'DAYU', 3981, 275, 19),
	(52830, 'IPU MEA', 3981, 275, 19),
	(52831, 'KANDRIS', 3981, 275, 19),
	(52832, 'LAGAN', 3981, 275, 19),
	(52833, 'PUTUT TAWULUH', 3981, 275, 19),
	(52834, 'SIMPANG NANENG', 3981, 275, 19),
	(52835, 'WURAN', 3981, 275, 19),
	(52836, 'BALAWA', 3982, 275, 19),
	(52837, 'JURU BANU', 3982, 275, 19),
	(52838, 'KALI NAPU', 3982, 275, 19),
	(52839, 'MAIPE', 3982, 275, 19),
	(52840, 'MURUTUWU', 3982, 275, 19),
	(52841, 'SIONG', 3982, 275, 19),
	(52842, 'TAMPULANGIT', 3982, 275, 19),
	(52843, 'TELANG BARU', 3982, 275, 19),
	(52844, 'TELANG SIONG', 3982, 275, 19),
	(52845, 'BANTAI NAPU', 3983, 275, 19),
	(52846, 'GANDRUNG', 3983, 275, 19),
	(52847, 'KALAMUS', 3983, 275, 19),
	(52848, 'KUPANG BARU', 3983, 275, 19),
	(52849, 'LUAU JAWUK', 3983, 275, 19),
	(52850, 'PAKU BETO', 3983, 275, 19),
	(52851, 'PANGKAN', 3983, 275, 19),
	(52852, 'PATUNG', 3983, 275, 19),
	(52853, 'RUNGGU RAYA', 3983, 275, 19),
	(52854, 'SIMPANG BINGKUANG/BANGKUANG', 3983, 275, 19),
	(52855, 'TAMPA', 3983, 275, 19),
	(52856, 'TARINGSING', 3983, 275, 19),
	(52857, 'AMPARI BORA', 3984, 275, 19),
	(52858, 'BENTOT', 3984, 275, 19),
	(52859, 'BETANG NALONG', 3984, 275, 19),
	(52860, 'JANGO', 3984, 275, 19),
	(52861, 'KAMBITIN', 3984, 275, 19),
	(52862, 'KOTAM', 3984, 275, 19),
	(52863, 'LALAP', 3984, 275, 19),
	(52864, 'MUANI', 3984, 275, 19),
	(52865, 'PULAU PADANG', 3984, 275, 19),
	(52866, 'RAMANIA', 3984, 275, 19),
	(52867, 'TAMIANG', 3984, 275, 19),
	(52868, 'BAMBULUNG', 3985, 275, 19),
	(52869, 'BARARAWA', 3985, 275, 19),
	(52870, 'KETAP', 3985, 275, 19),
	(52871, 'KUPANG BERSIH', 3985, 275, 19),
	(52872, 'LAMPEONG', 3985, 275, 19),
	(52873, 'LEBO', 3985, 275, 19),
	(52874, 'MUARA PALANTAU', 3985, 275, 19),
	(52875, 'MURU DUYUNG', 3985, 275, 19),
	(52876, 'NAGALEAH', 3985, 275, 19),
	(52877, 'PINANG TUNGGAL', 3985, 275, 19),
	(52878, 'SUMBEREJO', 3985, 275, 19),
	(52879, 'TUMPUNG ULUNG', 3985, 275, 19),
	(52880, 'TUYAU', 3985, 275, 19),
	(52881, 'BARUYAN', 3986, 275, 19),
	(52882, 'BATUAH', 3986, 275, 19),
	(52883, 'LENGGANG', 3986, 275, 19),
	(52884, 'MALINTOT', 3986, 275, 19),
	(52885, 'PURI', 3986, 275, 19),
	(52886, 'SIBUNG', 3986, 275, 19),
	(52887, 'TANGKUM', 3986, 275, 19),
	(52888, 'TURAN AMIS', 3986, 275, 19),
	(52889, 'UNSUM', 3986, 275, 19),
	(52890, 'BAOK', 3987, 276, 19),
	(52891, 'BERONG', 3987, 276, 19),
	(52892, 'LAMPEONG I', 3987, 276, 19),
	(52893, 'LAMPEONG II', 3987, 276, 19),
	(52894, 'LAWARANG', 3987, 276, 19),
	(52895, 'LINON BESI I', 3987, 276, 19),
	(52896, 'LINON BESI II', 3987, 276, 19),
	(52897, 'MUARA MEA', 3987, 276, 19),
	(52898, 'PAYANG', 3987, 276, 19),
	(52899, 'TAMBABA', 3987, 276, 19),
	(52900, 'TANJUNG HARAPAN', 3987, 276, 19),
	(52901, 'BALITI', 3988, 276, 19),
	(52902, 'BATURAYA I', 3988, 276, 19),
	(52903, 'BATURAYA II', 3988, 276, 19),
	(52904, 'JAMAN', 3988, 276, 19),
	(52905, 'KANDUI', 3988, 276, 19),
	(52906, 'KETAPANG', 3988, 276, 19),
	(52907, 'MAJANGKAN', 3988, 276, 19),
	(52908, 'MALUNGAI', 3988, 276, 19),
	(52909, 'PAYANG ARA', 3988, 276, 19),
	(52910, 'PELARI', 3988, 276, 19),
	(52911, 'RARAWA', 3988, 276, 19),
	(52912, 'SANGKURANG', 3988, 276, 19),
	(52913, 'SIWAU', 3988, 276, 19),
	(52914, 'TAPEN RAYA', 3988, 276, 19),
	(52915, 'TONGKA', 3988, 276, 19),
	(52916, 'WALUR', 3988, 276, 19),
	(52917, 'BENGAHON', 3989, 276, 19),
	(52918, 'HARAGANDANG', 3989, 276, 19),
	(52919, 'HURUNG ENEP', 3989, 276, 19),
	(52920, 'IPU', 3989, 276, 19),
	(52921, 'JUJU BARU', 3989, 276, 19),
	(52922, 'KARENDAN', 3989, 276, 19),
	(52923, 'LAHEI I', 3989, 276, 19),
	(52924, 'LAHEI II', 3989, 276, 19),
	(52925, 'MUARA BAKAH', 3989, 276, 19),
	(52926, 'MUARA INU', 3989, 276, 19),
	(52927, 'MUARA PARI', 3989, 276, 19),
	(52928, 'MUKUT', 3989, 276, 19),
	(52929, 'RAHADEN', 3989, 276, 19),
	(52930, 'BENAO HILIR', 3990, 276, 19),
	(52931, 'BENAO HULU', 3990, 276, 19),
	(52932, 'JANGKANG BARU', 3990, 276, 19),
	(52933, 'JANGKANG LAMA', 3990, 276, 19),
	(52934, 'KARAMUAN', 3990, 276, 19),
	(52935, 'LUWE HILIR', 3990, 276, 19),
	(52936, 'LUWE HULU', 3990, 276, 19),
	(52937, 'NIHAN HILIR', 3990, 276, 19),
	(52938, 'NIHAN HULU', 3990, 276, 19),
	(52939, 'PAPAR PUJUNG', 3990, 276, 19),
	(52940, 'TELUK MALEWAI', 3990, 276, 19),
	(52941, 'KAMAWEN', 3991, 276, 19),
	(52942, 'MONTALLAT I', 3991, 276, 19),
	(52943, 'MONTALLAT II', 3991, 276, 19),
	(52944, 'PARING LAHUNG', 3991, 276, 19),
	(52945, 'PEPAS', 3991, 276, 19),
	(52946, 'RUBEI', 3991, 276, 19),
	(52947, 'RUJI', 3991, 276, 19),
	(52948, 'SIKAN', 3991, 276, 19),
	(52949, 'TUMPUNG LAUNG I', 3991, 276, 19),
	(52950, 'TUMPUNG LAUNG II', 3991, 276, 19),
	(52951, 'GANDRING', 3992, 276, 19),
	(52952, 'HAJAK', 3992, 276, 19),
	(52953, 'JAMBU', 3992, 276, 19),
	(52954, 'JINGAH', 3992, 276, 19),
	(52955, 'LIANG NAGA', 3992, 276, 19),
	(52956, 'LIANGBUAH', 3992, 276, 19),
	(52957, 'MALAWAKEN', 3992, 276, 19),
	(52958, 'PANAEN', 3992, 276, 19),
	(52959, 'SABUH', 3992, 276, 19),
	(52960, 'SIKUI', 3992, 276, 19),
	(52961, 'BINTANG NINGGI I (SATU)', 3993, 276, 19),
	(52962, 'BINTANG NINGGI II (DUA)', 3993, 276, 19),
	(52963, 'BUKIT SAWIT', 3993, 276, 19),
	(52964, 'BUNTOK BARU', 3993, 276, 19),
	(52965, 'BUTONG', 3993, 276, 19),
	(52966, 'PANDRAN PERMAI', 3993, 276, 19),
	(52967, 'PANDRAN RAYA', 3993, 276, 19),
	(52968, 'TAWAN JAYA', 3993, 276, 19),
	(52969, 'TERINSING/TRINSING', 3993, 276, 19),
	(52970, 'TRAHEAN', 3993, 276, 19),
	(52971, 'BERINGIN RAYA', 3994, 276, 19),
	(52972, 'DATAI NIRUI', 3994, 276, 19),
	(52973, 'LANJAS', 3994, 276, 19),
	(52974, 'LEMO I', 3994, 276, 19),
	(52975, 'LEMO II', 3994, 276, 19),
	(52976, 'MELAYU', 3994, 276, 19),
	(52977, 'PENDREH/PANDREH', 3994, 276, 19),
	(52978, 'RIMBA SARI', 3994, 276, 19),
	(52979, 'SEI/SUNGAI RAHAYU I', 3994, 276, 19),
	(52980, 'SEI/SUNGAI RAHAYU II', 3994, 276, 19),
	(52981, 'BENANGIN I / I', 3995, 276, 19),
	(52982, 'BENANGIN II / 2', 3995, 276, 19),
	(52983, 'BENANGIN III / 3', 3995, 276, 19),
	(52984, 'BENANGIN V / 5', 3995, 276, 19),
	(52985, 'JAMUT', 3995, 276, 19),
	(52986, 'LIJU', 3995, 276, 19),
	(52987, 'MAMPUAK I', 3995, 276, 19),
	(52988, 'MAMPUAK II (PANTUNG)', 3995, 276, 19),
	(52989, 'MUARA WAKAT', 3995, 276, 19),
	(52990, 'SAMPIRANG I', 3995, 276, 19),
	(52991, 'SAMPIRANG II', 3995, 276, 19),
	(52992, 'SEI/SUNGAI LIJU', 3995, 276, 19),
	(52993, 'KARETAU RAMBANGUN', 3996, 277, 19),
	(52994, 'KARETAU SARIAN', 3996, 277, 19),
	(52995, 'LAWANG KANJI', 3996, 277, 19),
	(52996, 'TUMBANG ANOI', 3996, 277, 19),
	(52997, 'TUMBANG MAHUROI', 3996, 277, 19),
	(52998, 'TUMBANG MARAYA', 3996, 277, 19),
	(52999, 'TUMBANG MARIKOI', 3996, 277, 19),
	(53000, 'TUMBANG POSU', 3996, 277, 19)
	`)
}
