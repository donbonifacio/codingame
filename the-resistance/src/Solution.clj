(ns Solution
  (:gen-class))

(def morse-table {\A ".-"
                  \B "-..."
                  \C "-.-."
                  \D "-.."
                  \E ".-"
                  \F "..-." \G "--." \H  "...."
                  \I "..-." \J ".---" \K "-.-." \L ".-.."
                  \M "--." \N "-..." \O "---" \P ".--."
                  \Q "--.-" \R ".-.." \S "..." \T "-"
                  \U "..-." \V "...-" \W ".---" \X "-..-"
                  \Y "-.--" \Z "--.."})

(defn morse
  ([text]
   (morse text ""))
  ([text sep]
   (clojure.string/join sep (map #(get morse-table %) text))))

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
