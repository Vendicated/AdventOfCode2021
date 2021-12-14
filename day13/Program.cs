using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;


var lines = File.ReadAllLines("../../../input.txt");
var rowCount = 
    lines
        .Where(l => l.Contains(','))
        .Select(l => int.Parse(l[..l.IndexOf(',')]))
        .Max() + 1;
var colCount = 
    lines
        .Where(l => l.Contains(','))
        .Select(l => int.Parse(l[(l.IndexOf(',')+1)..]))
        .Max() + 1;

var coords = new bool[colCount][];
for (var i = 0; i < colCount; i++)
{
    coords[i] = new bool[rowCount];
}

var yFolds = new List<int>();
var xFolds = new List<int>();

foreach (var line in lines)
{
    var split = line.Split(',');
    if (split.Length == 2)
    {
        coords[int.Parse(split[1])][int.Parse(split[0])] = true;
    }
    else if (line.StartsWith("fold along y="))
    {
        yFolds.Add(int.Parse(line["fold along y=".Length..]));
    } else if (line.StartsWith("fold along x="))
    {
        xFolds.Add(int.Parse(line["fold along x=".Length..]));
    }
}

int CountDots(bool[][] arr) => arr.Sum(r => r.Sum(b => b ? 1 : 0));

bool[][] MergeX(int foldX, bool[][] arr)
{
    if (foldX == -1) return arr;
    var left = arr.Select(l => l[..foldX]).ToArray();
    var right = arr.Select(l => l[(foldX + 1)..].Reverse().ToArray()).ToArray();

    for (var i = 0; i < left.Length; i++)
    {
        for (var j = 0; j < left[0].Length; j++)
        {
            left[i][j] = left[i][j] || right[i][j];
        }
    }
    return left;
}

var dotsAfterFirstFold = CountDots(MergeX(xFolds[0], coords));
bool[][] topCoords = coords;
for (var idx = 0; idx < yFolds.Count; idx++)
{
    var foldY = yFolds[idx];
    var foldX = idx < xFolds.Count ? xFolds[idx] : -1;
    
    var botCoords = MergeX(foldX, topCoords[(foldY + 1)..].Reverse().ToArray());
    topCoords = MergeX(foldX, topCoords[..foldY]);

    for (var i = 0; i < topCoords.Length; i++)
    {
        for (var j = 0; j < topCoords[0].Length; j++)
        {
            topCoords[i][j] = topCoords[i][j] || botCoords[i][j];
        }
    }
}

foreach (var foldedCoord in topCoords)
{
    foreach (var b in foldedCoord)
    {
        Console.Write(b ? "#" : ".");
    }
    Console.WriteLine();
}

Console.WriteLine($"After first fold: {dotsAfterFirstFold}");
var visibleDots = CountDots(topCoords);
