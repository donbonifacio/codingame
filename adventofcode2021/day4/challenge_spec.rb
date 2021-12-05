require_relative './challenge.rb'

TEST_DRAW_NUMBERS = "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1".freeze

TEST_BOARDS = "

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7".freeze

TEST_DATA = (TEST_DRAW_NUMBERS + TEST_BOARDS).freeze
INPUT_DATA = File.read('./input.txt').freeze

RSpec.describe Day4 do
  context 'when working with provided data' do
    it 'properly deserializes bingo data' do
      bingo = Day4.deserialize(TEST_DATA)

      expect(bingo.draw_numbers).to eql([7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1])
      expect(bingo.boards[0].at(2, 2)).to eql(14)
      expect(bingo.boards[1].at(2, 2)).to eql(7)
      expect(bingo.boards[2].at(3, 3)).to eql(6)
    end

    it 'has expected data for TEST_DATA' do
      expect(Day4.winner_score(TEST_DATA)).to eq(4512)
    end

    it 'has expected data for TEST_DATA but for a column win' do
      data = ('24,9,26,6,3'.freeze + TEST_BOARDS).freeze
      expect(Day4.winner_score(data)).to eq(771)
    end
  end

  context 'when working with input data' do
    it 'has expected data for INPUT_DATA' do
      expect(Day4.winner_score(INPUT_DATA)).to eq(51_034)
    end
  end
end
