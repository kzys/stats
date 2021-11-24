# stats

[![codecov](https://codecov.io/gh/kzys/stats/branch/main/graph/badge.svg?token=ELKDO3ZZ8T)](https://codecov.io/gh/kzys/stats)

stats is a command line tool to calculate summary statistics

```
% ps uax | head | awk '{print $6}'
RSS
1908
435984
81360
86868
25852
7272
167028
18096
29676
% ps uax | stats -k6
avg: 8680.55
min: 0.00
p50: 1668.00
p90: 13304.00
p99: 85712.00
max: 558864.00

histogram (binSize=55886.50, bins=10)
[     0.00,  55886.50) 500 █████████████████████████████████████████████████████▍
[ 55886.50, 111773.00)   8 ▊
[111773.00, 167659.50)   2 ▏
[167659.50, 223546.00)   1  
[223546.00, 279432.50)   1  
[279432.50, 335319.00)   1  
[335319.00, 391205.50)   0  
[391205.50, 447092.00)   0  
[447092.00, 502978.50)   0  
[502978.50, 558865.00)   0  
% 
