package main

import (
	"fmt"
	"strconv"
	"strings"
)

const day4Input = `
5-90,4-90
52-52,3-51
46-81,45-80
15-48,49-75
14-81,14-81
23-44,23-37
60-60,48-61
10-58,58-75
12-86,4-77
8-80,7-89
82-92,91-95
67-67,35-67
92-92,65-91
21-96,20-95
29-30,30-81
15-97,6-97
52-96,51-98
16-74,16-74
7-14,15-58
10-10,24-25
63-93,58-93
1-99,2-4
17-17,17-96
46-69,45-88
10-90,17-66
41-99,41-45
9-55,8-56
14-54,54-54
2-16,1-99
10-99,9-97
75-97,96-97
9-79,19-80
1-1,3-41
5-54,4-53
90-96,59-70
21-88,89-97
4-99,3-98
98-99,1-98
3-48,7-47
5-90,6-89
45-94,1-93
11-51,12-36
5-95,1-5
83-83,56-84
5-82,4-4
6-79,7-92
70-92,70-92
32-59,32-32
55-80,54-66
6-53,55-91
1-76,1-77
55-80,30-54
36-38,18-37
9-38,39-39
12-23,24-88
28-95,23-25
36-43,35-65
46-53,17-47
6-9,10-87
62-62,57-62
74-79,75-81
14-15,14-85
15-65,9-66
18-99,17-20
36-93,83-94
15-41,14-18
74-91,73-91
61-96,69-96
74-78,73-76
1-89,1-90
78-91,90-90
2-45,11-45
51-82,76-78
44-89,45-85
66-92,67-93
3-88,1-2
23-86,22-68
1-39,2-39
4-62,63-63
43-71,44-82
48-99,48-97
98-99,35-89
83-87,43-83
9-25,9-9
6-8,7-9
89-97,32-78
18-59,19-75
49-59,49-53
2-2,1-99
88-93,59-92
1-2,1-98
53-54,11-54
19-58,6-57
66-85,65-65
48-77,16-39
22-89,23-99
1-69,37-69
10-73,74-74
16-22,16-23
31-79,30-74
73-79,5-74
10-68,11-95
45-71,23-45
12-12,13-69
79-88,58-87
88-97,87-96
25-28,29-93
11-88,87-89
51-80,52-79
37-38,11-38
85-88,28-87
29-76,28-98
90-93,9-89
9-81,8-81
21-94,66-94
55-56,49-56
25-67,16-25
23-48,49-49
16-54,24-54
33-75,33-33
44-88,71-89
98-98,3-98
32-32,33-34
24-98,24-25
7-19,5-8
52-54,52-53
53-58,54-96
24-97,25-98
5-44,45-47
9-13,14-42
39-39,39-74
61-97,60-90
82-83,83-98
15-44,1-16
65-67,64-83
67-87,66-68
16-71,71-80
41-56,56-57
44-65,44-67
66-81,19-65
6-93,92-93
7-91,5-92
29-82,29-92
70-71,71-77
20-90,26-89
14-56,15-81
19-54,18-20
5-25,5-6
10-54,9-95
11-94,8-10
4-91,4-5
14-92,92-92
8-88,89-89
22-77,22-38
5-95,6-97
42-89,43-90
7-39,6-38
45-55,6-46
11-68,12-68
37-54,42-53
21-42,21-35
80-97,34-79
23-94,93-95
74-89,53-74
24-85,24-84
2-97,2-2
10-12,13-31
29-38,30-39
96-96,86-97
31-94,30-50
94-99,16-95
4-92,91-93
6-6,7-93
2-93,1-93
28-56,29-55
23-40,23-58
68-68,24-69
79-93,40-83
30-97,29-98
43-56,42-57
5-91,5-87
80-80,54-79
16-63,15-63
1-1,2-95
73-77,72-75
4-99,3-95
7-98,7-97
29-72,72-99
83-88,56-62
42-42,41-61
10-97,6-98
9-96,10-97
8-93,8-92
57-59,36-58
47-96,61-88
13-97,23-75
16-17,34-53
2-99,2-98
50-89,35-68
7-19,6-62
64-64,62-69
40-40,3-40
23-73,22-74
1-26,25-90
96-99,36-94
14-21,15-77
57-72,53-57
11-12,14-32
36-39,27-38
22-22,21-96
43-85,43-84
7-75,12-75
76-76,77-77
1-1,2-99
8-19,7-21
4-14,8-15
15-89,13-89
26-26,25-51
48-98,97-99
3-93,4-47
17-88,99-99
98-99,41-97
30-87,41-87
25-41,29-42
14-16,16-66
38-85,86-88
2-2,3-71
4-23,22-94
27-60,32-46
14-15,14-81
42-64,47-65
32-32,31-92
1-98,2-97
4-5,12-59
7-76,6-76
94-95,35-95
4-94,94-95
25-72,24-35
56-77,52-56
49-82,42-48
47-76,1-61
41-50,50-50
24-70,18-69
53-66,25-91
97-99,46-92
6-97,14-88
17-88,1-89
43-43,42-89
14-85,84-90
24-69,69-70
47-75,15-76
23-76,76-77
81-97,81-81
22-64,23-82
33-34,12-34
7-7,8-17
63-63,58-63
66-94,93-95
54-73,21-72
10-41,11-95
60-60,16-59
13-95,94-94
1-33,2-96
28-29,28-30
2-66,3-3
25-28,25-52
65-79,75-79
12-28,27-91
4-96,95-96
35-72,34-71
16-18,17-79
3-97,99-99
13-22,13-14
8-55,4-55
30-96,61-67
53-53,54-96
10-11,11-86
81-97,80-97
6-8,9-97
41-54,11-54
15-71,15-15
69-97,90-96
1-99,2-2
3-97,2-98
14-82,21-82
1-95,1-94
11-11,10-72
93-99,31-94
8-11,5-10
4-33,3-34
9-89,94-94
2-19,18-43
95-96,70-72
61-84,60-85
95-97,17-96
33-70,32-71
36-36,37-50
6-85,6-6
22-49,22-75
5-93,5-6
86-86,8-86
91-94,90-92
26-56,56-56
13-30,12-95
11-55,12-43
21-78,20-79
69-94,69-93
85-87,10-45
33-84,80-82
16-78,16-92
25-64,39-65
33-33,34-78
9-57,9-9
94-94,66-95
10-48,7-48
57-88,63-87
4-56,2-2
23-94,22-64
45-66,46-67
44-94,43-95
29-66,66-66
30-75,74-74
93-97,41-94
3-38,2-12
2-48,49-49
20-94,19-98
22-55,7-55
1-58,2-88
76-81,77-80
2-37,37-68
22-44,23-61
21-91,90-91
98-98,2-97
7-93,8-95
23-98,12-24
56-56,56-96
52-77,37-78
91-93,3-92
57-96,56-95
96-99,1-97
44-57,45-85
8-84,5-6
2-95,94-98
91-95,91-99
5-94,4-4
26-87,86-86
42-80,79-81
21-33,32-40
29-69,29-69
79-79,79-97
23-91,28-95
11-36,1-35
4-29,5-91
7-14,16-95
46-85,26-89
29-95,11-18
4-94,2-95
3-18,18-18
16-35,3-15
3-39,2-39
22-34,23-33
26-99,98-99
10-85,84-85
42-42,40-42
1-1,1-35
6-97,7-99
4-23,3-8
62-63,63-82
1-80,2-81
1-19,1-27
13-91,14-91
33-78,34-79
7-95,8-8
7-82,75-83
26-94,11-27
21-23,21-96
19-24,20-66
26-99,25-80
38-79,39-80
39-95,94-95
25-26,9-26
92-98,74-93
43-48,45-49
36-81,16-80
2-99,3-6
73-74,73-77
2-88,14-87
28-58,27-57
27-42,28-34
2-3,3-81
64-78,77-79
92-94,26-93
4-45,25-45
13-24,5-25
16-31,15-33
27-37,34-38
13-88,12-83
19-91,90-91
44-76,45-83
13-15,2-2
57-82,18-89
26-98,25-98
5-73,3-87
10-12,11-13
97-99,13-98
30-47,30-30
3-62,48-58
5-27,6-11
4-57,56-57
92-98,1-90
75-97,80-97
24-90,90-91
10-81,11-55
46-57,46-53
15-62,2-16
3-82,1-82
13-36,2-14
8-88,8-9
8-86,87-94
84-87,5-83
6-96,6-99
62-88,89-91
57-71,56-70
6-41,8-40
47-87,87-96
15-98,16-99
7-95,6-80
36-48,36-58
13-90,89-91
39-53,38-47
29-52,30-41
18-87,87-87
39-92,60-91
30-98,28-99
8-8,9-77
10-54,10-54
45-46,32-45
5-96,5-99
98-98,3-98
1-48,6-47
49-61,60-62
46-47,5-47
20-49,6-89
29-81,28-81
7-72,71-71
39-98,43-86
23-27,27-27
6-87,7-7
1-87,6-69
18-95,19-30
38-88,38-65
19-53,52-53
42-67,41-67
7-23,8-24
13-14,13-35
62-64,23-54
26-48,4-34
64-96,18-29
56-56,33-56
10-81,9-10
45-53,24-96
11-70,29-69
54-69,91-97
40-72,39-73
4-65,64-81
17-19,20-60
39-73,39-74
62-62,62-87
58-60,25-59
5-82,1-6
83-84,28-83
51-99,51-81
16-54,17-89
65-66,64-68
8-64,9-64
2-5,5-48
27-97,6-98
24-88,24-86
12-43,13-13
52-80,79-79
15-81,15-82
1-93,1-1
36-77,78-78
3-63,23-62
5-71,5-98
1-94,7-93
5-75,5-79
22-26,21-28
56-56,24-55
43-43,42-67
5-48,47-57
13-93,14-94
12-60,13-59
13-27,12-87
3-68,69-79
18-96,7-13
9-29,9-10
45-49,49-49
6-86,4-7
25-90,26-94
33-39,40-42
14-84,14-79
8-60,12-76
1-18,17-30
28-42,41-42
3-92,87-92
15-46,46-46
94-98,34-91
2-78,12-77
2-33,31-31
71-95,72-72
80-80,79-98
38-83,37-82
75-76,64-76
53-56,57-83
34-57,35-58
38-78,38-39
7-38,6-38
85-87,60-86
16-43,15-99
21-94,22-22
1-89,90-94
1-2,1-92
83-84,6-47
26-26,25-36
14-18,14-72
3-10,14-66
32-82,32-86
1-62,2-2
15-84,16-83
62-63,49-63
56-79,9-56
12-93,39-92
81-86,80-88
97-97,61-97
10-99,8-10
14-24,12-15
5-8,8-80
39-97,38-60
1-94,1-95
32-54,31-53
73-88,61-72
79-90,20-78
20-99,21-98
60-73,61-74
7-13,8-31
7-79,7-35
85-86,2-85
27-35,9-34
2-96,1-96
45-61,46-51
2-2,2-8
87-87,59-86
4-84,5-83
35-95,36-97
35-42,43-64
4-8,7-64
73-74,17-74
8-79,5-59
53-95,18-94
11-87,2-86
34-80,35-80
22-48,23-48
36-78,79-84
20-42,28-42
61-87,59-86
19-98,20-51
3-3,4-37
43-94,45-81
39-40,33-45
92-93,52-93
1-73,72-74
40-56,40-57
63-86,62-85
34-79,33-78
58-60,23-59
3-30,29-30
61-61,60-61
67-87,2-86
11-83,68-82
13-65,11-65
30-70,34-69
26-87,14-14
3-90,89-91
14-42,15-15
12-32,7-9
1-99,99-99
20-95,19-20
41-78,41-75
30-40,15-39
27-43,27-36
3-95,3-97
51-52,14-52
72-72,72-94
40-40,40-83
5-94,4-4
43-62,18-42
18-92,18-19
13-83,12-12
26-36,15-35
10-67,25-68
11-70,11-68
65-96,65-66
31-96,19-32
25-27,17-26
3-62,2-62
5-11,5-33
71-71,53-72
6-85,6-86
47-48,43-47
32-94,33-33
5-89,5-6
86-98,85-94
6-6,6-97
47-48,48-92
17-56,56-56
4-80,3-81
94-98,25-93
23-48,5-22
8-79,79-79
41-75,11-74
3-86,4-66
43-82,42-43
5-88,1-94
44-44,44-95
88-98,8-21
7-95,94-95
74-97,73-96
72-72,45-73
9-10,9-96
64-90,64-64
4-53,53-53
8-83,9-83
1-84,1-84
6-81,5-81
16-56,10-57
32-35,31-39
22-86,82-87
29-94,3-93
56-88,57-89
11-23,23-23
48-69,68-81
55-70,56-78
3-16,20-94
19-48,11-48
13-73,1-72
19-26,19-93
25-60,60-80
14-80,80-81
71-90,71-91
17-67,15-67
96-96,11-97
45-55,26-56
28-54,29-29
83-96,83-99
40-94,40-98
1-97,2-97
28-28,27-96
86-86,21-86
3-72,4-71
18-95,94-95
19-19,19-47
53-89,52-88
42-42,42-88
97-98,18-92
24-94,11-24
4-20,15-21
44-44,45-70
26-37,36-82
23-79,28-93
33-91,34-98
2-99,2-2
10-48,9-49
12-12,13-13
16-16,10-16
3-79,12-80
19-19,24-99
50-90,2-89
91-92,55-92
12-94,13-93
8-89,39-76
7-8,7-80
5-36,35-97
27-67,26-66
41-92,44-76
50-65,51-66
61-82,60-60
54-96,55-96
57-64,51-63
21-35,22-35
25-71,6-71
49-94,13-48
31-67,31-32
78-87,79-85
91-91,67-91
24-90,13-89
9-18,8-85
9-18,9-22
13-81,80-81
11-29,12-30
19-32,33-66
14-97,14-99
55-69,68-88
38-59,39-71
49-63,55-64
18-95,18-67
64-74,73-80
21-44,21-45
77-81,76-81
56-57,11-57
30-99,31-98
21-25,44-87
10-66,11-65
7-61,6-15
5-95,5-5
1-93,1-2
64-82,73-81
39-88,39-89
8-67,9-9
59-99,58-58
11-74,97-98
61-83,7-59
23-96,24-96
9-80,8-87
18-80,18-81
1-90,1-89
15-68,15-90
11-96,42-96
41-61,42-62
56-65,33-84
34-56,4-21
7-29,2-28
51-91,90-91
5-5,6-88
49-95,49-49
25-89,25-89
60-77,85-88
16-24,11-24
68-81,50-69
9-60,9-89
90-90,5-90
14-82,14-92
1-99,17-98
34-69,35-68
4-38,10-34
9-64,57-81
24-96,95-95
5-98,21-93
11-11,10-53
93-93,25-93
2-2,3-96
3-3,3-98
32-47,33-33
51-68,25-69
89-97,59-96
15-75,76-76
15-92,9-93
15-27,5-16
89-89,6-98
14-92,90-91
10-27,27-29
26-82,26-81
58-86,58-58
8-80,3-8
55-56,4-59
44-96,44-97
1-89,1-1
22-35,23-34
6-86,3-6
2-44,1-14
10-80,9-64
17-79,18-85
98-98,29-98
39-99,38-39
73-75,68-74
11-34,11-12
19-39,18-34
31-78,31-90
28-69,29-69
32-99,22-33
11-78,79-94
17-51,52-52
1-94,8-44
2-95,2-95
9-92,8-92
83-94,1-84
43-45,1-44
19-56,18-32
64-88,21-88
10-63,9-63
13-57,57-57
4-36,33-36
10-82,82-90
43-56,43-55
37-99,98-99
14-94,15-92
24-89,89-89
14-49,1-72
15-81,81-88
18-26,25-95
91-97,48-92
46-79,92-94
55-60,59-96
5-98,4-99
22-94,23-94
60-64,49-63
69-94,43-70
58-97,34-59
63-92,42-64
16-93,11-14
5-6,6-18
19-78,79-94
96-98,4-74
11-92,12-92
99-99,6-98
95-95,5-95
16-52,16-52
90-90,61-91
24-39,23-77
27-64,26-82
1-2,1-99
31-62,32-32
45-73,60-74
11-12,11-89
2-91,90-96
21-47,22-50
95-95,19-67
10-93,93-99
5-79,4-76
9-70,9-98
58-87,59-87
4-78,5-78
33-50,32-40
47-98,46-75
40-92,41-66
56-85,55-56
81-81,36-82
11-44,11-12
3-3,2-86
83-83,9-82
21-96,21-35
80-85,73-85
79-79,78-94
97-97,7-96
27-92,21-24
58-77,28-57
14-87,19-86
15-96,45-96
5-78,5-90
95-95,17-96
1-35,1-1
12-93,11-94
89-90,23-90
54-98,55-98
13-81,14-56
8-67,9-57
2-59,5-48
46-54,47-92
8-35,8-8
98-99,17-97
27-90,27-91
10-34,33-33
98-99,10-99
40-73,40-74
13-14,13-47
15-22,16-21
97-98,46-90
5-46,25-46
25-82,15-81
36-57,53-53
42-92,43-92
59-74,75-94
27-87,27-67
85-86,2-86
37-53,36-54
99-99,12-96
91-98,90-97
19-43,85-96
99-99,7-99
71-73,16-72
50-99,49-49
24-95,96-97
43-43,42-97
33-96,34-98
4-5,4-86
15-87,4-86
41-96,41-42
89-94,11-32
83-91,71-84
2-85,3-85
53-53,53-87
2-69,68-77
90-91,17-91
49-50,69-92
15-18,15-97
38-79,47-78
4-60,5-5
20-92,47-91
99-99,1-94
70-71,10-71
45-49,12-48
33-35,36-54
2-85,99-99
76-91,25-90
69-72,71-80
6-17,5-92
3-88,2-87
31-39,31-32
33-93,73-94
37-97,58-97
16-16,15-15
7-98,9-97
35-53,53-53
83-83,57-89
82-83,14-83
39-45,82-87
59-66,58-65
21-96,33-97
65-66,15-72
6-93,6-7
14-92,15-15
5-35,34-91
19-91,18-92
40-42,41-73
21-70,21-21
55-95,54-94
36-77,77-77
81-97,71-80
49-67,43-48
25-93,24-94
53-91,18-90
28-52,35-51
75-75,65-76
73-73,51-72
3-3,3-57
97-98,6-98
19-73,74-74
65-86,64-64
43-90,44-91
3-98,2-97
5-84,98-99
55-98,54-77
22-29,21-22
10-72,5-73
52-61,53-95
4-80,80-80
79-81,77-80
19-97,2-97
76-95,96-97
11-84,11-96
69-79,70-89
7-20,3-8
10-96,9-96
5-28,5-29
21-28,21-88
40-46,39-47
26-38,26-37
30-97,97-97
36-59,35-60
1-2,1-26
45-65,4-46
92-95,58-93
58-70,66-70
6-8,7-88
29-72,29-73
77-92,77-91
86-92,60-93
55-75,55-56
9-58,10-10
4-36,4-37
3-94,5-95
4-98,97-99
11-86,12-87
94-94,16-63
4-97,96-98
22-74,21-44
4-6,4-5
5-54,5-5
21-50,21-86
27-84,28-85
30-95,30-30
49-49,50-50
3-55,2-56
14-30,14-14
41-77,40-81
35-95,30-96
51-91,52-83
52-52,16-52
55-56,4-63
16-63,15-64
3-3,4-52
10-66,11-11
11-64,12-71
5-91,4-92
47-60,60-60
32-95,94-98
78-79,78-79
5-72,6-72
64-81,47-82
3-64,7-64
`

