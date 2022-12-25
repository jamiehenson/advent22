class Coordinate
  attr_accessor :x, :y

  def initialize(x, y)
    @x = x
    @y = y
  end
end

def create_cave(rock_coords, xh, xl, yh, yl)
  cave = Array.new(yh - yl + 1) { Array.new(xh - xl + 1) { '.' } }
  
  rock_coords.each do |rock|
    for i in 1..rock.length - 1 do
      if rock[i - 1].x == rock[i].x
        sorted = [rock[i - 1].y, rock[i].y].sort
        for j in sorted.first..sorted.last do
          cave[j - yl][rock[i].x - xl] = "#"
        end
      elsif rock[i - 1].y == rock[i].y
        sorted = [rock[i - 1].x, rock[i].x].sort
        for j in sorted.first..sorted.last do
          cave[rock[i].y - yl][j - xl] = "#"
        end
      end
    end
  end

  cave
end

def simulate_sand(cave, xl, yh, yl)
  sand_coords = []
  pouring = true

  while (pouring) do
    motion = true
    sx = 500
    sy = 0

    while motion do
      if sy > yh || sx < xl
        pouring = false
        motion = false
      elsif cave[sy + 1 - yl][sx - xl] == "."
        sy += 1
      elsif cave[sy + 1 - yl][sx - 1 - xl] == "."
        sy += 1
        sx -= 1
      elsif cave[sy + 1 - yl][sx + 1 - xl] == "."
        sy += 1
        sx += 1
      else
        motion = false
        sand_coords << Coordinate.new(sx, sy)
        cave[sy - yl][sx - xl] = "o"
      end
    end
  end

  sand_coords
end

def main
  rocks = File.open("rocks.txt").readlines.map(&:chomp)

  xl, xh, yl, yh = 500, 500, 0, 0
  rock_coords = Array.new(rocks.length) { [] }

  rocks.each_with_index do |line, index|
    line.split(" -> ").each do |coord|
      x, y = coord.split(",").map(&:to_i)
      xh = x if x > xh
      xl = x if x < xl
      yh = y if y > yh
      yl = y if y < yl
      rock_coords[index] << Coordinate.new(x, y)
    end
  end

  cave = create_cave(rock_coords, xh, xl, yh, yl)

  sand_coords = simulate_sand(cave, xl, yh, yl)

  puts "Sand grains: #{sand_coords.length}"
  puts cave.map { |x| x.join(' ') }
end

main