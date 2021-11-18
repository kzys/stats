# stats

stats is a command line tool to calculate summary statistics

```
% ps uax | head | awk '{print $6}'
RSS
73916
75508
1732
40132
4528
689648
22900
27168
7288
% ps uax | stats -k6
avg: 9525.53
min: 0.00
p50: 2264.00
p90: 18680.00
p99: 123668.00
max: 689652.00

histogram (bins=10)
458: █████████████████████████████████████████████████████████▎
  8: █ 
  1: ▏
  0:  
  0:  
  0:  
  0:  
  0:  
  0:  
  0:  
% 
