package kel0

import "github.com/danangkonang/rest-api/config"

func Kel2() {
	db := config.Connect()
	db.Exec(`INSERT INTO kelurahan (id,kelurahan,kecamatan_id,kota_id,provinsi_id) VALUES
		(2001, 'SUKAMANAH', 229, 14, 3),
		(2002, 'SUKARESMI', 229, 14, 3),
		(2003, 'BOJONGHALEUANG', 230, 14, 3),
		(2004, 'CIKANDE', 230, 14, 3),
		(2005, 'CIPANGERAN', 230, 14, 3),
		(2006, 'GIRIMUKTI', 230, 14, 3),
		(2007, 'JATI', 230, 14, 3),
		(2008, 'SAGULING', 230, 14, 3),
		(2009, 'BUNINAGARA', 231, 14, 3),
		(2010, 'CICANGKANG GIRANG', 231, 14, 3),
		(2011, 'CIKADU', 231, 14, 3),
		(2012, 'CINTAKARYA', 231, 14, 3),
		(2013, 'MEKARWANGI', 231, 14, 3),
		(2014, 'PASIRPOGOR', 231, 14, 3),
		(2015, 'PUNCAKSARI', 231, 14, 3),
		(2016, 'RANCA SENGGANG', 231, 14, 3),
		(2017, 'SINDANGKERTA', 231, 14, 3),
		(2018, 'WANGUNSARI', 231, 14, 3),
		(2019, 'WENINGGALIH', 231, 14, 3),
		(2020, 'ALUH ALUH KECIL', 232, 15, 3),
		(2021, 'ALUH ALUH KECIL MUARA', 232, 15, 3),
		(2022, 'ALUH-ALUH BESAR', 232, 15, 3),
		(2023, 'BAKAMBAT', 232, 15, 3),
		(2024, 'BALIMAU', 232, 15, 3),
		(2025, 'BUNIPAH', 232, 15, 3),
		(2026, 'HANDIL BARU', 232, 15, 3),
		(2027, 'HANDIL BUJUR', 232, 15, 3),
		(2028, 'KUIN BESAR', 232, 15, 3),
		(2029, 'KUIN KECIL', 232, 15, 3),
		(2030, 'LABAT MUARA', 232, 15, 3),
		(2031, 'PEMURUS', 232, 15, 3),
		(2032, 'PODOK', 232, 15, 3),
		(2033, 'PULANTAN', 232, 15, 3),
		(2034, 'SEI/SUNGAI MUSANG', 232, 15, 3),
		(2035, 'SIMPANG WARGA', 232, 15, 3),
		(2036, 'SIMPANG WARGA DALAM', 232, 15, 3),
		(2037, 'TANIPAH', 232, 15, 3),
		(2038, 'TERAPU', 232, 15, 3),
		(2039, 'APUAI', 233, 15, 3),
		(2040, 'ARANIO', 233, 15, 3),
		(2041, 'ARTAIN', 233, 15, 3),
		(2042, 'BELANGIAN', 233, 15, 3),
		(2043, 'BENUA RIAM', 233, 15, 3),
		(2044, 'BUNGLAI', 233, 15, 3),
		(2045, 'KALAAN', 233, 15, 3),
		(2046, 'PAAU', 233, 15, 3),
		(2047, 'RANTAU BALAI', 233, 15, 3),
		(2048, 'RANTAU BUJUR', 233, 15, 3),
		(2049, 'TIWINGAN BARU', 233, 15, 3),
		(2050, 'TIWINGAN LAMA', 233, 15, 3),
		(2051, 'ASTAMBUL', 234, 15, 3),
		(2052, 'ASTAMBUL SEBERANG', 234, 15, 3),
		(2053, 'BANUA ANYAR I (BENUA ANYAR I)', 234, 15, 3),
		(2054, 'BANUA ANYAR II (BENUA ANYAR II)', 234, 15, 3),
		(2055, 'DANAU SALAK', 234, 15, 3),
		(2056, 'JATI', 234, 15, 3),
		(2057, 'KALAMPAYAN', 234, 15, 3),
		(2058, 'KALAMPAYAN ULU', 234, 15, 3),
		(2059, 'KALIUKAN', 234, 15, 3),
		(2060, 'KELAMPAIAN ILIR', 234, 15, 3),
		(2061, 'LIMAMAR', 234, 15, 3),
		(2062, 'LOKGABANG', 234, 15, 3),
		(2063, 'MUNGGU RAYA', 234, 15, 3),
		(2064, 'PASAR JATI', 234, 15, 3),
		(2065, 'PEMATANG HAMBAWANG', 234, 15, 3),
		(2066, 'PINGARAN ILIR', 234, 15, 3),
		(2067, 'PINGARAN ULU', 234, 15, 3),
		(2068, 'SEI/SUNGAI ALAT', 234, 15, 3),
		(2069, 'SEI/SUNGAI TUAN ILIR', 234, 15, 3),
		(2070, 'SEI/SUNGAI TUAN ULU', 234, 15, 3),
		(2071, 'TAMBAK DANAU', 234, 15, 3),
		(2072, 'TAMBANGAN', 234, 15, 3),
		(2073, 'BABIRIK', 235, 15, 3),
		(2074, 'HANDIL PURAI', 235, 15, 3),
		(2075, 'HAUR KUNING', 235, 15, 3),
		(2076, 'JAMBU BURUNG', 235, 15, 3),
		(2077, 'JAMBU RAYA', 235, 15, 3),
		(2078, 'KAMPUNG BARU', 235, 15, 3),
		(2079, 'LAWAHAN', 235, 15, 3),
		(2080, 'MUARA HALAYUNG', 235, 15, 3),
		(2081, 'PINDAHAN BARU', 235, 15, 3),
		(2082, 'RUMPIANG', 235, 15, 3),
		(2083, 'SELAT MAKMUR', 235, 15, 3),
		(2084, 'TAMBAK PADI', 235, 15, 3),
		(2085, 'BANYU HIRANG', 236, 15, 3),
		(2086, 'GAMBUT', 236, 15, 3),
		(2087, 'GAMBUT BARAT', 236, 15, 3),
		(2088, 'GUNTUNG PAPUYU', 236, 15, 3),
		(2089, 'GUNTUNG UJUNG', 236, 15, 3),
		(2090, 'KAYU BAWANG', 236, 15, 3),
		(2091, 'KELADAN BARU', 236, 15, 3),
		(2092, 'MAKMUR', 236, 15, 3),
		(2093, 'MALINTANG', 236, 15, 3),
		(2094, 'MALINTANG BARU', 236, 15, 3),
		(2095, 'SEI/SUNGAI KUPANG', 236, 15, 3),
		(2096, 'TAMBAK SIRANG BARU', 236, 15, 3),
		(2097, 'TAMBAK SIRANG DARAT', 236, 15, 3),
		(2098, 'TAMBAK SIRANG LAUT', 236, 15, 3),
		(2099, 'ABIRAU', 237, 15, 3),
		(2100, 'AWANG BANGKAL BARAT', 237, 15, 3),
		(2101, 'AWANG BANGKAL TIMUR', 237, 15, 3),
		(2102, 'BALAU', 237, 15, 3),
		(2103, 'BIIH', 237, 15, 3),
		(2104, 'JINGAH HABANG ILIR', 237, 15, 3),
		(2105, 'JINGAH HABANG ULU', 237, 15, 3),
		(2106, 'KARANG INTAN', 237, 15, 3),
		(2107, 'KIRAM', 237, 15, 3),
		(2108, 'LIHUNG', 237, 15, 3),
		(2109, 'LOKTANGGA', 237, 15, 3),
		(2110, 'MALI MALI', 237, 15, 3),
		(2111, 'MANDI ANGIN BARAT', 237, 15, 3),
		(2112, 'MANDI ANGIN TIMUR', 237, 15, 3),
		(2113, 'MANDI KAPAU BARAT', 237, 15, 3),
		(2114, 'MANDI KAPAU TIMUR', 237, 15, 3),
		(2115, 'PADANG PANJANG', 237, 15, 3),
		(2116, 'PANDAK DAUN', 237, 15, 3),
		(2117, 'PASAR LAMA', 237, 15, 3),
		(2118, 'PENYAMBARAN', 237, 15, 3),
		(2119, 'PULAU NYIUR', 237, 15, 3),
		(2120, 'SEI/SUNGAI ALANG', 237, 15, 3),
		(2121, 'SEI/SUNGAI ARFAT', 237, 15, 3),
		(2122, 'SEI/SUNGAI ASAM', 237, 15, 3),
		(2123, 'SEI/SUNGAI BESAR', 237, 15, 3),
		(2124, 'SEI/SUNGAI LANDAS', 237, 15, 3),
		(2125, 'BELAYUNG BARU', 238, 15, 3),
		(2126, 'BENUA HANYAR (BANUA HANYAR)', 238, 15, 3),
		(2127, 'KERTAK HANYAR I', 238, 15, 3),
		(2128, 'KERTAKHANYAR II', 238, 15, 3),
		(2129, 'MANARAP BARU', 238, 15, 3),
		(2130, 'MANARAP LAMA', 238, 15, 3),
		(2131, 'MANARAP TENGAH', 238, 15, 3),
		(2132, 'MANDAR SARI', 238, 15, 3),
		(2133, 'MEKAR RAYA', 238, 15, 3),
		(2134, 'PASAR KEMIS', 238, 15, 3),
		(2135, 'PEMANGKIH BARU', 238, 15, 3),
		(2136, 'SEI/SUNGAI LAKUM', 238, 15, 3),
		(2137, 'SIMPANG EMPAT', 238, 15, 3),
		(2138, 'BOJONGKANTONG', 239, 15, 3),
		(2139, 'KUJANGSARI', 239, 15, 3),
		(2140, 'LANGENSARI', 239, 15, 3),
		(2141, 'MUKTISARI', 239, 15, 3),
		(2142, 'REJASARI', 239, 15, 3),
		(2143, 'WARINGINSARI', 239, 15, 3),
		(2144, 'ANTASAN SUTUN', 240, 15, 3),
		(2145, 'KELILING BENTENG TENGAH', 240, 15, 3),
		(2146, 'KELILING BENTENG ULU', 240, 15, 3),
		(2147, 'PENGGALAMAN', 240, 15, 3),
		(2148, 'SEI/SUNGAI BATANG', 240, 15, 3),
		(2149, 'SEI/SUNGAI BATANG ILIR', 240, 15, 3),
		(2150, 'SEI/SUNGAI RANGAS', 240, 15, 3),
		(2151, 'SEI/SUNGAI RANGAS HAMBUKU', 240, 15, 3),
		(2152, 'SEI/SUNGAI RANGAS TENGAH', 240, 15, 3),
		(2153, 'SEI/SUNGAI RANGAS ULU', 240, 15, 3),
		(2154, 'TANGKAS', 240, 15, 3),
		(2155, 'TELOK SELONG', 240, 15, 3),
		(2156, 'TELOK SELONG ULU', 240, 15, 3),
		(2157, 'BINCAU', 241, 15, 3),
		(2158, 'BINCAU MUARA', 241, 15, 3),
		(2159, 'CINDAI ALUS', 241, 15, 3),
		(2160, 'INDRA SARI', 241, 15, 3),
		(2161, 'JAWA', 241, 15, 3),
		(2162, 'JAWA LAUT', 241, 15, 3),
		(2163, 'KERATON', 241, 15, 3),
		(2164, 'LABUAN TABU', 241, 15, 3),
		(2165, 'MURUNG KENANGA', 241, 15, 3),
		(2166, 'MURUNG KERATON', 241, 15, 3),
		(2167, 'PASAYANGAN BARAT', 241, 15, 3),
		(2168, 'PASAYANGAN SELATAN', 241, 15, 3),
		(2169, 'PASAYANGAN UTARA', 241, 15, 3),
		(2170, 'PESAYANGAN', 241, 15, 3),
		(2171, 'SEI/SUNGAI PARING', 241, 15, 3),
		(2172, 'SEI/SUNGAI SIPAI', 241, 15, 3),
		(2173, 'SEKUMPUL', 241, 15, 3),
		(2174, 'TAMBAK BARU', 241, 15, 3),
		(2175, 'TAMBAK BARU ILIR', 241, 15, 3),
		(2176, 'TAMBAK BARU ULU', 241, 15, 3),
		(2177, 'TANJUNG REMA', 241, 15, 3),
		(2178, 'TANJUNG REMA DARAT', 241, 15, 3),
		(2179, 'TUNGGUL IRANG', 241, 15, 3),
		(2180, 'TUNGGUL IRANG ILIR', 241, 15, 3),
		(2181, 'TUNGGUL IRANG ULU', 241, 15, 3),
		(2182, 'TUNGKARAN', 241, 15, 3),
		(2183, 'AKAR BAGANTUNG', 242, 15, 3),
		(2184, 'AKAR BARU', 242, 15, 3),
		(2185, 'ANTASAN SENOR', 242, 15, 3),
		(2186, 'ANTASAN SENOR ILIR', 242, 15, 3),
		(2187, 'DALAM PAGAR', 242, 15, 3),
		(2188, 'DALAM PAGAR ULU', 242, 15, 3),
		(2189, 'KERAMAT', 242, 15, 3),
		(2190, 'KERAMAT BARU', 242, 15, 3),
		(2191, 'MEKAR', 242, 15, 3),
		(2192, 'MELAYU ILIR', 242, 15, 3),
		(2193, 'MELAYU TENGAH', 242, 15, 3),
		(2194, 'MELAYU ULU', 242, 15, 3),
		(2195, 'PEKAUMAN', 242, 15, 3),
		(2196, 'PEKAUMAN DALAM', 242, 15, 3),
		(2197, 'PEKAUMAN ULU', 242, 15, 3),
		(2198, 'PEMATANG BARU', 242, 15, 3),
		(2199, 'SEI/SUNGAI KITANO', 242, 15, 3),
		(2200, 'TAMBAK ANYAR', 242, 15, 3),
		(2201, 'TAMBAK ANYAR ILIR', 242, 15, 3),
		(2202, 'TAMBAK ANYAR ULU', 242, 15, 3),
		(2203, 'BARU', 243, 15, 3),
		(2204, 'BAWAHAN PASAR', 243, 15, 3),
		(2205, 'BAWAHAN SEBERANG', 243, 15, 3),
		(2206, 'BAWAHAN SELAN', 243, 15, 3),
		(2207, 'GUNUNG ULIN', 243, 15, 3),
		(2208, 'LOK TAMU', 243, 15, 3),
		(2209, 'MANGKA LAWAT', 243, 15, 3),
		(2210, 'MATARAMAN', 243, 15, 3),
		(2211, 'PASIRAMAN', 243, 15, 3),
		(2212, 'PEMATANG DANAU', 243, 15, 3),
		(2213, 'SEI/SUNGAI JATI', 243, 15, 3),
		(2214, 'SIMPANG TIGA', 243, 15, 3),
		(2215, 'SURIAN', 243, 15, 3),
		(2216, 'TAKUTI', 243, 15, 3),
		(2217, 'TANAH ABANG', 243, 15, 3),
		(2218, 'BATULAWANG', 244, 15, 3),
		(2219, 'BINANGUN', 244, 15, 3),
		(2220, 'HEGARSARI', 244, 15, 3),
		(2221, 'KARYAMUKTI', 244, 15, 3),
		(2222, 'MULYASARI', 244, 15, 3),
		(2223, 'PATARUMAN', 244, 15, 3),
		(2224, 'SINARTANJUNG', 244, 15, 3),
		(2225, 'SUKAMUKTI', 244, 15, 3),
		(2226, 'ALIMUKIM', 245, 15, 3),
		(2227, 'ANTARAKU', 245, 15, 3),
		(2228, 'ATI IM', 245, 15, 3),
		(2229, 'BENTENG', 245, 15, 3),
		(2230, 'KERTAK EMPAT', 245, 15, 3),
		(2231, 'LOBANG BARU', 245, 15, 3),
		(2232, 'LOKTUNGGUL', 245, 15, 3),
		(2233, 'LUMPANGI', 245, 15, 3),
		(2234, 'MANGKAUK', 245, 15, 3),
		(2235, 'MANIAPUN', 245, 15, 3),
		(2236, 'PANYIURAN', 245, 15, 3),
		(2237, 'PENGARON', 245, 15, 3),
		(2238, 'ANGKIPIH', 246, 15, 3),
		(2239, 'PERAMASAN ATAS', 246, 15, 3),
		(2240, 'PERAMASAN BAWAH', 246, 15, 3),
		(2241, 'REMO', 246, 15, 3),
		(2242, 'KARANGPANIMBAL', 247, 15, 3),
		(2243, 'MEKARHARJA', 247, 15, 3),
		(2244, 'PURWAHARJA', 247, 15, 3),
		(2245, 'RAHARJA', 247, 15, 3),
		(2246, 'BALIANGIN', 248, 15, 3),
		(2247, 'BATANG BANYU', 248, 15, 3),
		(2248, 'BATU TANAM', 248, 15, 3),
		(2249, 'GUNUNG BATU', 248, 15, 3),
		(2250, 'MADUREJO', 248, 15, 3),
		(2251, 'PASAR BARU', 248, 15, 3),
		(2252, 'SEI/SUNGAI LURUS', 248, 15, 3),
		(2253, 'BELIMBING BARU', 249, 15, 3),
		(2254, 'BELIMBING LAMA', 249, 15, 3),
		(2255, 'HAKIM MAKMUR', 249, 15, 3),
		(2256, 'KAHELAAN', 249, 15, 3),
		(2257, 'KUPANG REJO', 249, 15, 3),
		(2258, 'PAKUTIK', 249, 15, 3),
		(2259, 'RANTAU BAKULA', 249, 15, 3),
		(2260, 'RANTAU NANGKA', 249, 15, 3),
		(2261, 'SEI/SUNGAI PINANG', 249, 15, 3),
		(2262, 'SUMBER BARU', 249, 15, 3),
		(2263, 'SUMBER HARAPAN', 249, 15, 3),
		(2264, 'ABUMBUN JAYA', 250, 15, 3),
		(2265, 'GUDANG HIRANG', 250, 15, 3),
		(2266, 'GUDANG TENGAH', 250, 15, 3),
		(2267, 'KELILING BENTENG ILIR', 250, 15, 3),
		(2268, 'LOK BAINTAN', 250, 15, 3),
		(2269, 'LOK BAINTAN DALAM', 250, 15, 3),
		(2270, 'LOKBUNTAR', 250, 15, 3),
		(2271, 'PAKU ALAM', 250, 15, 3),
		(2272, 'PEJAMBUAN', 250, 15, 3),
		(2273, 'PEMAKUAN', 250, 15, 3),
		(2274, 'PEMATANG PANJANG', 250, 15, 3),
		(2275, 'PEMBANTANAN', 250, 15, 3),
		(2276, 'SEI/SUNGAI BAKUNG', 250, 15, 3),
		(2277, 'SEI/SUNGAI BANGKAL', 250, 15, 3),
		(2278, 'SEI/SUNGAI LULUT', 250, 15, 3),
		(2279, 'SEI/SUNGAI PINANG BARU', 250, 15, 3),
		(2280, 'SEI/SUNGAI PINANG LAMA', 250, 15, 3),
		(2281, 'SEI/SUNGAI TABUK KERAMAT', 250, 15, 3),
		(2282, 'SEI/SUNGAI TABUK KOTA', 250, 15, 3),
		(2283, 'SEI/SUNGAI TANDIPAH', 250, 15, 3),
		(2284, 'TAJAU LANDUNG', 250, 15, 3),
		(2285, 'BANGKAL TENGAH', 251, 15, 3),
		(2286, 'JARUJU', 251, 15, 3),
		(2287, 'JARUJU LAUT', 251, 15, 3),
		(2288, 'LAYAP BARU', 251, 15, 3),
		(2289, 'MEKAR SARI', 251, 15, 3),
		(2290, 'PANDAN SARI', 251, 15, 3),
		(2291, 'PEMANGKIH BARU', 251, 15, 3),
		(2292, 'PEMANGKIH DARAT', 251, 15, 3),
		(2293, 'PEMANGKIH TENGAH', 251, 15, 3),
		(2294, 'TAIBAH RAYA', 251, 15, 3),
		(2295, 'TAMPANG AWANG', 251, 15, 3),
		(2296, 'TATAH BANGKAL', 251, 15, 3),
		(2297, 'TATAH LAYAP', 251, 15, 3),
		(2298, 'LOKTANAH', 252, 15, 3),
		(2299, 'RAMPAH', 252, 15, 3),
		(2300, 'RANTAU BUJUR', 252, 15, 3),
		(2301, 'TELAGA BARU', 252, 15, 3),
		(2302, 'BABELAN KOTA', 253, 16, 3),
		(2303, 'BAHAGIA', 253, 16, 3),
		(2304, 'BUNI BAKTI', 253, 16, 3),
		(2305, 'HURIP JAYA', 253, 16, 3),
		(2306, 'KEBALEN', 253, 16, 3),
		(2307, 'KEDUNG PENGAWAS', 253, 16, 3),
		(2308, 'KEDUNGJAYA', 253, 16, 3),
		(2309, 'MUARA BAKTI', 253, 16, 3),
		(2310, 'PANTAI HURIP', 253, 16, 3),
		(2311, 'BANTAR GEBANG', 254, 16, 3),
		(2312, 'CIKETING UDIK', 254, 16, 3),
		(2313, 'CIKIWUL', 254, 16, 3),
		(2314, 'SUMUR BATU', 254, 16, 3),
		(2315, 'BINTARA', 255, 16, 3),
		(2316, 'BINTARA JAYA', 255, 16, 3),
		(2317, 'JAKA SAMPURNA', 255, 16, 3),
		(2318, 'KOTA BARU', 255, 16, 3),
		(2319, 'KRANJI', 255, 16, 3),
		(2320, 'JAKA MULYA', 256, 16, 3),
		(2321, 'JAKA SETIA', 256, 16, 3),
		(2322, 'KAYURINGIN JAYA', 256, 16, 3),
		(2323, 'MARGAJAYA', 256, 16, 3),
		(2324, 'PEKAYON JAYA', 256, 16, 3),
		(2325, 'AREN JAYA', 257, 16, 3),
		(2326, 'BEKASI JAYA', 257, 16, 3),
		(2327, 'DUREN JAYA', 257, 16, 3),
		(2328, 'MARGAHAYU', 257, 16, 3),
		(2329, 'HARAPAN BARU', 258, 16, 3),
		(2330, 'HARAPAN JAYA', 258, 16, 3),
		(2331, 'KALI ABANG TENGAH', 258, 16, 3),
		(2332, 'MARGA MULYA', 258, 16, 3),
		(2333, 'PERWIRA', 258, 16, 3),
		(2334, 'TELUK PUCUNG', 258, 16, 3),
		(2335, 'BOJONGMANGGU', 259, 16, 3),
		(2336, 'KARANGINDAH', 259, 16, 3),
		(2337, 'KARANGMULYA', 259, 16, 3),
		(2338, 'MEDALKRISNA', 259, 16, 3),
		(2339, 'SUKABUNGAH', 259, 16, 3),
		(2340, 'SUKAMUKTI', 259, 16, 3),
		(2341, 'JAYA BAKTI', 260, 16, 3),
		(2342, 'JAYA LAKSANA', 260, 16, 3),
		(2343, 'LENGGAH JAYA', 260, 16, 3),
		(2344, 'LENGGAH SARI', 260, 16, 3),
		(2345, 'SETIA JAYA', 260, 16, 3),
		(2346, 'SETIALAKSANA', 260, 16, 3),
		(2347, 'SINDANG JAYA', 260, 16, 3),
		(2348, 'SINDANGSARI', 260, 16, 3),
		(2349, 'CIBARUSAHJAYA', 261, 16, 3),
		(2350, 'CIBARUSAHKOTA', 261, 16, 3),
		(2351, 'RIDOGALIH', 261, 16, 3),
		(2352, 'RIDOMANAH', 261, 16, 3),
		(2353, 'SINDANGMULYA', 261, 16, 3),
		(2354, 'SIRNAJATI', 261, 16, 3),
		(2355, 'WIBAWAMULYA', 261, 16, 3),
		(2356, 'CIKEDOKAN', 262, 16, 3),
		(2357, 'DANAU INDAH', 262, 16, 3),
		(2358, 'GANDAMEKAR', 262, 16, 3),
		(2359, 'GANDASARI', 262, 16, 3),
		(2360, 'JATIWANGI', 262, 16, 3),
		(2361, 'KALI JAYA', 262, 16, 3),
		(2362, 'MEKAR WANGI', 262, 16, 3),
		(2363, 'SUKADANAU', 262, 16, 3),
		(2364, 'TELAGA ASIH', 262, 16, 3),
		(2365, 'TELAGA MURNI', 262, 16, 3),
		(2366, 'TELAJUNG', 262, 16, 3),
		(2367, 'CICAU', 263, 16, 3),
		(2368, 'HEGARMUKTI', 263, 16, 3),
		(2369, 'JAYAMUKTI', 263, 16, 3),
		(2370, 'PASIRPANJI', 263, 16, 3),
		(2371, 'PASIRTANJUNG', 263, 16, 3),
		(2372, 'SUKAMAHI', 263, 16, 3),
		(2373, 'CIANTRA', 264, 16, 3),
		(2374, 'CIBATU', 264, 16, 3),
		(2375, 'PASIRSARI', 264, 16, 3),
		(2376, 'SERANG', 264, 16, 3),
		(2377, 'SUKADAMI', 264, 16, 3),
		(2378, 'SUKARESMI', 264, 16, 3),
		(2379, 'SUKASEJATI', 264, 16, 3),
		(2380, 'CIPAYUNG', 265, 16, 3),
		(2381, 'HEGARMANAH', 265, 16, 3),
		(2382, 'JATIBARU', 265, 16, 3),
		(2383, 'JATIREJA', 265, 16, 3),
		(2384, 'KARANGSARI', 265, 16, 3),
		(2385, 'LABANSARI', 265, 16, 3),
		(2386, 'SERTAJAYA', 265, 16, 3),
		(2387, 'TANJUNGBARU', 265, 16, 3),
		(2388, 'CIKARANG KOTA', 266, 16, 3),
		(2389, 'HARJA MEKAR', 266, 16, 3),
		(2390, 'KARANG BARU', 266, 16, 3),
		(2391, 'KARANGASIH', 266, 16, 3),
		(2392, 'KARANGRAHARJA', 266, 16, 3),
		(2393, 'MEKARMUKTI', 266, 16, 3),
		(2394, 'PASIR GOMBONG', 266, 16, 3),
		(2395, 'SIMPANGAN', 266, 16, 3),
		(2396, 'TANJUNGSARI', 266, 16, 3),
		(2397, 'WALUYA', 266, 16, 3),
		(2398, 'WANGUNHARJA', 266, 16, 3),
		(2399, 'JATIKARYA', 267, 16, 3),
		(2400, 'JATIRADEN', 267, 16, 3),
		(2401, 'JATIRANGGA', 267, 16, 3),
		(2402, 'JATIRANGGON', 267, 16, 3),
		(2403, 'JATISAMPURNA', 267, 16, 3),
		(2404, 'JATI MEKAR', 268, 16, 3),
		(2405, 'JATIASIH', 268, 16, 3),
		(2406, 'JATIKRAMAT', 268, 16, 3),
		(2407, 'JATILUHUR', 268, 16, 3),
		(2408, 'JATIRASA', 268, 16, 3),
		(2409, 'JATISARI', 268, 16, 3),
		(2410, 'KARANG BAHAGIA', 269, 16, 3),
		(2411, 'KARANG RAHAYU', 269, 16, 3),
		(2412, 'KARANG SENTOSA', 269, 16, 3),
		(2413, 'KARANG SETRA', 269, 16, 3),
		(2414, 'KARANG SETU', 269, 16, 3),
		(2415, 'KARANGANYAR', 269, 16, 3),
		(2416, 'KARANGMUKTI', 269, 16, 3),
		(2417, 'SUKARAYA', 269, 16, 3),
		(2418, 'BOJONGSARI', 270, 16, 3),
		(2419, 'KARANG HARUM', 270, 16, 3),
		(2420, 'KARANG MEKAR', 270, 16, 3),
		(2421, 'KARANG SAMBUNG', 270, 16, 3),
		(2422, 'KEDUNGWARINGIN', 270, 16, 3),
		(2423, 'MEKAR JAYA', 270, 16, 3),
		(2424, 'WARINGIN JAYA', 270, 16, 3),
		(2425, 'HARAPAN MULYA', 271, 16, 3),
		(2426, 'KALI BARU', 271, 16, 3),
		(2427, 'MEDAN SATRIA', 271, 16, 3),
		(2428, 'PEJUANG', 271, 16, 3),
		(2429, 'JAYASAKTI', 272, 16, 3),
		(2430, 'PANTAI BAHAGIA', 272, 16, 3),
		(2431, 'PANTAI BAKTI', 272, 16, 3),
		(2432, 'PANTAI HARAPAN JAYA', 272, 16, 3),
		(2433, 'PANTAI MEKAR', 272, 16, 3),
		(2434, 'PANTAI SEDERHANA', 272, 16, 3),
		(2435, 'CIMUNING', 273, 16, 3),
		(2436, 'MUSTIKA JAYA', 273, 16, 3),
		(2437, 'MUSTIKA SARI', 273, 16, 3),
		(2438, 'PEDURENAN (PADURENAN)', 273, 16, 3),
		(2439, 'BANTARJAYA', 274, 16, 3),
		(2440, 'BANTARSARI', 274, 16, 3),
		(2441, 'KARANG HARJA', 274, 16, 3),
		(2442, 'KARANGHAUR', 274, 16, 3),
		(2443, 'KARANGJAYA', 274, 16, 3),
		(2444, 'KARANGPATRI', 274, 16, 3),
		(2445, 'KARANGREJA', 274, 16, 3),
		(2446, 'KARANGSEGAR', 274, 16, 3),
		(2447, 'KERTAJAYA', 274, 16, 3),
		(2448, 'KERTASARI', 274, 16, 3),
		(2449, 'SUMBERREJA', 274, 16, 3),
		(2450, 'SUMBERSARI', 274, 16, 3),
		(2451, 'SUMBERURIP', 274, 16, 3),
		(2452, 'JATIBENING', 275, 16, 3),
		(2453, 'JATIBENING BARU', 275, 16, 3),
		(2454, 'JATICEMPAKA', 275, 16, 3),
		(2455, 'JATIMAKMUR', 275, 16, 3),
		(2456, 'JATIWARINGIN', 275, 16, 3),
		(2457, 'JATI RAHAYU', 276, 16, 3),
		(2458, 'JATIMELATI', 276, 16, 3),
		(2459, 'JATIMURNI', 276, 16, 3),
		(2460, 'JATIWARNA', 276, 16, 3),
		(2461, 'BOJONG MENTENG', 277, 16, 3),
		(2462, 'BOJONG RAWALUMBU', 277, 16, 3),
		(2463, 'PENGASINAN', 277, 16, 3),
		(2464, 'SEPANJANG JAYA', 277, 16, 3),
		(2465, 'CILANGKARA', 278, 16, 3),
		(2466, 'JAYA SAMPURNA', 278, 16, 3),
		(2467, 'JAYAMULYA', 278, 16, 3),
		(2468, 'NAGACIPTA', 278, 16, 3),
		(2469, 'NAGASARI', 278, 16, 3),
		(2470, 'SIRNAJAYA', 278, 16, 3),
		(2471, 'SUKARAGAM', 278, 16, 3),
		(2472, 'SUKASARI', 278, 16, 3),
		(2473, 'BABAKAN', 279, 16, 3),
		(2474, 'BAKTI JAYA', 279, 16, 3),
		(2475, 'BURANGKENG', 279, 16, 3),
		(2476, 'CIBENING', 279, 16, 3),
		(2477, 'CIJENGKOL', 279, 16, 3),
		(2478, 'CIKARAGEMAN', 279, 16, 3),
		(2479, 'CILEDUK', 279, 16, 3),
		(2480, 'KADEMANGAN', 279, 16, 3),
		(2481, 'KERTARAHAYU', 279, 16, 3),
		(2482, 'KRANGGAN', 279, 16, 3),
		(2483, 'LUBANGBUAYA', 279, 16, 3),
		(2484, 'MUKTIJAYA', 279, 16, 3),
		(2485, 'MUNCUL', 279, 16, 3),
		(2486, 'RAGEMANUNGGAL', 279, 16, 3),
		(2487, 'SETU', 279, 16, 3),
		(2488, 'TAMAN RAHAYU', 279, 16, 3),
		(2489, 'TAMAN SARI', 279, 16, 3),
		(2490, 'SUKABUDI', 280, 16, 3),
		(2491, 'SUKADAYA', 280, 16, 3),
		(2492, 'SUKAKERTA', 280, 16, 3),
		(2493, 'SUKAMEKAR', 280, 16, 3),
		(2494, 'SUKARINGIN', 280, 16, 3),
		(2495, 'SUKATENANG', 280, 16, 3),
		(2496, 'SUKAWANGI', 280, 16, 3),
		(2497, 'SUKABAKTI', 281, 16, 3),
		(2498, 'SUKAMAJU', 281, 16, 3),
		(2499, 'SUKAMANTRI', 281, 16, 3),
		(2500, 'SUKARAHAYU', 281, 16, 3),
		(2501, 'SUKARAJA', 281, 16, 3),
		(2502, 'SUKARAPIH', 281, 16, 3),
		(2503, 'SUKAWIJAYA', 281, 16, 3),
		(2504, 'JATIMULYA', 282, 16, 3),
		(2505, 'LAMBANGJAYA', 282, 16, 3),
		(2506, 'LAMBANGSARI', 282, 16, 3),
		(2507, 'MANGUNJAYA', 282, 16, 3),
		(2508, 'MEKARSARI', 282, 16, 3),
		(2509, 'SETIADARMA', 282, 16, 3),
		(2510, 'SETIAMEKAR', 282, 16, 3),
		(2511, 'SUMBER JAYA', 282, 16, 3),
		(2512, 'TAMBUN', 282, 16, 3),
		(2513, 'TRIDAYA SAKTI', 282, 16, 3),
		(2514, 'JALENJAYA (JEJALENJAYA)', 283, 16, 3),
		(2515, 'KARANGSATRIA', 283, 16, 3),
		(2516, 'SATRIAJAYA', 283, 16, 3),
		(2517, 'SATRIAMEKAR', 283, 16, 3),
		(2518, 'SRIAMUR', 283, 16, 3),
		(2519, 'SRIJAYA', 283, 16, 3),
		(2520, 'SRIMAHI', 283, 16, 3),
		(2521, 'SRIMUKTI', 283, 16, 3),
		(2522, 'PAHLAWAN SETIA', 284, 16, 3),
		(2523, 'PANTAI MAKMUR', 284, 16, 3),
		(2524, 'PUSAKA RAKYAT', 284, 16, 3),
		(2525, 'SAMUDRA JAYA', 284, 16, 3),
		(2526, 'SEGARA JAYA', 284, 16, 3),
		(2527, 'SEGARA MAKMUR', 284, 16, 3),
		(2528, 'SETIA ASIH', 284, 16, 3),
		(2529, 'SETIA MULYA', 284, 16, 3),
		(2530, 'BABAKAN MADANG', 285, 17, 3),
		(2531, 'BOJONG KONENG', 285, 17, 3),
		(2532, 'CIJAYANTI', 285, 17, 3),
		(2533, 'CIPAMBUAN', 285, 17, 3),
		(2534, 'CITARINGGUL', 285, 17, 3),
		(2535, 'KADUMANGU', 285, 17, 3),
		(2536, 'KARANG TENGAH', 285, 17, 3),
		(2537, 'SENTUL', 285, 17, 3),
		(2538, 'SUMUR BATU', 285, 17, 3),
		(2539, 'BALUNGBANG JAYA', 286, 17, 3),
		(2540, 'BUBULAK', 286, 17, 3),
		(2541, 'CILENDEK BARAT', 286, 17, 3),
		(2542, 'CILENDEK TIMUR', 286, 17, 3),
		(2543, 'CURUG', 286, 17, 3),
		(2544, 'CURUG MEKAR', 286, 17, 3),
		(2545, 'GUNUNGBATU', 286, 17, 3),
		(2546, 'LOJI', 286, 17, 3),
		(2547, 'MARGAJAYA', 286, 17, 3),
		(2548, 'MENTENG', 286, 17, 3),
		(2549, 'PASIR JAYA', 286, 17, 3),
		(2550, 'PASIR KUDA', 286, 17, 3),
		(2551, 'PASIR MULYA', 286, 17, 3),
		(2552, 'SEMPLAK', 286, 17, 3),
		(2553, 'SINDANGBARANG', 286, 17, 3),
		(2554, 'SITU GEDE', 286, 17, 3),
		(2555, 'BATUTULIS', 287, 17, 3),
		(2556, 'BOJONGKERTA', 287, 17, 3),
		(2557, 'BONDONGAN', 287, 17, 3),
		(2558, 'CIKARET', 287, 17, 3),
		(2559, 'CIPAKU', 287, 17, 3),
		(2560, 'EMPANG', 287, 17, 3),
		(2561, 'GENTENG', 287, 17, 3),
		(2562, 'HARJASARI', 287, 17, 3),
		(2563, 'KERTAMAYA', 287, 17, 3),
		(2564, 'LAWANG GINTUNG', 287, 17, 3),
		(2565, 'MUARASARI', 287, 17, 3),
		(2566, 'MULYAHARJA', 287, 17, 3),
		(2567, 'PAKUAN', 287, 17, 3),
		(2568, 'PAMOYANAN', 287, 17, 3),
		(2569, 'RANCAMAYA', 287, 17, 3),
		(2570, 'RANGGAMEKAR', 287, 17, 3),
		(2571, 'BABAKAN', 288, 17, 3),
		(2572, 'BABAKAN PASAR', 288, 17, 3),
		(2573, 'CIBOGOR', 288, 17, 3),
		(2574, 'CIWARINGIN', 288, 17, 3),
		(2575, 'GUDANG', 288, 17, 3),
		(2576, 'KEBON KELAPA', 288, 17, 3),
		(2577, 'PABATON', 288, 17, 3),
		(2578, 'PALEDANG', 288, 17, 3),
		(2579, 'PANARAGAN', 288, 17, 3),
		(2580, 'SEMPUR', 288, 17, 3),
		(2581, 'TEGAL PANJANG', 288, 17, 3),
		(2582, 'BARANANGSIANG', 289, 17, 3),
		(2583, 'KATULAMPA', 289, 17, 3),
		(2584, 'SINDANGRASA', 289, 17, 3),
		(2585, 'SINDANGSARI', 289, 17, 3),
		(2586, 'SUKASARI', 289, 17, 3),
		(2587, 'TAJUR', 289, 17, 3),
		(2588, 'BANTARJATI', 290, 17, 3),
		(2589, 'CIBULUH', 290, 17, 3),
		(2590, 'CILUAR', 290, 17, 3),
		(2591, 'CIMAHPAR', 290, 17, 3),
		(2592, 'CIPARIGI', 290, 17, 3),
		(2593, 'KEDUNGHALANG', 290, 17, 3),
		(2594, 'TANAH BARU', 290, 17, 3),
		(2595, 'TEGAL GUNDIL', 290, 17, 3),
		(2596, 'BOJONG BARU', 291, 17, 3),
		(2597, 'BOJONG GEDE', 291, 17, 3),
		(2598, 'CIMANGGIS', 291, 17, 3),
		(2599, 'KEDUNG WARINGIN', 291, 17, 3),
		(2600, 'PABUARAN', 291, 17, 3),
		(2601, 'RAGAJAYA', 291, 17, 3),
		(2602, 'RAWA PANJANG', 291, 17, 3),
		(2603, 'SUSUKAN', 291, 17, 3),
		(2604, 'WARINGIN JAYA', 291, 17, 3),
		(2605, 'BABAKAN RADEN', 292, 17, 3),
		(2606, 'BANTAR KUNING', 292, 17, 3),
		(2607, 'CARIU', 292, 17, 3),
		(2608, 'CIBATU TIGA', 292, 17, 3),
		(2609, 'CIKUTAMAHI', 292, 17, 3),
		(2610, 'KARYA MEKAR', 292, 17, 3),
		(2611, 'KUTA MEKAR', 292, 17, 3),
		(2612, 'MEKARWANGI', 292, 17, 3),
		(2613, 'SUKAJADI', 292, 17, 3),
		(2614, 'TEGAL PANJANG', 292, 17, 3),
		(2615, 'BENTENG', 293, 17, 3),
		(2616, 'BOJONG JENGKOL', 293, 17, 3),
		(2617, 'BOJONG RANGKAS', 293, 17, 3),
		(2618, 'CIAMPEA', 293, 17, 3),
		(2619, 'CIAMPEA UDIK', 293, 17, 3),
		(2620, 'CIBADAK', 293, 17, 3),
		(2621, 'CIBANTENG', 293, 17, 3),
		(2622, 'CIBUNTU', 293, 17, 3),
		(2623, 'CICADAS', 293, 17, 3),
		(2624, 'CIHIDEUNG ILIR', 293, 17, 3),
		(2625, 'CIHIDEUNG UDIK', 293, 17, 3),
		(2626, 'CINANGKA', 293, 17, 3),
		(2627, 'TEGAL WARU', 293, 17, 3),
		(2628, 'CEMPLANG', 294, 17, 3),
		(2629, 'CIARUTEUN ILIR', 294, 17, 3),
		(2630, 'CIARUTEUN UDIK', 294, 17, 3),
		(2631, 'CIBATOK 1', 294, 17, 3),
		(2632, 'CIBATOK 2', 294, 17, 3),
		(2633, 'CIJUJUNG', 294, 17, 3),
		(2634, 'CIMANGGU 1', 294, 17, 3),
		(2635, 'CIMANGGU 2', 294, 17, 3),
		(2636, 'DUKUH', 294, 17, 3),
		(2637, 'GALUGA', 294, 17, 3),
		(2638, 'GIRIMULYA', 294, 17, 3),
		(2639, 'LEUWEUNG KOLOT', 294, 17, 3),
		(2640, 'SITU ILIR', 294, 17, 3),
		(2641, 'SITU UDIK', 294, 17, 3),
		(2642, 'SUKAMAJU', 294, 17, 3),
		(2643, 'CIADEG', 295, 17, 3),
		(2644, 'CIBURAYUT', 295, 17, 3),
		(2645, 'CIBURUY', 295, 17, 3),
		(2646, 'CIGOMBONG', 295, 17, 3),
		(2647, 'CISALADA', 295, 17, 3),
		(2648, 'PASIR JAYA', 295, 17, 3),
		(2649, 'SROGOL', 295, 17, 3),
		(2650, 'TUGUJAYA', 295, 17, 3),
		(2651, 'WATESJAYA', 295, 17, 3),
		(2652, 'ARGAPURA', 296, 17, 3),
		(2653, 'BANGUNJAYA', 296, 17, 3),
		(2654, 'BANYU ASIH', 296, 17, 3),
		(2655, 'BANYU RESMI', 296, 17, 3),
		(2656, 'BANYU WANGI', 296, 17, 3),
		(2657, 'BATU JAJAR', 296, 17, 3),
		(2658, 'BUNAR', 296, 17, 3),
		(2659, 'CIGUDEG', 296, 17, 3),
		(2660, 'CINTAMANIK', 296, 17, 3),
		(2661, 'MEKARJAYA', 296, 17, 3),
		(2662, 'RENGASJAJAR', 296, 17, 3),
		(2663, 'SUKAMAJU', 296, 17, 3),
		(2664, 'SUKARAKSA', 296, 17, 3),
		(2665, 'TEGALLEGA', 296, 17, 3),
		(2666, 'WARGAJAYA', 296, 17, 3),
		(2667, 'CIBALUNG', 297, 17, 3),
		(2668, 'CIJERUK', 297, 17, 3),
		(2669, 'CIPELANG', 297, 17, 3),
		(2670, 'CIPICUNG', 297, 17, 3),
		(2671, 'PALASARI', 297, 17, 3),
		(2672, 'SUKAHARJA', 297, 17, 3),
		(2673, 'TAJUR HALANG', 297, 17, 3),
		(2674, 'TANJUNG SARI', 297, 17, 3),
		(2675, 'WARUNG MENTENG', 297, 17, 3),
		(2676, 'CILEUNGSI', 298, 17, 3),
		(2677, 'CILEUNGSI KIDUL', 298, 17, 3),
		(2678, 'CIPENJO', 298, 17, 3),
		(2679, 'CIPEUCANG', 298, 17, 3),
		(2680, 'DAYEUH', 298, 17, 3),
		(2681, 'GANDOANG', 298, 17, 3),
		(2682, 'JATISARI', 298, 17, 3),
		(2683, 'LIMUS NUNGGAL', 298, 17, 3),
		(2684, 'MAMPIR', 298, 17, 3),
		(2685, 'MEKARSARI', 298, 17, 3),
		(2686, 'PASIR ANGIN', 298, 17, 3),
		(2687, 'SETU SARI', 298, 17, 3),
		(2688, 'BABAKAN', 299, 17, 3),
		(2689, 'CIBENTANG', 299, 17, 3),
		(2690, 'CIBEUTEUNG MUARA', 299, 17, 3),
		(2691, 'CIBEUTEUNG UDIK', 299, 17, 3),
		(2692, 'CIHOE (CIHOWE)', 299, 17, 3),
		(2693, 'CISEENG', 299, 17, 3),
		(2694, 'KARIHKIL', 299, 17, 3),
		(2695, 'KURIPAN', 299, 17, 3),
		(2696, 'PARIGI MEKAR', 299, 17, 3),
		(2697, 'PUTAT NUTUG', 299, 17, 3),
		(2698, 'CITEUREUP', 300, 17, 3),
		(2699, 'GUNUNG SARI', 300, 17, 3),
		(2700, 'HAMBALANG', 300, 17, 3),
		(2701, 'KARANG ASEM BARAT', 300, 17, 3),
		(2702, 'KARANG ASEM TIMUR', 300, 17, 3),
		(2703, 'LEUWINUTUG', 300, 17, 3),
		(2704, 'PASIR MUKTI', 300, 17, 3),
		(2705, 'PUSPANEGARA', 300, 17, 3),
		(2706, 'PUSPASARI', 300, 17, 3),
		(2707, 'SANJA', 300, 17, 3),
		(2708, 'SUKAHATI', 300, 17, 3),
		(2709, 'TAJUR', 300, 17, 3),
		(2710, 'TANGKIL', 300, 17, 3),
		(2711, 'TARIKOLOT', 300, 17, 3),
		(2712, 'BABAKAN', 301, 17, 3),
		(2713, 'CIHERANG', 301, 17, 3),
		(2714, 'CIKARAWANG', 301, 17, 3),
		(2715, 'DRAMAGA', 301, 17, 3),
		(2716, 'NEGLASARI', 301, 17, 3),
		(2717, 'PETIR', 301, 17, 3),
		(2718, 'PURWASARI', 301, 17, 3),
		(2719, 'SINAR SARI', 301, 17, 3),
		(2720, 'SUKADAMAI', 301, 17, 3),
		(2721, 'SUKAWENING', 301, 17, 3),
		(2722, 'BOJONG KULUR', 302, 17, 3),
		(2723, 'BOJONG NANGKA', 302, 17, 3),
		(2724, 'CIANGSANA', 302, 17, 3),
		(2725, 'CICADAS', 302, 17, 3),
		(2726, 'CIKEAS UDIK', 302, 17, 3),
		(2727, 'GUNUNG PUTRI', 302, 17, 3),
		(2728, 'KARANGGAN', 302, 17, 3),
		(2729, 'NAGRAG', 302, 17, 3),
		(2730, 'TLAJUNG UDIK', 302, 17, 3),
		(2731, 'WANAHERANG', 302, 17, 3),
		(2732, 'CIBADUNG', 303, 17, 3),
		(2733, 'CIBINONG', 303, 17, 3),
		(2734, 'CIDOKOM', 303, 17, 3),
		(2735, 'CURUG', 303, 17, 3),
		(2736, 'GUNUNG SINDUR', 303, 17, 3),
		(2737, 'JAMPANG', 303, 17, 3),
		(2738, 'PABUARAN', 303, 17, 3),
		(2739, 'PADURENAN', 303, 17, 3),
		(2740, 'PENGASINAN', 303, 17, 3),
		(2741, 'RAWAKALONG', 303, 17, 3),
		(2742, 'BAGOANG', 304, 17, 3),
		(2743, 'BARENGKOK', 304, 17, 3),
		(2744, 'CIKOPOMAYAK', 304, 17, 3),
		(2745, 'CURUG', 304, 17, 3),
		(2746, 'JASINGA', 304, 17, 3),
		(2747, 'JUGALA JAYA', 304, 17, 3),
		(2748, 'KALONGSAWAH', 304, 17, 3),
		(2749, 'KOLEANG', 304, 17, 3),
		(2750, 'NEGLASARI', 304, 17, 3),
		(2751, 'PAMAGERSARI', 304, 17, 3),
		(2752, 'PANGAUR', 304, 17, 3),
		(2753, 'PANGRADIN', 304, 17, 3),
		(2754, 'SETU', 304, 17, 3),
		(2755, 'SIPAK', 304, 17, 3),
		(2756, 'TEGAL WANGI', 304, 17, 3),
		(2757, 'BALEKAMBANG', 305, 17, 3),
		(2758, 'BENDUNGAN', 305, 17, 3),
		(2759, 'CIBODAS', 305, 17, 3),
		(2760, 'JONGGOL', 305, 17, 3),
		(2761, 'SINGAJAYA', 305, 17, 3),
		(2762, 'SINGASARI', 305, 17, 3),
		(2763, 'SIRNAGALIH', 305, 17, 3),
		(2764, 'SUKAGALIH', 305, 17, 3),
		(2765, 'SUKAJAYA', 305, 17, 3),
		(2766, 'SUKAMAJU', 305, 17, 3),
		(2767, 'SUKAMANAH', 305, 17, 3),
		(2768, 'SUKANEGARA', 305, 17, 3),
		(2769, 'SUKASIRNA', 305, 17, 3),
		(2770, 'WENINGGALIH', 305, 17, 3),
		(2771, 'ATANG SENJAYA', 306, 17, 3),
		(2772, 'BOJONG', 306, 17, 3),
		(2773, 'JAMPANG', 306, 17, 3),
		(2774, 'KEMANG', 306, 17, 3),
		(2775, 'PABUARAN', 306, 17, 3),
		(2776, 'PARAKAN JAYA', 306, 17, 3),
		(2777, 'PONDOK UDIK', 306, 17, 3),
		(2778, 'SEMPLAK BARAT', 306, 17, 3),
		(2779, 'TEGAL', 306, 17, 3),
		(2780, 'BANTAR JATI', 307, 17, 3),
		(2781, 'BOJONG', 307, 17, 3),
		(2782, 'CIKAHURIPAN', 307, 17, 3),
		(2783, 'KEMBANG KUNING', 307, 17, 3),
		(2784, 'KLAPANUNGGAL', 307, 17, 3),
		(2785, 'LEUWIKARET', 307, 17, 3),
		(2786, 'LIGARMUKTI', 307, 17, 3),
		(2787, 'LULUT', 307, 17, 3),
		(2788, 'NAMBO', 307, 17, 3),
		(2789, 'BARENGKOK', 308, 17, 3),
		(2790, 'CIBEBER 1', 308, 17, 3),
		(2791, 'CIBEBER 2', 308, 17, 3),
		(2792, 'KARACAK', 308, 17, 3),
		(2793, 'KAREHKEL', 308, 17, 3),
		(2794, 'KARYASARI', 308, 17, 3),
		(2795, 'LEUWILIANG', 308, 17, 3),
		(2796, 'LEUWIMEKAR', 308, 17, 3),
		(2797, 'PABANGBON', 308, 17, 3),
		(2798, 'PURASARI', 308, 17, 3),
		(2799, 'PURASEDA', 308, 17, 3),
		(2800, 'BABAKAN SADENG', 309, 17, 3),
		(2801, 'KALONG 1', 309, 17, 3),
		(2802, 'KALONG 2', 309, 17, 3),
		(2803, 'LEUWISADENG', 309, 17, 3),
		(2804, 'SADENG', 309, 17, 3),
		(2805, 'SADENGKOLOT', 309, 17, 3),
		(2806, 'SIBANTENG', 309, 17, 3),
		(2807, 'WANGUN JAYA', 309, 17, 3),
		(2808, 'CIPAYUNG (CIPAYUNG DATAR)', 310, 17, 3),
		(2809, 'CIPAYUNG GIRANG', 310, 17, 3),
		(2810, 'GADOG', 310, 17, 3),
		(2811, 'KUTA', 310, 17, 3),
		(2812, 'MEGAMENDUNG', 310, 17, 3),
		(2813, 'PASIR ANGIN', 310, 17, 3),
		(2814, 'SUKAGALIH', 310, 17, 3),
		(2815, 'SUKAKARYA', 310, 17, 3),
		(2816, 'SUKAMAHI', 310, 17, 3),
		(2817, 'SUKAMAJU', 310, 17, 3),
		(2818, 'SUKAMANAH', 310, 17, 3),
		(2819, 'SUKARESMI', 310, 17, 3),
		(2820, 'BANTAR KARET', 311, 17, 3),
		(2821, 'CISARUA', 311, 17, 3),
		(2822, 'CURUG BITUNG', 311, 17, 3),
		(2823, 'HAMBARO', 311, 17, 3),
		(2824, 'KALONG LIUD', 311, 17, 3),
		(2825, 'MALASARI', 311, 17, 3),
		(2826, 'NANGGUNG', 311, 17, 3),
		(2827, 'PANGKAL JAYA', 311, 17, 3),
		(2828, 'PARAKAN MUNCANG', 311, 17, 3),
		(2829, 'SUKALUYU', 311, 17, 3),
		(2830, 'CIASIHAN', 312, 17, 3),
		(2831, 'CIASMARA', 312, 17, 3),
		(2832, 'CIBENING', 312, 17, 3),
		(2833, 'CIBITUNG KULON', 312, 17, 3),
		(2834, 'CIBITUNG WETAN', 312, 17, 3),
		(2835, 'CIBUNIAN', 312, 17, 3),
		(2836, 'CIMAYANG', 312, 17, 3),
		(2837, 'GUNUNG BUNDER 1', 312, 17, 3),
		(2838, 'GUNUNG BUNDER 2', 312, 17, 3),
		(2839, 'GUNUNG MENYAN', 312, 17, 3),
		(2840, 'GUNUNG PICUNG', 312, 17, 3),
		(2841, 'GUNUNG SARI', 312, 17, 3),
		(2842, 'PAMIJAHAN', 312, 17, 3),
		(2843, 'PASAREAN', 312, 17, 3),
		(2844, 'PURWABAKTI', 312, 17, 3),
		(2845, 'BOJONG INDAH', 313, 17, 3),
		(2846, 'BOJONG SEMPU', 313, 17, 3),
		(2847, 'COGREG', 313, 17, 3),
		(2848, 'IWUL', 313, 17, 3),
		(2849, 'JABON MEKAR', 313, 17, 3),
		(2850, 'PAMAGERSARI', 313, 17, 3),
		(2851, 'PARUNG', 313, 17, 3),
		(2852, 'WARU', 313, 17, 3),
		(2853, 'WARUJAYA', 313, 17, 3),
		(2854, 'CIBUNAR', 314, 17, 3),
		(2855, 'CIKUDA', 314, 17, 3),
		(2856, 'DAGO', 314, 17, 3),
		(2857, 'GINTUNG CILEJET', 314, 17, 3),
		(2858, 'GOROWONG', 314, 17, 3),
		(2859, 'JAGABAYA', 314, 17, 3),
		(2860, 'JAGABITA', 314, 17, 3),
		(2861, 'KABASIRAN', 314, 17, 3),
		(2862, 'LUMPANG', 314, 17, 3),
		(2863, 'PARUNGPANJANG', 314, 17, 3),
		(2864, 'PINGKU', 314, 17, 3),
		(2865, 'BANTARJAYA', 315, 17, 3),
		(2866, 'BANTARSARI', 315, 17, 3),
		(2867, 'CANDALI', 315, 17, 3),
		(2868, 'MEKARSARI', 315, 17, 3),
		(2869, 'PASIRGAOK', 315, 17, 3),
		(2870, 'RANCABUNGUR', 315, 17, 3),
		(2871, 'CIBODAS', 316, 17, 3),
		(2872, 'CIDOKOM', 316, 17, 3),
		(2873, 'CIPINANG', 316, 17, 3),
		(2874, 'GOBANG', 316, 17, 3),
		(2875, 'KAMPUNG SAWAH', 316, 17, 3),
		(2876, 'KERTA JAYA', 316, 17, 3),
		(2877, 'LEUWIBATU', 316, 17, 3),
		(2878, 'MEKAR SARI', 316, 17, 3),
		(2879, 'MEKARJAYA', 316, 17, 3),
		(2880, 'RABAK', 316, 17, 3),
		(2881, 'RUMPIN', 316, 17, 3),
		(2882, 'SUKAMULYA', 316, 17, 3),
		(2883, 'SUKASARI', 316, 17, 3),
		(2884, 'TAMAN SARI', 316, 17, 3),
		(2885, 'CIBADAK', 317, 17, 3),
		(2886, 'PABUARAN', 317, 17, 3),
		(2887, 'SIRNAJAYA', 317, 17, 3),
		(2888, 'SUKADAMAI', 317, 17, 3),
		(2889, 'SUKAHARJA', 317, 17, 3),
		(2890, 'SUKAMAKMUR', 317, 17, 3),
		(2891, 'SUKAMULYA', 317, 17, 3),
		(2892, 'SUKARESMI', 317, 17, 3),
		(2893, 'SUKAWANGI', 317, 17, 3),
		(2894, 'WARGAJAYA', 317, 17, 3),
		(2895, 'CITAYAM', 318, 17, 3),
		(2896, 'KALISUREN', 318, 17, 3),
		(2897, 'NANGGERANG', 318, 17, 3),
		(2898, 'SASAK PANJANG', 318, 17, 3),
		(2899, 'SUKMAJAYA', 318, 17, 3),
		(2900, 'TAJUR HALANG', 318, 17, 3),
		(2901, 'TONJONG', 318, 17, 3),
		(2902, 'CIBADAK', 319, 17, 3),
		(2903, 'KAYUMANIS', 319, 17, 3),
		(2904, 'KEBON PEDES', 319, 17, 3),
		(2905, 'KEDUNG BADAK', 319, 17, 3),
		(2906, 'KEDUNG JAYA', 319, 17, 3),
		(2907, 'KEDUNG WARINGIN', 319, 17, 3),
		(2908, 'KENCANA', 319, 17, 3),
		(2909, 'MEKARWANGI', 319, 17, 3),
		(2910, 'SUKADAMAI', 319, 17, 3),
		(2911, 'SUKARESMI', 319, 17, 3),
		(2912, 'TANAH SAREAL', 319, 17, 3),
		(2913, 'BABAKAN', 320, 17, 3),
		(2914, 'BATOK', 320, 17, 3),
		(2915, 'BOJONG', 320, 17, 3),
		(2916, 'CILAKU', 320, 17, 3),
		(2917, 'CIOMAS', 320, 17, 3),
		(2918, 'SINGABANGSA', 320, 17, 3),
		(2919, 'SINGABRAJA', 320, 17, 3),
		(2920, 'TAPOS', 320, 17, 3),
		(2921, 'TENJO', 320, 17, 3),
		(2922, 'CIBITUNG TENGAH', 321, 17, 3),
		(2923, 'CINANGNENG', 321, 17, 3),
		(2924, 'GUNUNG MALANG', 321, 17, 3),
		(2925, 'SITU DAUN', 321, 17, 3),
		(2926, 'TAPOS 1', 321, 17, 3),
		(2927, 'TAPOS 2', 321, 17, 3),
		(2928, 'BAREGBEG', 322, 18, 3),
		(2929, 'JELAT', 322, 18, 3),
		(2930, 'KARANGAMPEL', 322, 18, 3),
		(2931, 'MEKARJAYA', 322, 18, 3),
		(2932, 'PETIRHILIR', 322, 18, 3),
		(2933, 'PUSAKANAGARA', 322, 18, 3),
		(2934, 'SAGULING', 322, 18, 3),
		(2935, 'SUKAMAJU', 322, 18, 3),
		(2936, 'SUKAMULYA', 322, 18, 3),
		(2937, 'BENTENG', 323, 18, 3),
		(2938, 'CIAMIS', 323, 18, 3),
		(2939, 'CIGEMBOR', 323, 18, 3),
		(2940, 'CISADAP', 323, 18, 3),
		(2941, 'IMBANAGARA', 323, 18, 3),
		(2942, 'IMBANAGARA RAYA', 323, 18, 3),
		(2943, 'KERTASARI', 323, 18, 3),
		(2944, 'LINGGASARI', 323, 18, 3),
		(2945, 'MALEBER', 323, 18, 3),
		(2946, 'PANYINGKIRAN', 323, 18, 3),
		(2947, 'PAWINDAN', 323, 18, 3),
		(2948, 'SINDANGRASA', 323, 18, 3),
		(2949, 'CIDOLOG', 324, 18, 3),
		(2950, 'CIDOLOG', 324, 18, 3),
		(2951, 'CIKARANG', 324, 18, 3),
		(2952, 'CIPAMINGKIS', 324, 18, 3),
		(2953, 'CIPARAY', 324, 18, 3),
		(2954, 'HEGARMANAH', 324, 18, 3),
		(2955, 'JANGGALA', 324, 18, 3),
		(2956, 'JELEGONG', 324, 18, 3),
		(2957, 'MEKARJAYA', 324, 18, 3),
		(2958, 'SUKASARI', 324, 18, 3),
		(2959, 'TEGALLEGA', 324, 18, 3),
		(2960, 'CIHAURBEUTI', 325, 18, 3),
		(2961, 'CIJULANG', 325, 18, 3),
		(2962, 'PADAMULYA', 325, 18, 3),
		(2963, 'PAMOKOLAN', 325, 18, 3),
		(2964, 'PASIRTAMIANG', 325, 18, 3),
		(2965, 'SUKAHAJI', 325, 18, 3),
		(2966, 'SUKAHURIP', 325, 18, 3),
		(2967, 'SUKAMAJU', 325, 18, 3),
		(2968, 'SUKAMULYA', 325, 18, 3),
		(2969, 'SUKASETIA', 325, 18, 3),
		(2970, 'SUMBERJAYA', 325, 18, 3),
		(2971, 'BOJONGMENGGER', 326, 18, 3),
		(2972, 'CIHARALANG', 326, 18, 3),
		(2973, 'CIJEUNGJING', 326, 18, 3),
		(2974, 'DEWASARI', 326, 18, 3),
		(2975, 'HANDAPHERANG', 326, 18, 3),
		(2976, 'KARANGANYAR', 326, 18, 3),
		(2977, 'KARANGKAMULYAN', 326, 18, 3),
		(2978, 'KERTABUMI', 326, 18, 3),
		(2979, 'KERTAHARJA', 326, 18, 3),
		(2980, 'PAMALAYAN', 326, 18, 3),
		(2981, 'UTAMA', 326, 18, 3),
		(2982, 'CIKONENG', 327, 18, 3),
		(2983, 'CIMARI', 327, 18, 3),
		(2984, 'DARMACAANG', 327, 18, 3),
		(2985, 'GEGEMPALAN', 327, 18, 3),
		(2986, 'KUJANG', 327, 18, 3),
		(2987, 'MARGALUYU', 327, 18, 3),
		(2988, 'NASOL', 327, 18, 3),
		(2989, 'PANARAGAN', 327, 18, 3),
		(2990, 'SINDANGSARI', 327, 18, 3),
		(2991, 'BEBER', 328, 18, 3),
		(2992, 'BOJONGMALANG', 328, 18, 3),
		(2993, 'CIMARAGAS', 328, 18, 3),
		(2994, 'JAYARAKSA', 328, 18, 3),
		(2995, 'RAKSABAYA', 328, 18, 3),
		(2996, 'BANGBAYANG', 329, 18, 3),
		(2997, 'BUNISEURI', 329, 18, 3),
		(2998, 'CIAKAR', 329, 18, 3),
		(2999, 'CIEURIH', 329, 18, 3),
		(3000, 'CIPAKU', 329, 18, 3)
	`)
}
