# Solutions for Day 4.1
module Day4
  module_function

  # Represents the bingo boards and draw numbers
  class Bingo
    attr_reader :draw_numbers, :boards

    def initialize(draw_numbers, boards)
      @draw_numbers = draw_numbers
      @boards = boards
    end
  end

  # Represents a bingo board
  class Board
    attr_reader :matrix

    def initialize(raw_board)
      @matrix = raw_board
    end

    def at(x_idx, y_idx)
      matrix[x_idx][y_idx]
    end

    def score
      1
    end
  end

  def winner_score(raw_bingo)
    bingo = deserialize(raw_bingo)
    board_score(Board.new)
  end

  def board_score(board)
    1
  end

  def deserialize(raw_bingo)
    lines = raw_bingo.lines.map(&:strip)

    draw_numbers = lines.first.split(',').map(&:strip).map(&:to_i)
    lines = lines.drop(2)

    boards = []
    loop do
      board_lines = lines.take(5)
      raw_board = board_lines.map { |line| line.split(/\s+/).map(&:strip).map(&:to_i) }
      boards << Board.new(raw_board)
      lines = lines.drop(6)
      break if lines.count < 5
    end

    Bingo.new(draw_numbers, boards)
  end
end
