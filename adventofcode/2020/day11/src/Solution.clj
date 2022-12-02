(ns Solution
  (:gen-class))

(defn count-occupied
  "Given a collection of sets, counts the ones that are occupied."
  [seats]
  #_(count (filter #(= % \#) seats)) ; simpler but adds +2s
  (reduce (fn [acc curr]
            (if (= curr \#)
              (inc acc)
              acc))
          0 seats))

(defn gather-adjacent
  "Returls a collection with all the adjacent cells. The middle one is
  injected as an array (it's a hack to be able to get it upstream, should
  be refactored)"
  [lines x y]
  (for [line (range -1 2)
        column (range -1 2)]
    (let [pos-x (+ column x)
          pos-y (+ line y)]
      (if (and (= pos-x x) (= pos-y y))
        [(get (get lines pos-y) pos-x)]
        (get (get lines pos-y) pos-x)))))


(defn calculate-next
  "Given all the seats and a position, returns the next interation for that
  position."
  [input-lines x y]
  (let [adjacent (gather-adjacent input-lines x y)
        occupied (count-occupied adjacent)
        target (first (nth adjacent 4))]
    (cond
      (and (= target \L) (zero? occupied)) "#"
      (and (= target \#) (> occupied 3)) "L"
      :else (str target))))

(defn run-round
  "Runs the algorithm for the given input. It has some setup that could maybe
  be extracted and ran once. It will iterate per position to find the next
  iteration, and then builds everything up. Definetely could use improvements"
  [input]
  (let [input-lines (clojure.string/split-lines input)
        size-y (count input-lines)
        size-x (count (first input-lines))]
    (->> (for [y (range 0 size-y)
               x (range 0 size-x)]
            [x y])
         (map #(calculate-next input-lines (first %) (second %)))
         (partition size-x)
         (map clojure.string/join)
         (clojure.string/join "\n"))))

(defn run
  "Runs round after round until the rounds don't change anymore."
  ([input]
   (run input 0))
  ([input round-number]
   (let [new-round (run-round input)]
     (cond
        (= new-round input)
          (count-occupied new-round)
        (> round-number 1000)
          :failure
        :else
          (recur new-round (inc round-number))))))

(defn -main
  "The main function, will read data from the stdin and call the worker
  function."
  [& args]
  (println "Seats " (run (slurp "input.txt")))
  (shutdown-agents))
