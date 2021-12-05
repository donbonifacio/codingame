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

    def play_number(number)
      5.times do |x|
        5.times do |y|
          matrix[x][y] = '*' if at(x, y) == number
        end
      end
    end

    def won?
      matrix.each do |line|
        return true if line.join == '*****'
      end

      5.times do |idx|
        return true if matrix.map { |line| line[idx] }.join == '*****'
      end

      false
    end

    def score(drawn_number)
      unmarked = matrix
                 .flatten
                 .reject { |entry| entry == '*' }
                 .sum

      unmarked * drawn_number
    end
  end

  def winner_score(raw_bingo, lucky = :first)
    bingo = deserialize(raw_bingo)
    result = if lucky == :first
               play_until_winner(bingo)
             else
               play_until_last_winner(bingo)
             end

    raise 'No winner' unless result[:winner]

    result[:board].score(result[:winner_number])
  end

  def play_until_winner(bingo)
    play(bingo) do |result|
      return result
    end
  end

  def play_until_last_winner(bingo)
    winners = []
    play(bingo) do |result|
      winners << result if result[:winner]
    end
    winners.last
  end

  def play(bingo)
    boards = bingo.boards.clone
    bingo.draw_numbers.each do |drawn_number|
      # puts "-- Number: #{drawn_number}"
      to_delete = []
      boards.each do |board|
        board.play_number(drawn_number)
        if board.won?
          # puts "Won? #{board.won?} #{board.inspect}"
          result = { winner: true,
                     board: board,
                     winner_number: drawn_number }
          yield result
          to_delete << board
        end
      end
      to_delete.each { |board| boards.delete(board) }
    end
    { winner: false }
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
