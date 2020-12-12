(ns Solution
  (:gen-class))

(defn count-occupied
  [seats]
  (reduce (fn [acc curr]
            (if (= curr \#)
              (inc acc)
              acc))
          0 seats)
  )

(defn gather-adjacent
  [raw x y]
  (let [lines (clojure.string/split-lines raw)]
    (for [line (range -1 2)
          column (range -1 2)]
      (let [pos-x (+ column x)
            pos-y (+ line y)]
        (if (and (= pos-x x) (= pos-y y))
          [(get (get lines pos-y) pos-x)]
          (get (get lines pos-y) pos-x)))
      )))

(defn calculate-next
  [raw x y]
  (let [adjacent (gather-adjacent raw x y)
        occupied (count-occupied adjacent)
        target (first (nth adjacent 4))]
    (cond
      (and (= target \L) (zero? occupied)) "#"
      (and (= target \#) (> occupied 3)) "L"
      :else (str target))))

; (flatten (interpose "\n" (partition n list))))

(defn run
  "Runs the algorithm for the given input."
  [input]
  (->> (for [y (range 0 10)
             x (range 0 10)]
          [x y])
       (map #(calculate-next input (first %) (second %)))
       (partition 10)
       (map clojure.string/join)
       (clojure.string/join "\n")))


(defn -main
  "The main function, will read data from the stdin and call the worker
  function."
  [& args]
  (println 1))