require_relative './challenge1.rb'

RSpec.describe 'Day1.1' do
  it 'runs the default example' do
    data = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]
    expect(Day1.measure_increases(data)).to eq(7)

  end
end
