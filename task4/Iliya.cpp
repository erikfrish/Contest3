ld dpg[510];
bool dpgwas[510];

ld gt(int k)
{
  if (k < 3)
    return 0;

  if (dpgwas[k]) return dpg[k];

  dpgwas[k] = 1;
  dpg[k] = getS(pts[0], pts[k / 2], pts[k - 1]) + gt((k - 3) / 2) + gt((k - 3) / 2 + k % 2);
  return dpg[k];
}

void solve()
{
  cin >> n;
  for (int i = 0; i < n; i++)
    for (int j = 0; j < n; j++)
      dpwas[i][j] = false;
  double ang = 2 * M_PI / n;
  
  ld R = 1.0 / (2 * sin(M_PI / n));
  for (int i = 0; i < n; i++)
  {
    pts.push_back(Point(R * sin(i * ang), R * cos(i * ang)));
  }

  int i = 0;
  ld maxans = 0;
  for (int j = i+1; j < n; j++)
  {
    for (int k = j + 1; k < n; k++)
    {
      ld a1 = gt((j - 1) - (i + 1) + 1);
      ld a2 = gt((k - 1) - (j + 1) + 1);
      ld a3 = gt(n - k - 1);
      ld pans = getS(pts[i], pts[j], pts[k]) + a1 + a2 + a3;
      if (pans > maxans)
      {
        maxans = pans;
      }
    }
  }
  cout << maxans << en;
}