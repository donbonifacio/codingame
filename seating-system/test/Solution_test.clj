(ns Solution-test
  (:require [clojure.test :refer :all]
            [Solution :as solution]))

(def round-one
"#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##")

(def round-two
"#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##")

#_(deftest calculate-next-test
  (is (= "#" (solution/calculate-next round-one 0 0)))
  (is (= "." (solution/calculate-next round-one 1 0)))
  (is (= "L" (solution/calculate-next round-one 1 1)))
  (is (= "." (solution/calculate-next round-one 3 2))))

(deftest run-test
  (is (= round-two (solution/run round-one))))