func day4() {

	trimmed := strings.TrimSuffix(day4Input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")

	var score int

	for _, line := range strings.Split(strings.TrimSuffix(trimmed, "\n"), "\n") {

		pairs := strings.Split(line, ",")
		first := pairs[0]
		firstArray := strings.Split(first, "-")
		firstExploded := explode(firstArray[0], firstArray[1])
		fmt.Println(fmt.Sprintf("first: %v", firstExploded))

		second := pairs[1]
		secondArray := strings.Split(second, "-")
		secondExploded := explode(secondArray[0], secondArray[1])
		fmt.Println(fmt.Sprintf("second: %v", secondExploded))

		intersection := intersectTwo(firstExploded, secondExploded)
		fmt.Println(fmt.Sprintf("intersection: %v", intersection))

		if len(intersection) > 0 {
			score += 1
		}

		//		if slices.Compare(intersection, firstExploded) == 0 || slices.Compare(intersection, secondExploded) == 0 {
		//			score += 1
		//		}
	}

	fmt.Println(score)

}

func explode(start, end string) []int {
	var s, e int

	s, _ = strconv.Atoi(start)
	e, _ = strconv.Atoi(end)

	var result []int
	for i := s; i <= e; i++ {
		result = append(result, i)
	}
	return result
}

func intersectTwo[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}
