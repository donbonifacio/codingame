(ns Player
  (:gen-class))

; The while loop represents the game.
; Each iteration represents a turn of the game
; where you are given inputs (the heights of the mountains)
; and where you have to print an output (the index of the mountain to fire on)
; The inputs you are given are automatically updated according to your last actions.
(defn mountain-sizes
  []
  (loop [i 8
         sizes []]
    (if (= i 0)
      sizes
      (let [mountainH (read)]
        ; mountainH: represents the height of one mountain.
        (recur (dec i) (conj sizes mountainH)))
      )))

(defn -main [& args]
  (while true
    ; The index of the mountain to fire on.
    (let [sizes (mountain-sizes)
          height-map (->> sizes
                          (map-indexed (fn [idx height] [height idx]))
                          (into {}))
          m (apply max sizes)
          target (get height-map m)]
      #_(binding [*out* *err*]
        (println target m sizes height-map))

      (println target))))
