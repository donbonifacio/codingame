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

(def round-three
"#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##")

(def round-four
"#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##")

(def round-five
"#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##")

(deftest calculate-next-test
  (defn helper [raw x y] (solution/calculate-next (clojure.string/split-lines raw) x y))
  (is (= "#" (helper round-one 0 0)))
  (is (= "." (helper round-one 1 0)))
  (is (= "L" (helper round-one 1 1)))
  (is (= "." (helper round-one 3 2))))

(deftest run-round-test
  (is (= round-two (solution/run-round round-one)))
  (is (= round-three (solution/run-round round-two)))
  (is (= round-four (solution/run-round round-three)))
  (is (= round-five (solution/run-round round-four))))

(deftest run-test
  (is (= 37 (solution/run round-one))))
