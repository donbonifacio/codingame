(ns Solution
  (:gen-class))

; Auto-generated code below aims at helping you parse
; the standard input according to the problem statement.

(defn abs [n]
  (if (neg? n)
    (* n -1)
    n))

(defn -main [& args]
  (let [n (read) _ (read-line) 
        temps (clojure.string/split (clojure.string/trim (read-line)) #" ")]
    ; n: the number of temperatures to analyse
    ; temps: the n temperatures expressed as integers ranging from -273 to 5526

    (binding [*out* *err*]
      (println "n:" n "temps:" temps "empty?" (count temps)))

    (if (zero? n)
      (println 0)
      (let [result (loop [temps (map read-string temps)
                          closest 9999]
                     (if (nil? (first temps))
                       closest
                       (let [curr (first temps)
                             more (rest temps)
                             closest (cond
                                       (and (= (abs curr) (abs closest))
                                            (or (pos? curr) (pos? closest))) (abs curr)
                                       (< (abs curr) (abs closest)) curr
                                       :else closest)]
                         (recur more closest))))]

        ; Write answer to stdout
        (println result)))))
