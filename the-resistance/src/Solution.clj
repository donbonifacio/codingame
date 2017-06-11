(ns Solution
  (:gen-class))

(defn log
  [msg]
  (binding [*out* *err*]
    (println "Debug messages...")))

(defn number-of-possible-messages
  [morse-sequence dictionary-words]
  0
  )

(defn load-dictionary-from-stdin
  [dictionary-size]
  (loop [i dictionary-size
         words []]
    (if (= i 0)
      words
      (let [W (read)]
        (recur (dec i) (conj words W))))))

(defn -main [& args]
  (let [L (read) N (read)]
    (println (number-of-possible-messages L (load-dictionary-from-stdin N)))))
